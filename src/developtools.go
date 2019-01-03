package main

import (
	"encoding/json"
	"log"
	"net/rpc"
	"os"

	"github.com/aws/aws-lambda-go/lambda/messages"
)

var launchRequest = `{
	"version": "1.0",
	"session": {
	  "new": true,
	  "sessionId": "amzn1.echo-api.session.123456789012",
	  "application": {
		"applicationId": "amzn1.ask.skill.987654321"
	  },
	  "user": {
		"userId": "amzn1.ask.account.testUser"
	  },
	  "attributes": {}
	},
	"context": {
	  "AudioPlayer": {
		"playerActivity": "IDLE"
	  },
	  "System": {
		"application": {
		  "applicationId": "amzn1.ask.skill.987654321"
		},
		"user": {
		  "userId": "amzn1.ask.account.testUser"
		},
		"device": {
		  "supportedInterfaces": {
			"AudioPlayer": {}
		  }
		}
	  }
	},
	"request": {
	  "type": "LaunchRequest",
	  "requestId": "amzn1.echo-api.request.1234",
	  "timestamp": "2016-10-27T18:21:44Z",
	  "locale": "en-US"
	}
  }`

func isDelegate() bool {
	return os.Getenv("DELEGATE") != ""
}

func isClient() bool {
	return os.Getenv("TESTCLIENT") != ""
}

/*
Connect as Client
*/
func connectToServer() {
	client, err := rpc.Dial("tcp", "192.168.178.30:1234")
	resp := new(messages.InvokeResponse)
	req := messages.InvokeRequest{
		Payload: []byte(launchRequest),
	}
	if err != nil {
		log.Println("Error: ", err)
	}
	defer client.Close()
	err = client.Call("Function.Invoke", req, resp)
	if err != nil {
		log.Println("Error: ", err)
	}
	log.Printf("Response.Payload %+v", string(resp.Payload))
	log.Printf("Response.Error %+v", resp.Error)
	os.Exit(0)
}

func delegateRemote(req interface{}) (Response, error) {
	client, err := rpc.Dial("tcp", "yrnrwxodb39dwkmr.myfritz.net:1234")
	defer client.Close()
	if err != nil {
		log.Println("Error callin Remote: ", err)
	}

	// marshall for remoting
	payload, err2 := json.Marshal(req)
	if err2 != nil {
		log.Println("Error Marshalling: ", err2)
	}
	iResp := new(messages.InvokeResponse)
	iReq := messages.InvokeRequest{
		Payload: []byte(payload),
	}
	err = client.Call("Function.Invoke", iReq, iResp)
	if err != nil {
		log.Println("Error: ", err)
	}

	// unmarshal to response object
	var r = Response{}
	json.Unmarshal(iResp.Payload, &r)

	return r, err
}
