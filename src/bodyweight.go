package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

const SPEECH_START_TRAINING = `Du kannst die Trainings mit <break strength="x-strong"/> 
	leg los <break strength="x-strong"/>starten.`
const SPEECH_UNKNOWN = `Ich kann dich leider nicht verstehen.`
const SPEECH_EXIT_IF_MUTE = `Wenn Du nichts mehr sagts, wird das Programm beendet.`

func timeAsStr(sec int) string {
	var res string
	min := sec / 60
	s := sec % 60
	if min > 0 {
		res += fmt.Sprintf("%d Minuten", min)
	}
	if s > 0 {
		if min > 0 {
			res += " und "
		}
		res += fmt.Sprintf("%d Sekunden", s)
	}
	return res
}

func count(word string) string {
	var res string
	br := "<break time=\"500ms\"/>"
	for _, r := range []string{"drei", "zwei", "eins"} {
		res += fmt.Sprintf("%s%s", br, r)
	}
	return res + br + word
}

func breakFor(ms int) string {
	divider := 10000
	var res string
	br := `<break time="%dms"/>`
	for ms >= divider {
		ms -= divider
		res += fmt.Sprintf(br, divider)
	}
	if ms > 0 {
		res += fmt.Sprintf(br, ms)
	}
	return res
}

func addTimeInfo(sec, half, breakTime, breakTimeSub int) (string, int) {
	waitTime := timeAsStr(sec)
	if sec == half {
		waitTime += ". Du hast bereits die hälfte geschafft"
	}
	breakTimeSub += 700 // für Context Wechsel
	breakTime -= breakTimeSub
	res := fmt.Sprintf("%snoch %s.", breakFor(breakTime), waitTime)
	if len(waitTime) < 11 {
		breakTimeSub = 1080
	} else if len(waitTime) < 25 {
		breakTimeSub = 1350
	} else if len(waitTime) < 48 {
		breakTimeSub = 2450
	} else if len(waitTime) < 64 {
		breakTimeSub = 3500
	} else {
		breakTimeSub = 5233
	}
	return res, breakTimeSub
}

func timeText(sec int) string {
	res := "Du musst die Übung " + timeAsStr(sec) + " durchhalten."
	res += count("start")
	half := sec / 2
	count("start")
	var breakTimeSub int
	var breakTime int
	var line string
	intervall := 30
	for sec > intervall {
		sec--
		breakTime += 1000
		if sec%intervall == 0 || sec == half {
			line, breakTimeSub = addTimeInfo(sec, half, breakTime, breakTimeSub)
			res += line
			breakTime = 0
		}
	}
	res += breakFor(intervall*1000 - 4000 - breakTimeSub)
	res += count("stop")
	return res
}

func handleStartTraining(ctx context.Context, event Request) (interface{}, error) {
	if event.RequestBody.DialogState != "COMPLETED" {
		return responseBuilder().addDelegateDirective(&event.RequestBody.Intent, false), nil
	}
	user := event.RequestBody.Intent.Slots["user"].Value
	training := event.RequestBody.Intent.Slots["ex_number"].Resolutions.ResolutionsPerAuthority[0].Values[0]["value"].ID
	log.Printf("Session User: %+v", user)
	log.Printf("Training: %+v", training)

	//get zustand aus db für den User
	// zeige zustand auf (tag + )

	text := fmt.Sprintf("Herzlich Willkommen %s, wir starten mit der ersten Übung. ", user) + timeText(4*60+30)
	return responseBuilder().ssml(text).
		setSessionAttribute("user", user).
		reprompt("wenn du beginnen moechtest, sage starte die erste Übung"), nil

	//erkläre Übung
	// sind sie bereit?
	// starte Übung
	// spiele audio playback

}

func handleUnknown(ctx context.Context, event Request) (interface{}, error) {
	return responseBuilder().ssml(SPEECH_UNKNOWN + " " + SPEECH_START_TRAINING).
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
			ssml(`"Willkommen beim <lang xml:lang="en-US">Bodyweight Training</lang>. Du kannst die Trainings mit <break strength="x-strong"/> starte training <break strength="x-strong"/>starten."`).
			reprompt("start das Training mit \"starte training\""), nil
	case SESSION_END:
		log.Println("End-Reason: ", event.RequestBody.Reason)
		return nil, nil
	case AUDIOPLAYER_STARTED, AUDIOPLAYER_NEARLYFINISHED, AUDIOPLAYER_FINISHED:
		log.Println("Audioplayer Request")
		return nil, nil

	case INTENT_REQUEST:
		switch event.RequestBody.Intent.Name {
		case STOP_INTENT:
			return responseBuilder().speak("Auf Wiedersehen und bis bald").withShouldEndSession(), nil
		case HELP_INTENT:
			return responseBuilder().speak("Du kannst ein Training mit starte Training starten"), nil
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

func main() {
	if isClient() {
		connectToServer()
	}

	log.Println("Starting FN")
	lambda.Start(HandleRequest)
}
