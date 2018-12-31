package main

type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Context Context `json:"context"`
}

type Context struct {
}

type Session struct {
	New         bool
	SessionID   string            `json:"sessionId"`
	Application Application       `json:"application"`
	Attributes  map[string]string `json:"attributes"`
}

type User struct {
	UserID      string `json:"userId"`
	AccessTolen string `json:"accessToken"`
	Permissions interface{}
}

type Application struct {
	ApplicationID string `json:"applicationId"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
}

type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	Card             Card         `json:"card"`
	Reprompt         Reprompt     `json:"reprompt"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

// type Request struct {
// 	Type string `json:"type"`
// 	RequestID string `json:"requestId"`
// }

type DBItem struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
