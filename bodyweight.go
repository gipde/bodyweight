package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

var result = `{
	"version": "1.0",
	"sessionAttributes": {},
	"response": {
	  "outputSpeech": {
		"type": "PlainText",
		"text": "Welcome to the Alexa Skills Kit sample. Please tell me your favorite color by saying, my favorite color is red"
	  },
	  "card": {
		"type": "Simple",
		"title": "SessionSpeechlet - Welcome",
		"content": "SessionSpeechlet - Welcome to the Alexa Skills Kit sample. Please tell me your favorite color by saying, my favorite color is red"
	  },
	  "reprompt": {
		"outputSpeech": {
		  "type": "PlainText",
		  "text": "Please tell me your favorite color by saying, my favorite color is red"
		}
	  },
	  "shouldEndSession": false
	}
  }`

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log.Printf("Context: %v\n\nEvent: %v\n", ctx, name)
	//	return fmt.Sprintf("Hello Dear %s!", name.Name), nil
	return result, nil
}

func main() {
	lambda.Start(HandleRequest)
}
