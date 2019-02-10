package training

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

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
	assert.Equal(t, breakFor(10), `<break time="10ms"/>`)
	assert.Equal(t, breakFor(10000), `<break time="10000ms"/>`)
	assert.Equal(t, breakFor(20000), `<break time="10000ms"/><break time="10000ms"/>`)
	assert.Equal(t, breakFor(31234), `<break time="10000ms"/><break time="10000ms"/><break time="10000ms"/><break time="1234ms"/>`)
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
					if !(p == 3 && w == 9 && d == 4 && u == 2) {
						assert.Equal(t, false, state.switchToNextTraining())
					} else {
						assert.Equal(t, true, state.switchToNextTraining()) // end of alle trainings
					}
				}
			}
		}
	}
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
	t.Log("-------------------------------------------------")
	for s.Level == basisProgram {
		t.Log("State: ",s)
		if s.Day == 0 && s.Unit == 0 {
			t.Log(fmt.Sprintf("Week: %d", s.Week))
			t.Log("-------------------------------------------------")
		}
		if s.Unit == 0 {
			t.Log(fmt.Sprintf("Day: %d", s.Day))
			t.Log("-------------------------------------------------")
		}
		t.Log(s.InstructTraining())
		if s.WasLastUnit() {
			t.Log("end of day")
			t.Log("=================================================")
		}
	}
}
