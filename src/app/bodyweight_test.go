package app

import (
	"bodyweight/database"
	"bodyweight/tools"
	"bodyweight/training"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	amazonTestUser = "AmazonTestUser"
)

func speakyfy(r string) string {
	return "<speak>" + r + "</speak>"
}

func BuildRequest(requestType string, intent *Intent) Request {
	r := Request{
		RequestBody: RequestBody{
			Type: requestType,
		},
		Session: Session{
			User: User{
				UserID: amazonTestUser,
			},
		},
	}
	if intent != nil {
		r.RequestBody.Intent = *intent
	}
	return r
}

func TestMain(m *testing.M) {
	tools.SetupTestDB()
	db = database.Accessor()
	defer db.DeleteDB()

	m.Run()
}

func TestLaunchHandlerNewUser(t *testing.T) {
	m, _ := triggerLaunchRequest()

	// neuer, unbekannter User
	assert.True(t, m.Card.Type == "Simple")
	assert.True(t, m.OutputSpeech.Type == "SSML")
	assert.Equal(t, m.OutputSpeech.SSML, speakyfy(speechWelcome+speechDefineUser))
	assert.Equal(t, m.Reprompt.OutputSpeech.SSML, speakyfy(speechStartTraining+speechExitIfMute))
	assert.Equal(t, m.ShouldEndSession, false)
}

func TestLaunchWithDefineUser(t *testing.T) {
	name := "Bob"
	// trigger Define User Intent
	m, _ := triggerDefineUserIntent(name)

	assert.Equal(t, "SSML", m.OutputSpeech.Type)
	assert.False(t, m.ShouldEndSession)
	assert.True(t, strings.Contains(m.OutputSpeech.SSML, name))

	// trigger Launch Request
	m, _ = triggerLaunchRequest()

	assert.Equal(t, m.OutputSpeech.SSML,
		speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name)+speechStartTraining))
	assert.False(t, m.ShouldEndSession)
}

func TestChangeToNewUser(t *testing.T) {
	//  we assume, that the only user is bob
	name1 := "Bob"
	name2 := "Alice"

	m, _ := triggerDefineUserIntent(name2)
	m, _ = triggerLaunchRequest()
	assert.Equal(t, m.OutputSpeech.SSML,
		speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name2)+speechStartTraining))

	m, _ = triggerDefineUserIntent(name1)
	m, _ = triggerLaunchRequest()
	assert.Equal(t, m.OutputSpeech.SSML,
		speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name1)+speechStartTraining))

	entries, _ := db.GetEntries(amazonTestUser)
	assert.Equal(t, 2, len(*entries))
	assert.Equal(t, name2, (*entries)[0].UserName)
	assert.Equal(t, name1, (*entries)[1].UserName)
}

func TestLaunchHandlerKnownUser(t *testing.T) {
	name1 := "Bob"
	m, _ := triggerLaunchRequest()
	assert.Equal(t, speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name1)+speechStartTraining),
		m.OutputSpeech.SSML)
}

func TestStartTrainingWithoutUserSet(t *testing.T) {
	db.DeleteDB()
	db.CreateDBIfNotExists()
	m, e := triggerStartTraining("")
	assert.NotNil(t, m.Directives)
	assert.Equal(t, "Dialog.Delegate", m.Directives[0].Type)
	assert.Nil(t, e)

	m, e = triggerStartTraining("werner")
	state := training.GetBeginningState()
	assert.Equal(t, speakyfy(fmt.Sprintf("Herzlich Willkommen zurück %s. ", "werner")+
		training.AnnounceDailyTraining(&state)), m.OutputSpeech.SSML)
}

func triggerRequest(req Request) (*ResponseBody, error) {
	r, e := HandleRequest(nil, req)
	m := &r.(*Response).ResponseBody
	return m, e
}

func triggerLaunchRequest() (*ResponseBody, error) {
	return triggerRequest(BuildRequest(alexaLaunchRequest, nil))
}

func triggerDefineUserIntent(name string) (*ResponseBody, error) {
	return triggerRequest(BuildRequest(alexaIntentRequest, &Intent{
		Name: alexaDefineUserIntent,
		Slots: map[string]Slot{"user": {
			Value: name,
		}},
	}))
}

func triggerStartTraining(name string) (*ResponseBody, error) {
	i := &Intent{
		Name: alexaStartTrainingIntent,
	}
	if name != "" {
		i.Slots = map[string]Slot{"user": {
			Value: name,
		}}
	}
	return triggerRequest(BuildRequest(alexaIntentRequest, i))
}
