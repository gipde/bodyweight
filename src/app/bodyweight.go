package app

import (
	// "os/user"
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
	speechWelcomNewUser = `Hallo %s. Schön dass Du mit mir trainierst. `

	speechExplainTraining     = `Lass Dir zunächst das Training erklären. `
	speechExplainThisExercise = `Lass Dir diese nächste Übung erklären. `
	speechExplainNextExercise = `Lass Dir die nächste Übung erklären. `
	speechStart               = `Sage: <break strength="x-strong"/>bereit,<break strength="x-strong"/>`
	speechStartThisExercise   = speechStart + `wenn du mit dieser Übung beginnen möchtest. `
	speechStartNextExercise   = speechStart + `wenn du mit der nächsten Übung beginnen möchtests. `
	speechDone                = `Geschafft. `
	speechReadyForToday       = `Jetzt kannst Du Dich erholen. Bis zum nächsten mal. `
	speechStufenintervall     = `Jede Übung in diesem Block wird 4,5 bis 7,5 Minuten ausgeführt.
	Du beginnst mit einer Wiederholung, machst eine Pause, machst zwei Wiederholungen und so weiter. 
	Ab der Hälft der Zeit reduzierst Du die Sätze jeweils um eine Wiederholung. Wenn Du bereits 
	vorher nicht mehr kannst, kannst Du auch früher schon reduzieren. Ist noch Zeit übrig, 
	beginnst du mit einer neuen Stufe wieder mit zunächst 1 Wiederholung`
	speechIntervallsatz = `Beim Intervallsatz wird sind von jeder Übung 3 Sätze mit 6-12 Wiederholungen
	durchzuführen. Bei einem Satz solltest Du eigentlich bis zum Muskelversagen kommen. Für einen Satz hast Du
	genau 3 Minuten zeit. Falls Du früher ferwig wirst, kannst Du die Zeit als Pause nutzen. Die Sätze beginnen damit
	immer genau im Abstand von 3 Minuten. Wechsle bei einseitigen Übungen nach jedem Satz die Seite,es sei denn die 
	Wiederholungen sind im Wechsel durchzuführen. Beginnst zunächst mit der schwächeren Seite`
	speechSupersatz = `Ein Übungspaar bildet einen Supersatz, der jeweils 4 Minuten dauert. Bei der ersten Übung
	sind jeweils 1 - 5 Wiederholungen, bei der zweiten Übung sind 6-12 Wiederholungen zu absolvieren. Pro Paar sind 2 
	Supersätze direkt nacheinander durchzuführen. Bei einseitigen Übungen wechseln Sie nach jeder Wiederholung die Seite.`
	speechHochintensitaetssatz = `Acht Sätze mit jeweils 20 Sekunden Training gefolgt von je 10 Sekunden Pause insgesamt 4 Minuten lang.`
	speechZirkelintervall      = `Das Zirkelintervall besteht aus insgesamt 3 verschiedenen Übungen, von denen jeweils 
	eine angegebene Anzahl von Wiederholungen durchzuführen ist. Ohne Pause führen sie die Übungen im Wechsel durch.
	Versuchen Sie das Zirkelintervall insgesamt 20 Minuten durchzuführen.`

	speechUnknown    = "Ich kann dich leider nicht verstehen. "
	speechExitIfMute = "Wenn Du nichts mehr sagts, wird das Programm beendet. "
	speechEnde       = "Auf Wiedersehen und bis bald. "
	speechPersonal   = `%s, es ist schön, dass du wieder da bist. `
	speechHelp       = `Du brauchst Hilfe?
	
	In diesem Skill kannst Du mit:
	- erkläre das Training Dir das Training erklären lassen
	- erkläre die nächste Übung - dir die nächste Übung erklären lassen`
)

var db database.DB

