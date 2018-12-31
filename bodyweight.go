package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event interface{}) (interface{}, error) {
	log.Println("Handler reached")
	log.Println(getValue("ff"))
	log.Println(getValue("state"))
	log.Printf("Context: %v\n", ctx)
	log.Printf("Event:   %v\n", event)
	return Response{
		OutputSpeech: OutputSpeech{
			Type: "PlainText",
			Text: "Welcome",
		},
		Card: Card{
			Type:    "Simple",
			Title:   "SessionSpeechlet - Welcome",
			Content: "SessionSpeechlet - Welcome tho the Alexa Skill Kit Sample",
		},
		Reprompt: Reprompt{
			OutputSpeech: OutputSpeech{
				Type: "PlainText",
				Text: "Welcome",
			},
		},
	}, nil
}

func main() {
	log.Println("Starting FN")
	lambda.Start(HandleRequest)
}
