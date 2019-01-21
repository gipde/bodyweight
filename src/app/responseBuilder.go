package app

func responseBuilder() *Response {
	sessionAttributes := make(map[string]interface{})
	return &Response{
		Version:           "1.0",
		SessionAttributes: sessionAttributes,
		ResponseBody: ResponseBody{
			ShouldEndSession: false,
		},
	}
}

func (r *Response) withShouldEndSession() *Response {
	r.ResponseBody.ShouldEndSession = true
	return r
}

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

func (r *Response) withStandardCard() *Response {
	r.ResponseBody.Card = &Card{
		Type:  "Standard",
		Title: "Bodyweight Training",
		Text:  "Content",
		Image: &Image{
			SmallImageURL: "http://bla",
			LargeImageURL: "http://bla",
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

func (r *Request) getUser() string {
	return r.Session.User.UserID
}

func (r *Request) getSessionAttribute(key string) interface{} {
	return r.Session.Attributes[key]
}

func (r *Response) setSessionAttribute(key string, value interface{}) *Response {
	r.SessionAttributes[key] = value
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

func (r *Response) addDelegateDirective(intent *Intent, confirmed bool) *Response {
	r.ResponseBody.Directives = []Directive{
		{
			Type:          "Dialog.Delegate",
			UpdatedIntent: intent,
		}}

	if confirmed {
		intent.ConfirmationStatus = "CONFIRMED"
	}
	return r
}