func init() {
	if !debug {
		log.SetOutput(ioutil.Discard)
	} else {
		tools.SetDebug()
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

func getUser(event Request) *database.Entry {
	uid := event.Session.User.UserID // Amazon UID
	entry, _ := db.GetLastUsedEntry(uid)
	return entry
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

	user := getUser(event)
	switch event.RequestBody.Type {

	case alexaLaunchRequest:
		if user == nil {
			// First Usage

			return responseBuilder().
				speak(speechWelcome+speechDefineUser).
				withSimpleCard("Bodyweight Training", "Herzlich willkommen, du musst zunächst deinen Namen festlegen.").
				reprompt(speechDefineUser + speechExitIfMute), nil

		}

		// Welcome Back
		return responseBuilder().
			speak(speechWelcome+fmt.Sprintf(speechPersonal, user.UserName)+speechExplainTraining).
			withSimpleCard(user.TrainingState.ShortProgress(), user.TrainingState.DayShortDescription()).
			reprompt(speechExplainTraining + speechExitIfMute), nil

	case alexaSessionEndRequest:
		log.Println("End-Reason: ", event.RequestBody.Reason)
		return nil, nil

	case alexaAudioplayerStartedRequest, alexaAudioplayerNearlyFinishedRequest, alexaAudioplayerFinishedRequest:
		log.Println("Audioplayer Request")
		return nil, nil

	case alexaIntentRequest:

		// intents without user
		switch event.RequestBody.Intent.Name {
		case alexaDefineUserIntent:
			return defineUser(event)
		case alexaExplainTrainingMethodIntent:
			return handleExplainTrainingMethod(event)

		case alexaStopIntent:
			return responseBuilder().speak(speechEnde).withShouldEndSession(), nil
		case alexaHelpIntent:
			return responseBuilder().speak(speechHelp), nil
		case alexaAudioTestIntent:
			/* die Audioausgabe läuft parallel zum Dialog.
			D.h. wenn im Dialog nix mehr passiert, wird der Skill beendet */
			return responseBuilder().
				addAudioPlayerPlayDirective(
					"https://github.com/gipde/bodyweight/raw/master/contrib/alien-spaceship_daniel_simion.mp3"), nil

		case alexaFallbackIntent:
			return handleUnknown(ctx, event)
		}

		// intents which require user
		if user == nil {
			return responseBuilder().speak(speechDefineUser).reprompt(speechDefineUser + speechExitIfMute), nil
		}
		switch event.RequestBody.Intent.Name {
		case alexaExplainTrainingIntent:
			return responseBuilder().speak(user.TrainingState.ExplainTraining()+speechExplainThisExercise).
				withSimpleCard(user.TrainingState.ShortProgress(), user.TrainingState.DayShortDescription()).
				reprompt(speechExplainThisExercise + speechExitIfMute), nil
		case alexaExplainExerciseIntent:
			return responseBuilder().speak(user.TrainingState.ExplainExercise()+speechStartThisExercise).
				withSimpleCard(user.TrainingState.ShortProgress(), user.TrainingState.UnitShortDescription()).
				reprompt(speechStartThisExercise + speechExitIfMute), nil
		case alexaStartTrainingIntent:
			return handleStartTraining(user, event)
		}
	}
	return handleUnknown(ctx, event)
}

func handleExplainTrainingMethod(event Request) (*Response, error) {
	return responseBuilder().speak("nicht implementiert").reprompt(speechExitIfMute), nil
}

func handleStartTraining(user *database.Entry, event Request) (*Response, error) {
	//TODO: check if is directyl called
	text := user.TrainingState.InstructTraining()

	defer func() {
		user.Date = time.Now()
		db.CreateEntry(user)
	}()

	if user.TrainingState.WasLastUnit() {
		return responseBuilder().speak(text + speechDone + speechReadyForToday).withShouldEndSession(), nil
	}
	return responseBuilder().speak(text + speechDone + speechExplainNextExercise), nil

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

func getNewUserEntry(userid, username string) *database.Entry {
	return &database.Entry{
		PK: database.PK{
			AlexaID: userid,
			Date:    time.Now(),
		},
		TrainingState: training.GetBeginningState(),
		UserName:      username,
	}
}

func defineUser(event Request) (interface{}, error) {
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
				db.DeleteItem(entry.PK)
			}
		}
	}

	db.CreateEntry(&userEntry)

	return responseBuilder().
		speak(fmt.Sprintf(speechWelcomNewUser, user.Value) + speechExplainTraining).
		reprompt(speechExplainTraining + speechExitIfMute), nil

}

func handleUnknown(ctx context.Context, event Request) (interface{}, error) {
	return responseBuilder().speak(speechUnknown + " " + speechExplainTraining).
		reprompt(speechExitIfMute), nil
}
