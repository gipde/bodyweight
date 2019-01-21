package app

import (
	"bodyweight/database"
	"bodyweight/tools"
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
	m := triggerLaunchRequest()

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
	m := triggerDefineUserIntent(name)

	assert.Equal(t, "SSML", m.OutputSpeech.Type)
	assert.False(t, m.ShouldEndSession)
	assert.True(t, strings.Contains(m.OutputSpeech.SSML, name))

	// trigger Launch Request
	m = triggerLaunchRequest()

	assert.Equal(t, m.OutputSpeech.SSML,
		speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name)+speechStartTraining))
	assert.False(t, m.ShouldEndSession)
}

func TestChangeToNewUser(t *testing.T) {
	name1 := "Bob"
	name2 := "Alice"

	m := triggerDefineUserIntent(name2)
	m = triggerLaunchRequest()
	assert.Equal(t, m.OutputSpeech.SSML,
		speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name2)+speechStartTraining))

	m = triggerDefineUserIntent(name1)
	m = triggerLaunchRequest()
	assert.Equal(t, m.OutputSpeech.SSML,
		speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name1)+speechStartTraining))

	entries, _ := db.GetEntries(amazonTestUser)
	assert.Equal(t, 2, len(*entries))
	assert.Equal(t, name2, (*entries)[0].UserName)
	assert.Equal(t, name1, (*entries)[1].UserName)
}

func TestLaunchHandlerKnownUser(t *testing.T) {
	name1 := "Bob"
	m := triggerLaunchRequest()
	assert.Equal(t, speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name1)+speechStartTraining),
		m.OutputSpeech.SSML)
}

func TestStartTraining(t *testing.T) {

}

func triggerLaunchRequest() *ResponseBody {
	r, _ := HandleRequest(nil, BuildRequest(alexaLaunchRequest, nil))
	m := &r.(*Response).ResponseBody
	return m
}

func triggerDefineUserIntent(name string) *ResponseBody {
	r, _ := HandleRequest(nil, BuildRequest(alexaIntentRequest, &Intent{
		Name: alexaDefineUserIntent,
		Slots: map[string]Slot{"user": {
			Value: name,
		}},
	}))
	m := r.(*Response).ResponseBody
	return &m
}
