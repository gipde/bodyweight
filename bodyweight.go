package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log.Print("Hello from Lambda")
	log.Printf("Context: %v\nEvent: %v\n", ctx, name)
	return fmt.Sprintf("Hello Dear %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
