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
			OutputSpeech: &OutputSpeech{
				Type: "PlainText",
				Text: "",
			},
			Card: &Card{
				Type:  "Simple",
				Title: "Bodyweight Training",
				Text:  "SimpleCard",
			},
			Reprompt: &Reprompt{
				OutputSpeech: &OutputSpeech{
					Type: "PlainText",
					Text: "",
				},
			},
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
	r := getBaseResponse()
	r.ResponseBody.OutputSpeech.Text = text
	r.ResponseBody.Reprompt.OutputSpeech.Text = text
	return r
}

func (r Response) delegate() Response {
	m := make(map[string]Slot)
	m["user"] = Slot{
		Name: "user",
		//		Value:              "Vera",
		ConfirmationStatus: "NONE",
	}
	r.ResponseBody.Directives = []Directive{
		Directive{
			Type: "Dialog.Delegate",
			UpdatedIntent: Intent{
				Name:               "StartTraining",
				ConfirmationStatus: "NONE",
				Slots:              m,
			},
		}}
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
		return AnswerText("Willkommen beim Bodyweight Training").delegate(), nil
	case INTENT_REQUEST:
		switch event.RequestBody.Intent.Name {
		case STOP_INTENT:
			return EndSession("Auf Wiedersehen und bis bald"), nil
		case HELP_INTENT:
			return AnswerText("ach du brauchst Hilfe?"), nil
		case START_TRAINING:
			log.Println("Starte als User: ", event.RequestBody.Intent.Slots["user"].Value)
			return AnswerText("ich freue mich").delegate(), nil
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
