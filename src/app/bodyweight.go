package app

import (
	"bodyweight/tools"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"bodyweight/database"
	"bodyweight/training"

	"github.com/aws/aws-lambda-go/lambda"
)

/* TODO:
Alexa-Return warum immer nil
*/

const (
	// Launch Intent
	speechWelcome    = `Willkommen beim <lang xml:lang="en-US">Bodyweight Training</lang>. `
	speechDefineUser = `Du benutzt das <lang xml:lang="en-US">Bodyweight Training</lang> zum ersten Mal. 
	Du solltest zunächst deinen Namen festlegen. Sage hierzu bitte <break time="500ms"/>mein Name ist
	<break time="500ms"/> und deinen Vornamen. `
	speechWelcomNewUser = `Hallo %s. Schön dass Du mit mir trainierst. `

	speechExplainTraining     = `Lass Dir zunächst das Training erklären. `
	speechExplainThisExercise = `Lass Dir diese nächste Übung erklären. `
	speechExplainNextExercise = `Lass Dir die nächste Übung erklären. `
	speechStart               = `Sage: <break strength="x-strong"/>bereit,<break strength="x-strong"/> `
	speechStartThisExercise   = speechStart + `wenn du mit dieser Übung beginnen möchtest. `
	speechDone                = `Geschafft. `
	speechReadyForToday       = `Jetzt kannst Du Dich erholen. Bis zum nächsten mal. `

	speechUnknown    = "Ich kann dich leider nicht verstehen. "
	speechExitIfMute = "Wenn Du nichts mehr sagst, wird das Programm beendet. "
	speechEnde       = "Auf Wiedersehen und bis bald. "
	speechPersonal   = `%s, es ist schön, dass du wieder da bist. `

	speechStufenintervall = `Jede Übung in diesem Block wird 4 einhalb bis 7 einhalb Minuten ausgeführt.
	Du beginnst mit einer Wiederholung, machst eine Pause, machst zwei Wiederholungen und so weiter.
	Ab der Hälft der Zeit reduzierst Du die Sätze jeweils um eine Wiederholung. Wenn Du bereits
	vorher nicht mehr kannst, kannst Du auch früher schon reduzieren. Ist noch Zeit übrig,
	beginnst du mit einer neuen Stufe wieder mit zunächst einer Wiederholung`

	speechIntervallsatz = `Beim Intervallsatz sind von jeder Übung 3 Sätze mit 6-12 Wiederholungen
	durchzuführen. Bei einem Satz solltest Du eigentlich bis zum Muskelversagen kommen. Für einen Satz hast Du
	genau 3 Minuten zeit. Falls Du früher fertig wirst, kannst Du die Zeit als Pause nutzen. Die Sätze beginnen damit
	immer genau im Abstand von 3 Minuten. Wechsle bei einseitigen Übungen nach jedem Satz die Seite,es sei denn die
	Wiederholungen sind im Wechsel durchzuführen. Du beginnst zunächst mit der schwächeren Seite`

	speechSupersatz = `Ein Supersatz besteht aus einem Übungspaar und dauert 4 Minuten. Bei der ersten Übung
	sind jeweils 1 - 5 Wiederholungen, bei der zweiten Übung sind 6-12 Wiederholungen zu absolvieren. Pro Paar sind 2
	Supersätze direkt nacheinander durchzuführen. Bei einseitigen Übungen wechselst Du nach jeder Wiederholung die Seite.`

	speechHochintensitaetssatz = `Beim Hochintensitätssatz sind insgesamt 8 Sätze mit jeweils 20 Sekunden Training
	gefolgt von je 10 Sekunden Pause durchzuführen. Insgesamt dauert jede Übung 4 Minuten lang.`

	speechZirkelintervall = `Das Zirkelintervall besteht aus insgesamt 3 verschiedenen Übungen, von denen jeweils
	eine angegebene Anzahl von Wiederholungen durchzuführen ist. Ohne Pause führen sie die Übungen im Wechsel durch.
	Versuchen Sie das Zirkelintervall insgesamt 20 Minuten durchzuführen.`

	// TODO: hört sich noch nicht gut an
	speechHelp = `Du brauchst Hilfe?
	
	In diesem Skill kannst Du mit:
	- erkläre das Training - Dir das Training erklären lassen.
	- erkläre die nächste Übung - dir die nächste Übung erklären lassen.
	- ändere den Trainingsfortschritt - den Trainingsfortschritt ändern.
	- bereit - die nächste Übung starten oder mit
	- erkläre Trainingsmethode  - die jeweiligen Trainingsmethoden erklären.`

	textTitle           = "Bodyweight Training"
	textWelcome         = "Herzlich willkommen"
	textExplainTraining = "Das heutige Training:"
	textExplainExercise = "Die nächste Übung:"
	textDefineName      = textWelcome + ", du musst zunächst deinen Namen festlegen."
)

