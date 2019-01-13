package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gipde/bodyweight/training"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gipde/bodyweight/database"
)

const debug = true

const SPEECH_DEFINE_USER = `Du benutzt das <lang xml:lang="en-US">Bodyweight Training</lang> zum ersten Mal. 
	Du solltest zunächst deinen Namen festlegen. Sage hierzu bitte <break time="500ms"/>mein Name ist<break time="500ms"/> und deinen Vornamen. `
const SPEECH_START_TRAINING = `
	Du kannst die Trainings mit 
	<break strength="x-strong"/>starte das training<break strength="x-strong"/>
	starten. `

const SPEECH_UNKNOWN = "Ich kann dich leider nicht verstehen. "

const SPEECH_EXIT_IF_MUTE = "Wenn Du nichts mehr sagts, wird das Programm beendet. "

const SPEECH_ENDE = "Auf Wiedersehen und bis bald. "

const SPEECH_WELCOME = `Willkommen beim <lang xml:lang="en-US">Bodyweight Training</lang>. `

const SPEECH_PERSONAL = `%s, es ist schön, dass du wieder da bist. `

func handleStartTraining(ctx context.Context, event Request) (interface{}, error) {
	user := event.RequestBody.Intent.Slots["user"]
	if user.Value != "" {
		log.Println("we got user from Alexa")
		user.ConfirmationStatus = "CONFIRMED"
		db.CreateEntry(event.Session.User.UserID, user.Value, training.TrainingState{Level: training.BASISPROGRAM, Week: 0, Day: 0, Unit: 0}, "Start")
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
	db.CreateEntry(event.Session.User.UserID, user.Value, training.TrainingState{Level: training.BASISPROGRAM, Week: 0, Day: 0, Unit: 0}, "Start")
	return responseBuilder().
		speak(fmt.Sprintf("Hallo %s. Schön dass Du hier bist", user.Value)), nil

}

func handleUnknown(ctx context.Context, event Request) (interface{}, error) {
	return responseBuilder().speak(SPEECH_UNKNOWN + " " + SPEECH_START_TRAINING).
		reprompt(SPEECH_EXIT_IF_MUTE), nil
}

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

	case LAUNCH_REQUEST:

		entry, _ := db.GetLastUsedEntry(event.Session.User.UserID)
		if entry == nil {
			return responseBuilder().
				speak(SPEECH_WELCOME + SPEECH_DEFINE_USER).
				reprompt(SPEECH_START_TRAINING + SPEECH_EXIT_IF_MUTE), nil

		}

		return responseBuilder().
			speak(SPEECH_WELCOME + fmt.Sprintf(SPEECH_PERSONAL, entry.UserName) + SPEECH_START_TRAINING).
			reprompt(SPEECH_START_TRAINING + SPEECH_EXIT_IF_MUTE), nil

		// return responseBuilder().speak(timeText(4*60 + 30)), nil

	case SESSION_END:
		log.Println("End-Reason: ", event.RequestBody.Reason)
		return nil, nil

	case AUDIOPLAYER_STARTED, AUDIOPLAYER_NEARLYFINISHED, AUDIOPLAYER_FINISHED:
		log.Println("Audioplayer Request")
		return nil, nil

	case INTENT_REQUEST:
		switch event.RequestBody.Intent.Name {
		case STOP_INTENT:
			return responseBuilder().speak(SPEECH_ENDE).withShouldEndSession(), nil
		case HELP_INTENT:
			return responseBuilder().speak(SPEECH_START_TRAINING), nil
		case AUDIO_TEST:
			/* die Audioausgabe läuft parallel zum Dialog.
			D.h. wenn im Dialog nix mehr passiert, wird der Skill beendet */
			return responseBuilder().
				addAudioPlayerPlayDirective(
					"https://github.com/gipde/bodyweight/raw/master/contrib/alien-spaceship_daniel_simion.mp3"), nil
		case START_TRAINING:
			return handleStartTraining(ctx, event)
		case DEFINE_USER:
			return defineUser(ctx, event)
		case FALLBACK_INTENT:
			return handleUnknown(ctx, event)
		default:
			return handleUnknown(ctx, event)
		}
	}
	return handleUnknown(ctx, event)
}

var db database.DB

func init() {
	if !debug {
		log.SetOutput(ioutil.Discard)
	}
	db = database.NewDynamoDB("bodyweight")
}

func main() {
	if isClient() {
		connectToServer()
	}

	log.Println("Starting FN")
	lambda.Start(HandleRequest)
}
