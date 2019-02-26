package app

func responseBuilder() *Response {
	sessionAttributes := make(map[string]interface{})
	return &Response{
		Version:           "1.0",
		SessionAttributes: sessionAttributes,
		ResponseBody:      ResponseBody{},
	}
}

func (r *Response) withShouldEndSession() *Response {
	r.ResponseBody.ShouldEndSession = true
	return r
}

// audio is playing in background
func (r *Response) addAudioPlayerPlayDirective(url string) *Response {
	r.ResponseBody.Directives = []Directive{
		{
			Type:         "AudioPlayer.Play",
			PlayBehavior: "REPLACE_ALL",
			AudioItem: &AudioItem{
				Stream: &Stream{
					URL:                  url,
					Token:                "234234",
					OffsetInMilliseconds: 0,
				},
			},
		},
	}
	return r
}

func (r *Response) withSimpleCard(title, text string) *Response {
	r.ResponseBody.Card = &Card{
		Type:    "Simple",
		Title:   title,
		Content: text,
	}
	return r
}

func (r *Response) speak(payload string) *Response {
	r.ResponseBody.OutputSpeech = &OutputSpeech{
		Type: "SSML",
		SSML: "<speak>" + payload + "</speak>",
	}
	return r
}


// provide prompt speech to the user if no response for 8 seconds
func (r *Response) reprompt(text string) *Response {

	r.ResponseBody.Reprompt = &Reprompt{
		OutputSpeech: &OutputSpeech{
			Type: "SSML",
			SSML: "<speak>" + text + "</speak>",
		},
	}
	return r
}

// func (r *Response) addDelegateDirective(intent *Intent, confirmed bool) *Response {
// 	r.ResponseBody.Directives = []Directive{
// 		{
// 			Type:          "Dialog.Delegate",
// 			UpdatedIntent: intent,
// 		}}

// 	if confirmed {
// 		intent.ConfirmationStatus = "CONFIRMED"
// 	}
// 	return r
// }
