package app

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda/messages"
	"log"
	"net/rpc"
	"os"
)

func isDelegate() bool {
	return os.Getenv("DELEGATE") == "true"
}

// Delegate to Remote Bodyweight Instance
func delegateRemote(req interface{}) (Response, error) {
	var r = Response{}
	client, err := rpc.Dial("tcp", "yrnrwxodb39dwkmr.myfritz.net:1234")
	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
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
		return r, err
	}

	// unmarshal to response object
	err = json.Unmarshal(iResp.Payload, &r)
	return r, err
}
