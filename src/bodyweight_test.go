package main

import (
	"encoding/json"
	"log"
	"testing"
)

func BuildRequest(requestType string, intent string) Request {
	r := Request{
		RequestBody: RequestBody{
			Type: requestType,
		},
	}
	if requestType == INTENT_REQUEST {
		r.RequestBody.Intent = Intent{
			Name: intent,
		}
	}
	return r
}
func TestHandler(t *testing.T) {
	r, _ := HandleRequest(nil, BuildRequest(LAUNCH_REQUEST, ""))
	log.Printf("Result:  %+v\n", r)
	json, _ := json.MarshalIndent(r, "", "  ")
	log.Println("Json: ", string(json))
}

func TestResponseJson(t *testing.T) {
	r, _ := HandleRequest(nil, BuildRequest(INTENT_REQUEST,START_TRAINING))
	log.Printf("Result:  %+v\n", r)
	json, _ := json.MarshalIndent(r, "", "  ")
	log.Println("Json: ", string(json))
}
