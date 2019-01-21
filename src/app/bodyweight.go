package app

import (
	"bodyweight/tools"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"bodyweight/database"
	"bodyweight/training"

	"github.com/aws/aws-lambda-go/lambda"
)

/* TODO:
Alexa-Return warum immer nil
*/

const debug = true

const (
	// Launch Intent
	speechWelcome    = `Willkommen beim <lang xml:lang="en-US">Bodyweight Training</lang>. `
	speechDefineUser = `Du benutzt das <lang xml:lang="en-US">Bodyweight Training</lang> zum ersten Mal. 
	Du solltest zunächst deinen Namen festlegen. Sage hierzu bitte <break time="500ms"/>mein Name ist<break time="500ms"/> und deinen Vornamen. `
	speechStartTraining = `
	Du kannst die Trainings mit 
	<break strength="x-strong"/>starte das training<break strength="x-strong"/>
	starten. `

	speechUnknown    = "Ich kann dich leider nicht verstehen. "
	speechExitIfMute = "Wenn Du nichts mehr sagts, wird das Programm beendet. "
	speechEnde       = "Auf Wiedersehen und bis bald. "
	speechPersonal   = `%s, es ist schön, dass du wieder da bist. `
)

var db database.DB

func init() {
	if !debug {
		log.SetOutput(ioutil.Discard)
	}
}

// Start starts the Lambda Handler
func Start() {

	// start as Client
	if tools.IsTestClient() {
		tools.ConnectToServer()
		os.Exit(0)
	}

	// start as Server
	log.Println("Starting FN")
	db = database.Accessor()
	lambda.Start(HandleRequest)
}

// HandleRequest handles every Request / Base-Entry FN
func HandleRequest(ctx context.Context, event Request) (interface{}, error) {

	if event.RequestBody.Type == "IntentRequest" {
		log.Printf("Intent: %+v\n", event.RequestBody.Intent.Name)
	} else {
		log.Printf("Request: %+v\n", event.RequestBody.Type)
	}

	if isDelegate() {
		log.Println("we delegate")
		result, err := delegateRemote(event)
		log.Println("and got result: ")
		return result, err
	}

	switch event.RequestBody.Type {

	case alexaLaunchRequest:
		uid := event.Session.User.UserID // Amazon UID
		entry, _ := db.GetLastUsedEntry(uid)
		if entry == nil {
			// First Usage
			return responseBuilder().
				speak(speechWelcome+speechDefineUser).
				withSimpleCard("Bodyweight Training", "Herzlich willkommen").
				reprompt(speechStartTraining + speechExitIfMute), nil

		}

		// Welcome Back
		return responseBuilder().
			speak(speechWelcome+fmt.Sprintf(speechPersonal, entry.UserName)+speechStartTraining).
			withSimpleCard("Bodyweight Training", "Herzlich willkommen "+entry.UserName).
			reprompt(speechStartTraining + speechExitIfMute), nil

	case alexaSessionEndRequest:
		log.Println("End-Reason: ", event.RequestBody.Reason)
		return nil, nil

	case alexaAudioplayerStartedRequest, alexaAudioplayerNearlyFinishedRequest, alexaAudioplayerFinishedRequest:
		log.Println("Audioplayer Request")
		return nil, nil

	case alexaIntentRequest:
		switch event.RequestBody.Intent.Name {
		case alexaStopIntent:
			return responseBuilder().speak(speechEnde).withShouldEndSession(), nil
		case alexaHelpIntent:
			return responseBuilder().speak(speechStartTraining), nil
		case alexaAudioTestIntent:
			/* die Audioausgabe läuft parallel zum Dialog.
			D.h. wenn im Dialog nix mehr passiert, wird der Skill beendet */
			return responseBuilder().
				addAudioPlayerPlayDirective(
					"https://github.com/gipde/bodyweight/raw/master/contrib/alien-spaceship_daniel_simion.mp3"), nil
		case alexaStartTrainingIntent:
			return handleStartTraining(ctx, event)
		case alexaBereitIntent:
			return handleBereit(ctx, event)
		case alexaDefineUserIntent:
			return defineUser(ctx, event)
		case alexaFallbackIntent:
			return handleUnknown(ctx, event)
		default:
			return handleUnknown(ctx, event)
		}
	}
	return handleUnknown(ctx, event)
}

func getlastUserEntry(event Request) (dbEntry *database.Entry, notFoundMsg *Response) {
	// TODO: Es macht einen Unterschied, ob der Intent direkt gestartet wird, oder ob zunächst über einen
	// Launch Request gestartet wird.

	// without Slot-Value means called without a username set in slot
	user := event.RequestBody.Intent.Slots["user"]
	if user.Value != "" {
		log.Println("we got user from Alexa")
		user.ConfirmationStatus = "CONFIRMED"
		db.CreateEntry(getNewUserEntry(event.Session.User.UserID, user.Value))
	}

	entry, _ := db.GetLastUsedEntry(event.Session.User.UserID)
	if entry != nil {
		log.Println("we got user from DB")
		user.Value = entry.UserName
		user.ConfirmationStatus = "CONFIRMED"
	}

	log.Printf("Session User: %+v", user.Value)
	log.Printf("Complete State: %s", event.RequestBody.DialogState)

	// noch kein User festgelegt
	if user.Value == "" {
		return nil, responseBuilder().addDelegateDirective(&event.RequestBody.Intent, false)
	}

	return entry, nil

}

func handleStartTraining(ctx context.Context, event Request) (interface{}, error) {
	// TODO: Es macht einen Unterschied, ob der Intent direkt gestartet wird, oder ob zunächst über einen
	// Launch Request gestartet wird.

	entry, errmsg := getlastUserEntry(event)
	if errmsg != nil {
		return errmsg, nil
	}

	// zeige zustand auf (tag + )

	text := fmt.Sprintf("Herzlich Willkommen zurück %s. ", entry.UserName)
	text += training.AnnounceDailyTraining(&entry.TrainingState)
	return responseBuilder().withSimpleCard("Bodyweight Training", "für "+entry.UserName).speak(text), nil

	//erkläre Übung
	// sind sie bereit? sage: ich bin bereit
	// starte Übung
	// spiele audio playback

}

func handleBereit(ctx context.Context, event Request) (interface{}, error) {
	return nil, nil
}

func getNewUserEntry(userid, username string) *database.Entry {
	return &database.Entry{
		AlexaID:       userid,
		Date:          time.Now(),
		Desc:          "Start",
		TrainingState: training.GetBeginningState(),
		UserName:      username,
	}
}

func defineUser(ctx context.Context, event Request) (interface{}, error) {
	user := event.RequestBody.Intent.Slots["user"]
	userEntry := *getNewUserEntry(event.Session.User.UserID, user.Value)

	// we check, if this user already exists
	entries, _ := db.GetEntries(event.Session.User.UserID)
	if entries != nil {
		for _, entry := range *entries {
			if entry.UserName == user.Value {
				// set userEntry
				userEntry = entry
				userEntry.Date = time.Now()

				// delete old entry
				db.DeleteItem(entry.AlexaID, entry.Date)

			}
		}
	}

	db.CreateEntry(&userEntry)

	return responseBuilder().
		speak(fmt.Sprintf("Hallo %s. Schön dass Du mit mir trainierst", user.Value)), nil

}

func handleUnknown(ctx context.Context, event Request) (interface{}, error) {
	return responseBuilder().speak(speechUnknown + " " + speechStartTraining).
		reprompt(speechExitIfMute), nil
}
