package main

import (
	"encoding/json"
	"log"
	"testing"
)

func BuildRequest(requestType string) Request {
	return Request{
		RequestBody: RequestBody{
			Type: requestType,
		},
	}
}
func TestHandler(t *testing.T) {
	r, _ := HandleRequest(nil, BuildRequest(LAUNCH_REQUEST))
	log.Printf("Result:  %+v\n", r)
	json, _ := json.MarshalIndent(r, "", "  ")
	log.Println("Json: ", string(json))
}

func TestResponseJson(t *testing.T) {

}