var db database.DB

func init() {
	if !tools.IsDebug() {
		log.SetOutput(ioutil.Discard)
	}
}

// Start starts the Lambda Handler
func Start() {

	// start as Client
	if tools.IsTestClient() {
		tools.ConnectToServer()
		os.Exit(0)
	}

	// start as Server
	log.Println("Starting FN")
	db = database.Accessor()
	lambda.Start(HandleRequest)
}

// HandleRequest handles every Request / Base-Entry FN
func HandleRequest(ctx context.Context, event Request) (interface{}, error) {

	if event.RequestBody.Type == "IntentRequest" {
		log.Printf("Intent: %+v\n", event.RequestBody.Intent.Name)
	} else {
		log.Printf("Request: %+v\n", event.RequestBody.Type)
	}

	if isDelegate() {
		log.Println("we delegate")
		result, err := delegateRemote(event)
		log.Println("and got result: ")
		return result, err
	}

	user := getUser(event)
	switch event.RequestBody.Type {

	case alexaLaunchRequest:
		if user == nil {
			// First Usage

			return responseBuilder().
				speak(speechWelcome+speechDefineUser).
				withSimpleCard(textTitle, textDefineName).
				reprompt(speechDefineUser + speechExitIfMute), nil

		}

		// Welcome Back
		return responseBuilder().
			speak(speechWelcome+fmt.Sprintf(speechPersonal, user.UserName)+speechExplainTraining).
			withSimpleCard(textWelcome, user.UserName).
			// withSimpleCard(user.TrainingState.ShortProgress(), user.TrainingState.CardDayDescription()).
			reprompt(speechExplainTraining + speechExitIfMute), nil

	case alexaSessionEndRequest:
		log.Println("End-Reason: ", event.RequestBody.Reason)
		return nil, nil

	case alexaAudioplayerStartedRequest, alexaAudioplayerNearlyFinishedRequest, alexaAudioplayerFinishedRequest:
		log.Println("Audioplayer Request")
		return nil, nil

	case alexaIntentRequest:

		// intents without user
		switch event.RequestBody.Intent.Name {
		case alexaDefineUserIntent:
			return defineUser(event)
		case alexaExplainTrainingMethodIntent:
			return handleExplainTrainingMethod(event), nil

		case alexaStopIntent:
			return responseBuilder().speak(speechEnde).withShouldEndSession(), nil
		case alexaHelpIntent:
			return responseBuilder().speak(speechHelp).addAudioPlayerPlayDirective(
				"https://github.com/gipde/bodyweight/raw/master/contrib/audio1.mp3"), nil
		case alexaAudioTestIntent:
			/* die Audioausgabe läuft parallel zum Dialog.
			D.h. wenn im Dialog nix mehr passiert, wird der Skill beendet */
			return responseBuilder().
				addAudioPlayerPlayDirective(
					"https://github.com/gipde/bodyweight/raw/master/contrib/audio1.mp3"), nil

		case alexaFallbackIntent:
			return handleUnknown()
		}

		// intents which require user
		if user == nil {
			return responseBuilder().speak(speechDefineUser).reprompt(speechDefineUser + speechExitIfMute), nil
		}
		switch event.RequestBody.Intent.Name {
		case alexaExplainTrainingIntent:
			return responseBuilder().speak(user.TrainingState.ExplainTraining()+speechExplainThisExercise).
				withSimpleCard(textExplainTraining, user.TrainingState.CardDayDescription()).
				reprompt(speechExplainThisExercise + speechExitIfMute), nil
		case alexaExplainExerciseIntent:
			return responseBuilder().speak(user.TrainingState.ExplainExercise()+speechStartThisExercise).
				withSimpleCard(textExplainExercise, user.TrainingState.CardUnitDescription()).
				reprompt(speechStartThisExercise + speechExitIfMute), nil
		case alexaStartTrainingIntent:
			return handleStartTraining(user), nil
		}
	}
	return handleUnknown()
}

