package tools

import (
	"github.com/aws/aws-lambda-go/lambda/messages"
	"log"
	"math/rand"
	"net/rpc"

	"os"
	"strconv"
	"time"
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

/*
Connect as Client
*/

var debug=false

func IsDebug() bool {
	return debug
}

func SetDebug() {
	debug=true
}

func ConnectToServer() {
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

func IsTestClient() bool {
	return os.Getenv("TESTCLIENT") != ""
}

func SetupTestDB() {
	rand.Seed(time.Now().UTC().UnixNano())
	r := strconv.Itoa(rand.Int())
	tablename := "TESTTABLE_" + r

	os.Setenv("TEST_DB", tablename)
	log.Printf("Setup TEST DB Environment")
}