package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event Request) (Response, error) {
	log.Printf("Request: %+v\n", event)

	switch event.RequestBody.Type {
	case LAUNCH_REQUEST:
		return BuildTextResponse("Hallo Miriam"), nil
	case INTENT_REQUEST:
		name := event.RequestBody.X["intent"].(map[string]interface{})["name"]
		if name == STOP_INTENT {
			return BuildTextResponse("ah ein Intent"), nil
		}
	}
	return BuildTextResponse("das war wohl nix"), nil
}

func main() {
	log.Println("Starting FN")
	lambda.Start(HandleRequest)
}