func handleExplainTrainingMethod(event Request) *Response {
	method := event.RequestBody.Intent.Slots["Method"]
	resolutions := method.Resolutions.ResolutionsPerAuthority

	id := -1
	if resolutions != nil {
		r1 := resolutions[0]
		if r1.Status.Code == "ER_SUCCESS_MATCH" {
			id, _ = strconv.Atoi(r1.Values[0]["value"].ID)
		}
	}

	var text string
	if id == -1 {
		text = "diese Trainingsmethode kenne ich nicht."
	}

	switch training.TrainingMethod(id) {
	case training.StufenIntervall:
		text = speechStufenintervall
	case training.IntervallSatz:
		text = speechIntervallsatz
	case training.SuperSatz:
		text = speechSupersatz
	case training.HochIntensitaetsSatz:
		text = speechHochintensitaetssatz
	case training.ZirkelIntervall:
		text = speechZirkelintervall
	}
	return responseBuilder().speak(text).reprompt(speechExitIfMute)
}

func handleStartTraining(user *database.Entry) *Response {
	//TODO: check if is directyl called

	desc := user.TrainingState.CardUnitDescription()
	text := user.TrainingState.InstructTraining()

	defer func() {
		user.Date = time.Now()
		db.CreateEntry(user)
	}()

	if user.TrainingState.WasLastUnit() {
		return responseBuilder().
			withSimpleCard(textExplainExercise, desc).
			speak(text + speechDone + speechReadyForToday).
			withShouldEndSession()
	}
	return responseBuilder().
		withSimpleCard(textExplainExercise, desc).
		speak(text + speechDone + speechExplainNextExercise)

}

func handleUnknown() (interface{}, error) {
	return responseBuilder().speak(speechUnknown + " " + speechExplainTraining).
		reprompt(speechExitIfMute), nil
}

func getNewUserEntry(userid, username string) *database.Entry {
	return &database.Entry{
		PK: database.PK{
			AlexaID: userid,
			Date:    time.Now(),
		},
		TrainingState: training.GetBeginningState(),
		UserName:      username,
	}
}

func defineUser(event Request) (interface{}, error) {
	user := event.RequestBody.Intent.Slots["user"]
	userEntry := *getNewUserEntry(event.Session.User.UserID, user.Value)

	// we check, if this user already exists
	entries, _ := db.GetEntries(event.Session.User.UserID)
	if entries != nil {
		for _, entry := range *entries {
			if entry.UserName == user.Value {
				// set userEntry
				userEntry = entry
				userEntry.Date = time.Now()

				// delete old entry
				db.DeleteItem(entry.PK)
			}
		}
	}

	db.CreateEntry(&userEntry)

	return responseBuilder().
		speak(fmt.Sprintf(speechWelcomNewUser, user.Value)+speechExplainTraining).
		withSimpleCard(textWelcome, user.Value).
		reprompt(speechExplainTraining + speechExitIfMute), nil

}

func getUser(event Request) *database.Entry {
	uid := event.Session.User.UserID // Amazon UID
	entry, _ := db.GetLastUsedEntry(uid)
	return entry
}
