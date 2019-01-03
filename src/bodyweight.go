package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func getBaseResponse() Response {
	return Response{
		Version: "1.0",
		ResponseBody: ResponseBody{
			ShouldEndSession: false,
		},
	}
}

func EndSession(text string) Response {
	r := AnswerText(text)
	r.ResponseBody.ShouldEndSession = true
	return r
}

func AnswerText(text string) Response {
	return getBaseResponse().withText(text)
}

func (r Response) dialogDelegate(userOk bool) Response {
	// var confirmState string
	var value string
	// if userOk {
	// 	confirmState = "CONFIRMED"
	// 	value = "hans"
	// } else {
	// 	confirmState = "NONE"
	// }
	// log.Println(confirmState)
	m := make(map[string]Slot)
	m["user"] = Slot{
		Name:  "user",
		Value: value,
		// ConfirmationStatus: "NONE",
	}

	r.ResponseBody.Directives = []Directive{
		Directive{
			Type: "Dialog.Delegate",
			UpdatedIntent: &Intent{
				Name: "StartTraining",
				// ConfirmationStatus: "NONE",
				Slots: m,
			},
		}}

	return r
}

func (r Response) audioInterface(url string) Response {
	r.ResponseBody.Directives = []Directive{
		Directive{
			Type:         "AudioPlayer.Play",
			PlayBehavior: "REPLACE_ALL",
			AudioItem: &AudioItem{
				Stream: &Stream{
					URL:                  "http://soundbible.com/grab.php?id=2218&type=mp3",
					Token:                "234234",
					OffsetInMilliseconds: 0,
				},
			},
		},
	}
	return r
}

func (r Response) withCard() Response {
	r.ResponseBody.Card = &Card{
		Type:  "Simple",
		Title: "Bodyweight Training",
		Text:  "SimpleCard",
	}
	return r
}

func (r Response) withText(text string) Response {
	r.ResponseBody.OutputSpeech = &OutputSpeech{
		Type: "PlainText",
		Text: text,
	}
	return r
}
func (r Response) withReprompt(text string) Response {
	r.ResponseBody.Reprompt = &Reprompt{
		OutputSpeech: &OutputSpeech{
			Type: "PlainText",
			Text: text,
		},
	}
	return r
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
		log.Println("a launch Request")
		return AnswerText("Willkommen beim Bodyweight Training. Du kannst die Trainings mit \"leg los\" starten."), nil
	case SESSION_END:
		return EndSession("du willst beenden, gerne"), nil

	case INTENT_REQUEST:
		switch event.RequestBody.Intent.Name {
		case STOP_INTENT:
			return EndSession("Auf Wiedersehen und bis bald"), nil
		case HELP_INTENT:
			return AnswerText("Du kannst ein Training mit starte Training starten"), nil
		case AUDIO_TEST:
			return getBaseResponse().audioInterface("http://soundbible.com/grab.php?id=2218&type=mp3"), nil
		case START_TRAINING:
			if event.RequestBody.DialogState != "COMPLETED" {
				// das kann man sich dann wohl sparen
				userOK := event.RequestBody.Intent.Slots["user"].Value != ""
				return getBaseResponse().dialogDelegate(userOK), nil
			} else {
				return AnswerText("geschafft"), nil
			}
		default:
			return AnswerText("ich kann keinen passenden Intent finden"), nil
		}

	}
	return AnswerText("das war wohl nix"), nil
}

func main() {
	if isClient() {
		connectToServer()
	}

	log.Println("Starting FN")
	lambda.Start(HandleRequest)
}
