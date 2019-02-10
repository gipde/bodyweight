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


