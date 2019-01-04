package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

const debug = true

const SPEECH_START_TRAINING = `
	Du kannst die Trainings mit 
	<break strength="x-strong"/>starte das training<break strength="x-strong"/>
	starten.`

const SPEECH_UNKNOWN = "Ich kann dich leider nicht verstehen."

const SPEECH_EXIT_IF_MUTE = "Wenn Du nichts mehr sagts, wird das Programm beendet."

const SPEECH_ENDE = "Auf Wiedersehen und bis bald."

const SPEECH_WELCOME = `"Willkommen beim <lang xml:lang="en-US">Bodyweight Training</lang>.`

func handleStartTraining(ctx context.Context, event Request) (interface{}, error) {

	user := event.RequestBody.Intent.Slots["user"]
	log.Printf("Session User: %+v", user.Value)
	if user.Value != "" {
		user.ConfirmationStatus = "CONFIRMED"
	}

	if event.RequestBody.DialogState != "COMPLETED" {
		return responseBuilder().addDelegateDirective(&event.RequestBody.Intent, false), nil
	}
	training := event.RequestBody.Intent.Slots["ex_number"].Resolutions.ResolutionsPerAuthority[0].Values[0]["value"].ID
	log.Printf("Training: %+v", training)

	//get zustand aus db für den User
	// zeige zustand auf (tag + )

	text := fmt.Sprintf("Herzlich Willkommen %s. Wir starten mit der ersten Übung. ", user) + timeText(4*60+30)
	return responseBuilder().speak(text).
		setSessionAttribute("user", user).
		reprompt("wenn du beginnen moechtest, sage starte die erste Übung"), nil

	//erkläre Übung
	// sind sie bereit?
	// starte Übung
	// spiele audio playback

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
		return responseBuilder().
			speak(SPEECH_WELCOME + SPEECH_START_TRAINING).
			reprompt(SPEECH_START_TRAINING + SPEECH_EXIT_IF_MUTE), nil

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
		case FALLBACK_INTENT:
			return handleUnknown(ctx, event)
		default:
			return handleUnknown(ctx, event)
		}
	}
	return handleUnknown(ctx, event)
}

func init() {
	if !debug {
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	if isClient() {
		connectToServer()
	}

	log.Println("Starting FN")
	lambda.Start(HandleRequest)
}
