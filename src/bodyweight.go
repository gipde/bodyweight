package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func AnswerLaunch(text string) Response {
	return Response{
		Version: "1.0",
		ResponseBody: &ResponseBody{
			OutputSpeech: &OutputSpeech{
				Type: "PlainText",
				Text: text,
			},
			Card: &Card{
				Type:  "Simple",
				Title: "Bodyweight Training",
				Text:  "SimpleCard",
			},
			Reprompt: &Reprompt{
				OutputSpeech: &OutputSpeech{
					Type: "PlainText",
					Text: text,
				},
			},
			ShouldEndSession: false,
		},
	}
}

func HandleRequest(ctx context.Context, event interface{}) (interface{}, error) {
	log.Printf("Request: %+v\n", event)

	if isDelegate() {
		log.Println("we delegate")
		result, err := delegateRemote(event)
		log.Println("and got result: ", result)
		return result, err

	}

	switch event.(Request).RequestBody.Type {
	case LAUNCH_REQUEST:
		return AnswerLaunch("Willkommen beim Bodyweight Training"), nil
	case INTENT_REQUEST:
		name := event.(Request).RequestBody.X["intent"].(map[string]interface{})["name"]
		if name == STOP_INTENT {
			return AnswerLaunch("ah ein Intent"), nil
		}
	}
	return AnswerLaunch("das war wohl nix"), nil
}

func main() {
	if isClient() {
		connectToServer()
	}

	log.Println("Starting FN")
	lambda.Start(HandleRequest)
}
