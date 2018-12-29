package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

func HandleRequest(ctx context.Context, event interface{}) (interface{}, error) {
	log.Printf("Context: %v\n", ctx)
	log.Printf("Event:   %v\n", event)
	var f interface{}
	err := json.Unmarshal(result, &f)
	return f, err
}

type Item struct {
	ID    string `json:"id"`
	State int    `json:"state"`
}

func readDynamoDB() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("bodyweight"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String("state"),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("id:  ", item.ID)
	fmt.Println("State: ", item.State)

}

func main() {
	readDynamoDB()
	lambda.Start(HandleRequest)
}
