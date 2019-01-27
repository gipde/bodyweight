package app

const (
	sessionUser = "username"

	alexaLaunchRequest                    = "LaunchRequest"
	alexaIntentRequest                    = "IntentRequest"
	alexaSessionEndRequest                = "SessionEndedRequest"
	alexaAudioplayerStartedRequest        = "AudioPlayer.PlaybackStarted"
	alexaAudioplayerNearlyFinishedRequest = "AudioPlayer.PlaybackNearlyFinished"
	alexaAudioplayerFinishedRequest       = "AudioPlayer.PlaybackFinished"

	alexaStopIntent                  = "AMAZON.StopIntent"
	alexaHelpIntent                  = "AMAZON.HelpIntent"
	alexaFallbackIntent              = "AMAZON.FallbackIntent"

	alexaDefineUserIntent            = "DefineUser"
	alexaExplainTrainingIntent       = "ExplainTraining"
	alexaExplainExerciseIntent       = "ExplainExercise"
	alexaStartTrainingIntent         = "StartTraining"
	alexaExplainTrainingMethodIntent = "ExplainTrainingMethod"

	alexaAudioTestIntent             = "AudioTest"
)

// Request API Call Request
type Request struct {
	Version     string      `json:"version"`
	Session     Session     `json:"session"`
	Context     Context     `json:"context"`
	RequestBody RequestBody `json:"request"`
}

// RequestBody API Call
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

// Error API Call
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Intent API Call
type Intent struct {
	Name               string          `json:"name"`
	ConfirmationStatus string          `json:"confirmationStatus,omitempty"`
	Slots              map[string]Slot `json:"slots,omitEmpty"`
}

// Slot API Call
type Slot struct {
	Name               string      `json:"name"`
	Value              string      `json:"value,omitempty"`
	ConfirmationStatus string      `json:"confirmationStatus,omitempty"`
	Resolutions        Resolutions `json:"resolutions,omitempty"`
}

// Resolutions API Call
type Resolutions struct {
	ResolutionsPerAuthority []ResolutionsPerAuthority `json:"resolutionsPerAuthority"`
}

// ResolutionsPerAuthority API Call
type ResolutionsPerAuthority struct {
	Authority string             `json:"authority"`
	Status    Status             `json:"status"`
	Values    []map[string]Value `json:"values"`
}

// Status API Call
type Status struct {
	Code string `json:"code"`
}

// Value API Call
type Value struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Context API Call
type Context struct {
	System      System      `json:"system"`
	AudioPlayer interface{} `json:"AudioPlayer"`
}

// Session API Call
type Session struct {
	New         bool                   `json:"new"`
	SessionID   string                 `json:"sessionId"`
	Application Application            `json:"application"`
	Attributes  map[string]interface{} `json:"attributes"`
	User        User                   `json:"user"`
}

// System API Call
type System struct {
	APIAccessToken string      `json:"apiAccessToken"`
	APIEndpoint    string      `json:"apiEndpoint"`
	Application    Application `json:"application"`
	Device         interface{} `json:"device"`
	User           User        `json:"user"`
}

// User API Call
type User struct {
	UserID      string     `json:"userId"`
	AccessToken string     `json:"accessToken"`
	Permissions Permission `json:"permission"`
}

// Permission API Call
type Permission struct {
	ConsentToken string `json:"consentToken"`
}

// Application API Call
type Application struct {
	ApplicationID string `json:"applicationId"`
}

// OutputSpeech API Call
type OutputSpeech struct {
	Type         string `json:"type"`
	Text         string `json:"text"`
	SSML         string `json:"ssml,omitempty"`
	PlayBehavior string `json:"playBehavior,omitempty"`
}

// Image API Call
type Image struct {
	SmallImageURL string `json:"smallImageUrl"`
	LargeImageURL string `json:"lageImageUrl"`
}

// Card API Call
type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content,omitempty"`
	Text    string `json:"text,omitempty"`
	Image   *Image `json:"image,omitempty"`
}

// Response API Call
type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	ResponseBody      ResponseBody           `json:"response"`
}

// Reprompt API Call
type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}

// ResponseBody API Call
type ResponseBody struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	Directives       []Directive   `json:"directives,omitempty"`
	ShouldEndSession bool          `json:"shouldEndSession"`
}

// Stream API Call
type Stream struct {
	URL                   string `json:"url"`
	Token                 string `json:"token"`
	ExpectedPreviousToken string `json:"expectedPreviousToken,omitempty"`
	OffsetInMilliseconds  int    `json:"offsetInMilliseconds"`
}

// AudioItem API Call
type AudioItem struct {
	Stream *Stream `json:"stream,omitempty"`
}

// Directive API Call
type Directive struct {
	Type          string     `json:"type"`
	PlayBehavior  string     `json:"playBehavior,omitempty"`
	SlotToElicit  string     `json:"slotToElicit,omitempty"`
	SlotToConfirm string     `json:"slotToConfirm,omitempty"`
	UpdatedIntent *Intent    `json:"updatedIntent,omitempty"`
	AudioItem     *AudioItem `json:"audioItem,omitempt"`
}

// DBItem API Call
type DBItem struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
