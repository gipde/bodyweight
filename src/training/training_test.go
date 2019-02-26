package training

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortProgress(t *testing.T) {
	s := GetBeginningState()
	r := s.ShortProgress()
	assert.Exactly(t, "1. Woche, 1. Tag", r)
}

func TestGetShortInfo(t *testing.T) {
	s := GetBeginningState()
	unit := s.getExesD()[s.Unit]
	r := unit.getCardInfo()
	assert.Contains(t, r, "Liegestütz")
}

func TestCardDayDesc(t *testing.T) {
	s := GetBeginningState()
	for count := 0; count < 176; count++ {
		r := s.CardDayDescription()
		assert.Contains(t, r, "Trainingsmethode:")
		t.Log(r)

		// switch n-times if more units per day
		currentDay := s.Day
		for s.Day == currentDay {
			s.switchToNextTraining()
		}
	}
}

func TestUnitShortDesc(t *testing.T) {
	s := GetBeginningState()
	r := s.CardUnitDescription()
	assert.Contains(t, r, "Stufenintervalle")
	s.Week = 2
	r = s.CardUnitDescription()
	assert.Contains(t, r, "Intervallsätze")
	s.Week = 4
	r = s.CardUnitDescription()
	assert.Contains(t, r, "Supersatz")
}

func TestExplainTraining(t *testing.T) {
	s := GetBeginningState()
	r := s.ExplainTraining()
	assert.Contains(t, r, "Basisprogramm")
	assert.Contains(t, r, "1. Trainingswoche")
	assert.Contains(t, r, "1. Übungstag")
	s.Level = firstClass
	r = s.ExplainTraining()
	assert.Contains(t, r, "First Class")
}

func TestExplainExercise(t *testing.T) {
	s := GetBeginningState()
	t.Log(s.ExplainExercise())
	s.Week = 2
	t.Log(s.ExplainExercise())
	s.Week = 4
	t.Log(s.ExplainExercise())
}

func TestTimeAsString(t *testing.T) {
	assert.Equal(t, timeAsStr(0), "Ende")
	assert.Equal(t, timeAsStr(30), "30 Sekunden")
	assert.Equal(t, timeAsStr(97), "1 Minute und 37 Sekunden")
	assert.Equal(t, timeAsStr(127), "2 Minuten und 7 Sekunden")
}

func BenchmarkTimeAsString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		timeAsStr(n)
	}
}

func TestBreakFor(t *testing.T) {
	assert.Equal(t,`<break time="10ms"/> `, breakFor(10) )
	assert.Equal(t, `<break time="10000ms"/> `, breakFor(10000))
	assert.Equal(t, `<break time="10000ms"/> <break time="10000ms"/> `, breakFor(20000))
	assert.Equal(t, `<break time="10000ms"/> <break time="10000ms"/> <break time="10000ms"/> <break time="1234ms"/> `, breakFor(31234))
}

func TestCount(t *testing.T) {
	assert.Equal(t, countDown("test"), `Drei. Zwei. Eins. test`)
}

func TestIsLastTraining(t *testing.T) {
	state := GetBeginningState()

	state.switchToNextTraining()
	log.Println("nach 1. State:", state)
	assert.Equal(t, state.WasLastUnit(), false)

	state.switchToNextTraining()
	log.Println("nach 2. State:", state)
	assert.Equal(t, state.WasLastUnit(), false)

	state.switchToNextTraining()
	log.Println("nach 3. State:", state)
	assert.Equal(t, state.WasLastUnit(), false)

	state.switchToNextTraining()
	log.Println("nach 4. State:", state)
	assert.Equal(t, state.WasLastUnit(), true)

	state.switchToNextTraining()
	log.Println("nach 1. State:", state)
	assert.Equal(t, state.WasLastUnit(), false)
}

func TestSwitchNextUnit(t *testing.T) {
	state := GetBeginningState()
	for p := 0; p < 4; p++ {
		for w := 0; w < 10; w++ {
			for d := 0; d < len(trainings[w].TrainingDays); d++ {
				for u := 0; u < len(trainings[w].TrainingDays[d].Exercises[p]); u++ {
					assert.Equal(t, state, State{Level: trainingLevel(p), Week: w, Day: d, Unit: u})
					state.switchToNextTraining()
				}
			}
		}
	}
	assert.Equal(t, State{3, 6, 0, 0}, state)
}

func BenchmarkSwitchNextUnit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		state := GetBeginningState()
		state.switchToNextTraining()
	}
}

func TestInstructTrainingStufenintervall(t *testing.T) {
	s := State{Level: firstClass, Week: 0, Day: 3, Unit: 0}
	for i := 0; i < 4; i++ {
		t.Log(s.InstructTraining())
	}
}
func TestInstructTrainingIntervallsatz(t *testing.T) {
	s := State{Level: basisProgram, Week: 2, Day: 0, Unit: 0}
	t.Log(s.InstructTraining())
}
func TestInstructTrainingSuperSatz(t *testing.T) {
	s := State{Level: basisProgram, Week: 4, Day: 0, Unit: 0}
	t.Log(s.InstructTraining())
}
func TestHochintensitaetssatz(t *testing.T) {
	s := State{Level: basisProgram, Week: 7, Day: 0, Unit: 0}
	for i := 0; i < 3; i++ {
		t.Log(s.InstructTraining())
	}
}
func TestZirkelintervall(t *testing.T) {
	s := State{Level: firstClass, Week: 6, Day: 4, Unit: 0}
	t.Log(s.InstructTraining())
}

func TestCompleteTraining(t *testing.T) {

	s := GetBeginningState()
	for count := 0; count < 613; count++ {

		if s.Day == 0 && s.Unit == 0 {
			t.Log("-------------------------------------------------")
			t.Log(fmt.Sprintf("Week: %d", s.Week+1))
			t.Log("-------------------------------------------------")
		}
		if s.Unit == 0 {
			t.Log("-------------------------------------------------")
			t.Log(fmt.Sprintf("Day: %d", s.Day+1))
			t.Log("SPEAK::" + s.ExplainTraining())
			t.Log("-------------------------------------------------")
		}
		t.Log("+++++++++++++++++++++++++++++++++++++++++++++++++")
		t.Logf("%d State: %+v ", count, s)
		t.Log("+++++++++++++++++++++++++++++++++++++++++++++++++")
		t.Log("SPEAK::" + s.ExplainExercise())
		t.Log("SPEAK::" + s.InstructTraining())
		if s.WasLastUnit() {
			t.Log("end of day")
			t.Log("=================================================")
		}
	}
}

func Test_exercise_getUnitInfo(t *testing.T) {
	s := GetBeginningState()
	for i := 0; i < 100; i++ {
		t.Log(s)
		t.Log(s.CardUnitDescription())
		s.InstructTraining()
	}
}

func Test_exercise_getShortDayDesc(t *testing.T) {
	s := GetBeginningState()
	for i := 0; i < 100; i++ {
		t.Log(s.CardDayDescription())
		s.InstructTraining()
	}
}
