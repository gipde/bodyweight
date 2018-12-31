package main

import (
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	r, _ := HandleRequest(nil, BuildRequest(LAUNCH_REQUEST))
	log.Println("Result: ", r)
}
