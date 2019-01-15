package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"bodyweight/database"
	"bodyweight/training"

	"github.com/aws/aws-lambda-go/lambda"
)

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
	db = database.NewDynamoDB("bodyweight")
}

// Start starts the Lambda Handler
func Start() {
	if isClient() {
		connectToServer()
	}

	log.Println("Starting FN")
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

		entry, _ := db.GetLastUsedEntry(event.Session.User.UserID)
		if entry == nil {
			// First Usage
			return responseBuilder().
				speak(speechWelcome + speechDefineUser).
				reprompt(speechStartTraining + speechExitIfMute), nil

		}
		// Welcome Back
		// return responseBuilder().
		// 	speak(speechWelcome + fmt.Sprintf(speechPersonal, entry.UserName) + speechStartTraining).
		// 	reprompt(speechStartTraining + speechExitIfMute), nil

		return responseBuilder().speak(training.StartTraining(nil)), nil

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

func handleStartTraining(ctx context.Context, event Request) (interface{}, error) {
	// TODO: Es macht einen Unterschied, ob der Intent direkt gestartet wird, oder ob zunächst über einen
	// Launch Request gestartet wird.

	user := event.RequestBody.Intent.Slots["user"]
	if user.Value != "" {
		log.Println("we got user from Alexa")
		user.ConfirmationStatus = "CONFIRMED"
		db.CreateEntry(event.Session.User.UserID, user.Value, training.GetBeginningState(), "Start")
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
		return responseBuilder().addDelegateDirective(&event.RequestBody.Intent, false), nil
	}

	// zeige zustand auf (tag + )

	text := fmt.Sprintf("Herzlich Willkommen zurück %s. ", user.Value)
	text += training.AnnounceDailyTraining(&entry.TrainingState)
	return responseBuilder().speak(text), nil

	//erkläre Übung
	// sind sie bereit? sage: ich bin bereit
	// starte Übung
	// spiele audio playback

}

func defineUser(ctx context.Context, event Request) (interface{}, error) {
	user := event.RequestBody.Intent.Slots["user"]
	log.Printf("Session User: %+v", user.Value)
	db.CreateEntry(event.Session.User.UserID, user.Value, training.GetBeginningState(), "Start")
	return responseBuilder().
		speak(fmt.Sprintf("Hallo %s. Schön dass Du hier bist", user.Value)), nil

}

func handleUnknown(ctx context.Context, event Request) (interface{}, error) {
	return responseBuilder().speak(speechUnknown + " " + speechStartTraining).
		reprompt(speechExitIfMute), nil
}
