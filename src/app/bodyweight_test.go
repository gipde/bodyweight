package app

import (
	"bodyweight/database"
	"bodyweight/tools"

	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
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

	r := m.Run()
	err := db.DeleteDB()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(r)
}

func TestLaunchHandlerNewUser(t *testing.T) {
	m, _ := triggerLaunchRequest()

	// neuer, unbekannter User
	assert.True(t, m.Card.Type == "Simple")
	assert.True(t, m.OutputSpeech.Type == "SSML")
	assert.Equal(t, speakyfy(speechWelcome+speechDefineUser), m.OutputSpeech.SSML)
	assert.Equal(t, speakyfy(speechDefineUser+speechExitIfMute), m.Reprompt.OutputSpeech.SSML)
	assert.Equal(t, false, m.ShouldEndSession)
}

func TestStartTrainingWithoutUserSet(t *testing.T) {
	m, e := triggerStartTraining("")
	assert.Equal(t, speakyfy(speechDefineUser), m.OutputSpeech.SSML)
	assert.Equal(t, speakyfy(speechDefineUser+speechExitIfMute), m.Reprompt.OutputSpeech.SSML)
	assert.Nil(t, m.Directives)
	assert.Nil(t, e)
}

func TestLaunchWithDefineUser(t *testing.T) {
	name := "Bob"
	// trigger Define User Intent
	m, _ := triggerDefineUserIntent(name)

	assert.Equal(t, "SSML", m.OutputSpeech.Type)
	assert.False(t, m.ShouldEndSession)
	assert.Equal(t, speakyfy(fmt.Sprintf(speechWelcomNewUser, name)+speechExplainTraining), m.OutputSpeech.SSML)

	// trigger Launch Request
	m, _ = triggerLaunchRequest()

	assert.Equal(t, speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name)+speechExplainTraining), m.OutputSpeech.SSML)
	assert.False(t, m.ShouldEndSession)
}

func TestChangeToNewUser(t *testing.T) {
	//  we assume, that the only user is bob
	name1 := "Bob"
	name2 := "Alice"

	m, err := triggerDefineUserIntent(name2)
	if err != nil {
		assert.Fail(t, "Error:", err)
	}
	t.Logf("m: %+v", m.OutputSpeech.SSML)
	m, err = triggerLaunchRequest()
	if err != nil {
		assert.Fail(t, "Error:", err)
	}

	assert.Equal(t, speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name2)+speechExplainTraining),
		m.OutputSpeech.SSML)

	m, _ = triggerDefineUserIntent(name1)
	assert.Equal(t, speakyfy(fmt.Sprintf(speechWelcomNewUser+speechExplainTraining,name1)), m.OutputSpeech.SSML)

	m, _ = triggerLaunchRequest()
	assert.Equal(t, speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name1)+speechExplainTraining),
		m.OutputSpeech.SSML)

	entries, _ := db.GetEntries(amazonTestUser)
	assert.Equal(t, 2, len(*entries))
	assert.Equal(t, name2, (*entries)[0].UserName)
	assert.Equal(t, name1, (*entries)[1].UserName)
}

func TestLaunchHandlerKnownUser(t *testing.T) {
	name1 := "Bob"
	m, _ := triggerLaunchRequest()
	assert.Equal(t, speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name1)+speechExplainTraining),
		m.OutputSpeech.SSML)
}

// func TestStartTrainingWithUserSet(t *testing.T) {
// 	name := "Bob"
// 	m, _ := triggerStartTraining(name)
// 	assert.Equal(t, speakyfy(speechWelcome+fmt.Sprintf(speechPersonal, name)), m.OutputSpeech.SSML)
// }

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
