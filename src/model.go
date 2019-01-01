package main

import "log"

const LAUNCH_REQUEST = "LaunchRequest"
const INTENT_REQUEST = "IntentRequest"
const STOP_INTENT = "AMAZON.StopIntent"

type Request struct {
	Version string `json:"version"`
	Session struct {
		New         bool
		SessionID   string            `json:"sessionId"`
		Application Application       `json:"application"`
		Attributes  map[string]string `json:"attributes"`
	} `json:"session"`
	Context     struct{}    `json:"context"`
	RequestBody RequestBody `json:"request"`
}

type RequestBody struct {
	Type string                 `json:"type"`
	X    map[string]interface{} `json:-`
}

// type Session struct {
// 	New         bool
// 	SessionID   string            `json:"sessionId"`
// 	Application Application       `json:"application"`
// 	Attributes  map[string]string `json:"attributes"`
// }

type User struct {
	UserID      string `json:"userId"`
	AccessTolen string `json:"accessToken"`
	Permissions interface{}
}

type Application struct {
	ApplicationID string `json:"applicationId"`
}

type OutputSpeech struct {
	Type          string `json:"type"`
	Text          string `json:"text"`
	SSML          string `json:"ssml,omitempty"`
	playBehaviour string `json:"playBehaviour,omitempty"`
}

type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Text    string `json:"text,omitempty"`
	Image   []byte `json:"image,omitempty"`
}

type Response struct {
	Version           string            `json:"version"`
	SessionAttributes map[string]string `json:"sessionAttributes,omitempty"`
	ResponseBody      ResponseBody      `json:"response"`
}

type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
}

type ResponseBody struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech,omitempty"`
	Card             Card         `json:"card,omitempty"`
	Reprompt         Reprompt     `json:"reprompt,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession,omitempty"`
}

func BuildTextResponse(text string) Response {
	log.Printf("We build a Textresponse with: %s\n", text)
	r := Response{
		Version: "1.0",
		ResponseBody: ResponseBody{
			OutputSpeech: OutputSpeech{
				Type: "PlainText",
				Text: text,
			},
		},
	}
	log.Printf("R: %+v", r)
	return r
}

func BuildRequest(requestType string) Request {
	return Request{
		RequestBody: RequestBody{
			Type: requestType,
		},
	}
}

type DBItem struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}