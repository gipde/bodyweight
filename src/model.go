package main

const LAUNCH_REQUEST = "LaunchRequest"
const INTENT_REQUEST = "IntentRequest"
const STOP_INTENT = "AMAZON.StopIntent"
const HELP_INTENT = "AMAZON.HelpIntent"
const FALLBACK_INTENT = "AMAZON.FallbackIntent"
const START_TRAINING = "StartTraining"
const START_EXERCISE = "StartExercise"
const AUDIO_TEST = "AudioTest"
const SESSION_END = "SessionEndedRequest"
const AUDIOPLAYER_STARTED = "AudioPlayer.PlaybackStarted"
const AUDIOPLAYER_NEARLYFINISHED = "AudioPlayer.PlaybackNearlyFinished"
const AUDIOPLAYER_FINISHED = "AudioPlayer.PlaybackFinished"

type Request struct {
	Version     string      `json:"version"`
	Session     Session     `json:"session"`
	Context     Context     `json:"context"`
	RequestBody RequestBody `json:"request"`
}

type RequestBody struct {
	Type        string `json:"type"`
	RequestID   string `json:"requestId"`
	Timestamp   string `json:"timestamp"`
	DialogState string `json:"dialogState"`
	Locale      string `json:"locale"`
	Intent      Intent `json:"intent"`
	Reason      string `json:"reason"`
	Error       Error  `json:"error"`
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Intent struct {
	Name               string          `json:"name"`
	ConfirmationStatus string          `json:"confirmationStatus,omitempty"`
	Slots              map[string]Slot `json:"slots,omitEmpty"`
}

type Slot struct {
	Name               string      `json:"name"`
	Value              string      `json:"value,omitempty"`
	ConfirmationStatus string      `json:"confirmationStatus,omitempty"`
	Resolutions        Resolutions `json:"resolutions,omitempty"`
}

type Resolutions struct {
	ResolutionsPerAuthority []ResolutionsPerAuthority `json:"resolutionsPerAuthority"`
}

type ResolutionsPerAuthority struct {
	Authority string        `json:"authority"`
	Status    Status        `json:"status"`
	Values    []map[string]Value `json:"values"`
}

type Status struct {
	Code string `json:"code"`
}

type Value struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Context struct {
	System      System      `json:"system"`
	AudioPlayer interface{} `json:"AudioPlayer"`
}

type Session struct {
	New         bool                   `json:"new"`
	SessionID   string                 `json:"sessionId"`
	Application Application            `json:"application"`
	Attributes  map[string]interface{} `json:"attributes"`
	User        User                   `json:"user"`
}

type System struct {
	APIAccessToken string      `json:"apiAccessToken"`
	APIEndpoint    string      `json:"apiEndpoint"`
	Application    Application `json:"application"`
	Device         interface{} `json:"device"`
	User           User        `json:"user"`
}

type User struct {
	UserID      string     `json:"userId"`
	AccessToken string     `json:"accessToken"`
	Permissions Permission `json:"permission"`
}

type Permission struct {
	ConsentToken string `json:"consentToken"`
}

type Application struct {
	ApplicationID string `json:"applicationId"`
}

type OutputSpeech struct {
	Type         string `json:"type"`
	Text         string `json:"text"`
	SSML         string `json:"ssml,omitempty"`
	PlayBehavior string `json:"playBehavior,omitempty"`
}

type Image struct {
	SmallImageURL string `json:"smallImageUrl"`
	LargeImageURL string `json:"lageImageUrl"`
}
type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content,omitempty"`
	Text    string `json:"text,omitempty"`
	Image   *Image `json:"image,omitempty"`
}

type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	ResponseBody      ResponseBody           `json:"response"`
}

type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}

type ResponseBody struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	Directives       []Directive   `json:"directives,omitempty"`
	ShouldEndSession bool          `json:"shouldEndSession"`
}

type Stream struct {
	URL                   string `json:"url"`
	Token                 string `json:"token"`
	ExpectedPreviousToken string `json:"expectedPreviousToken,omitempty"`
	OffsetInMilliseconds  int    `json:"offsetInMilliseconds"`
}
type AudioItem struct {
	Stream *Stream `json:"stream,omitempty"`
}

type Directive struct {
	Type          string     `json:"type"`
	PlayBehavior  string     `json:"playBehavior,omitempty"`
	SlotToElicit  string     `json:"slotToElicit,omitempty"`
	SlotToConfirm string     `json:"slotToConfirm,omitempty"`
	UpdatedIntent *Intent    `json:"updatedIntent,omitempty"`
	AudioItem     *AudioItem `json:"audioItem,omitempt"`
}

type DBItem struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
