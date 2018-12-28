package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

var result = []byte(`{
	"version":"1.0",
	"sessionAttributes":{
 
	},
	"response":{
	   "outputSpeech":{
		  "type":"PlainText",
		  "text":"Welcome to the Alexa Skills Kit sample. Please tell me your favorite color by saying, my favorite color is red"
	   },
	   "card":{
		  "type":"Simple",
		  "title":"SessionSpeechlet - Welcome",
		  "content":"SessionSpeechlet - Welcome to the Alexa Skills Kit sample. Please tell me your favorite color by saying, my favorite color is red"
	   },
	   "reprompt":{
		  "outputSpeech":{
			 "type":"PlainText",
			 "text":"Please tell me your favorite color by saying, my favorite color is red"
		  }
	   },
	   "shouldEndSession":false
	}
 }`)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (interface{}, error) {
	log.Printf("Context: %v\n\nEvent: %v\n", ctx, name)
	var f interface{}
	err := json.Unmarshal(result, &f)
	return f, err
}

func main() {
	// log.Printf("starting...")
	// var f interface{}
	// json.Unmarshal(result, &f)
	// log.Printf("%v", f)
	lambda.Start(HandleRequest)
}
