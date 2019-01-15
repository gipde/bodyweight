package training

import (
	"testing"

	"gotest.tools/assert"
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
	assert.Equal(t, count("test"), `Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>test`)
}

func TestTimeText(t *testing.T) {
	assert.Equal(t, timeText(60), `Du musst die Übung 1 Minute durchhalten. Wir starten jetzt. Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>start<break time="10000ms"/><break time="10000ms"/><break time="9300ms"/>noch 30 Sekunden. Du hast bereits die hälfte geschafft.<break time="10000ms"/><break time="10000ms"/><break time="2500ms"/>Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>stop`)
	assert.Equal(t, timeText(120), `Du musst die Übung 2 Minuten durchhalten. Wir starten jetzt. Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>start<break time="10000ms"/><break time="10000ms"/><break time="9300ms"/>noch 1 Minute und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="7950ms"/>noch 1 Minute. Du hast bereits die hälfte geschafft.<break time="10000ms"/><break time="10000ms"/><break time="6850ms"/>noch 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="4650ms"/>Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>stop`)
	assert.Equal(t, timeText(245), `Du musst die Übung 4 Minuten und 5 Sekunden durchhalten. Wir starten jetzt. Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>start<break time="4300ms"/>noch 4 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 3 Minuten und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="6850ms"/>noch 3 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 2 Minuten und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="4850ms"/>noch 2 Minuten und 2 Sekunden. Du hast bereits die hälfte geschafft.noch 2 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 1 Minute und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="7950ms"/>noch 1 Minute.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="4650ms"/>Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>stop`)
	assert.Equal(t, timeText(375), `Du musst die Übung 6 Minuten und 15 Sekunden durchhalten. Wir starten jetzt. Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>start<break time="10000ms"/><break time="4300ms"/>noch 6 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 5 Minuten und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="6850ms"/>noch 5 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 4 Minuten und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="6850ms"/>noch 4 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 3 Minuten und 30 Sekunden.<break time="10000ms"/><break time="9850ms"/>noch 3 Minuten und 7 Sekunden. Du hast bereits die hälfte geschafft.<break time="2800ms"/>noch 3 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 2 Minuten und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="6850ms"/>noch 2 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 1 Minute und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="7950ms"/>noch 1 Minute.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="4650ms"/>Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>stop`)
	assert.Equal(t, timeText(500), `Du musst die Übung 8 Minuten und 20 Sekunden durchhalten. Wir starten jetzt. Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>start<break time="10000ms"/><break time="9300ms"/>noch 8 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 7 Minuten und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="6850ms"/>noch 7 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 6 Minuten und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="6850ms"/>noch 6 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 5 Minuten und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="6850ms"/>noch 5 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 4 Minuten und 30 Sekunden.<break time="10000ms"/><break time="6850ms"/>noch 4 Minuten und 10 Sekunden. Du hast bereits die hälfte geschafft.<break time="4067ms"/>noch 4 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 3 Minuten und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="6850ms"/>noch 3 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 2 Minuten und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="6850ms"/>noch 2 Minuten.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 1 Minute und 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="7950ms"/>noch 1 Minute.<break time="10000ms"/><break time="10000ms"/><break time="8220ms"/>noch 30 Sekunden.<break time="10000ms"/><break time="10000ms"/><break time="4650ms"/>Drei<break time="500ms"/>Zwei<break time="500ms"/>Eins<break time="500ms"/>stop`)
}

func BenchmarkTimeText(b *testing.B) {
	for n := 0; n < b.N; n++ {
		timeText(500)
	}
}

func switchNextUnit(t *testing.T,fn func(*State) bool) {
	state := GetBeginningState()
	for p := 0; p < 4; p++ {
		for w := 0; w < 10; w++ {
			for d := 0; d < len(trainings[w].TrainingDays); d++ {
				for u := 0; u < len(trainings[w].TrainingDays[d].Exercises[p]); u++ {
					assert.Equal(t, state, State{Level: trainingLevel(p), Week: w, Day: d, Unit: u})
					if !(p == 3 && w == 9 && d == 4 && u == 2) {
						assert.Assert(t, fn(&state) == false)
					} else {
						assert.Assert(t, fn(&state) == true) // end of alle trainings
					}
				}
			}
		}
	}
}

func TestSwitchNextUnit(t *testing.T) {
	switchNextUnit(t,switchToNextTrainingLinear)
	switchNextUnit(t,switchToNextTrainingRecursive)
}

func BenchmarkSwitchNextUnit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		state := GetBeginningState()
		SwitchToNextTraining(&state)
	}
}
