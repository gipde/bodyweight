package training

import (
	"fmt"
)

func GetCurrentState(current *State) string {
	week := trainings[current.Week]
	day := trainings[current.Week].TrainingDays[current.Day]
	exes := day.Exercises[current.Level]
	text := fmt.Sprintf("Du absolvierst zur Zeit das %s ", current.Level.name())
	text += fmt.Sprintf("und bist in der %d. Trainingswoche beim %d. Übungstag angelangt. ", current.Week+1, current.Day+1)
	text += fmt.Sprintf(`Das Training steht in dieser Woche unter dem Motto "%s". `, week.Description)
	text += fmt.Sprintf("Das anstehende Training enthält insgesagt %d verschiedene Übungen und ist mit %s durchzuführen. ", len(exes), day.Method.name())
	if day.Method == stufenIntervall {
		switch current.Level {
		case basisProgram:
			text += fmt.Sprintf("Im %s dauern die Intervalle 4 Minuten und 30 Sekunden. ", current.Level.name())
		case firstClass:
			text += fmt.Sprintf(`In der <lang xml:lang="en-US">%s</lang> dauern die Intervalle 5 Minuten und 30 Sekunden. `, current.Level.name())
		case masterClass:
			text += fmt.Sprintf(`In der <lang xml:lang="en-US">%s</lang> dauern die Intervalle 6 Minuten und 30 Sekunden. `, current.Level.name())
		case chiefClass:
			text += fmt.Sprintf(`In der <lang xml:lang="en-US">%s</lang> dauern die Intervalle 7 Minuten und 30 Sekunden. `, current.Level.name())
		}
	}
	for i, ex := range exes {
		text += fmt.Sprintf("%d. Übung: %s. ", i+1, ex.Exercise.get().Name)
	}
	return text
}

func GetCurrentExercise(current *State) string {
	exes := trainings[current.Week].TrainingDays[current.Day].Exercises[current.Level]

	// TODO: Sonderlösung für Interfallsätze und ?? 
	ex := exes[exes[0].Exercise]
	return fmt.Sprintf("Als nächste %s steht an: %s. Genauere Infos findest Du auf Seite %d im Buch. ", ex.Exercise.get().Type.name(), ex.Exercise.get().Name, ex.Exercise.get().Page)
}

func checkState(state *State) error {
	return nil
}

// GetBeginningState initial Trainig State
func GetBeginningState() State {
	return State{
		Level: basisProgram,
		Week:  0,
		Day:   0,
		Unit:  0,
	}
}

//StartTraining starting Training
func StartTraining(state *State) string {
	return timeText(4*60 + 30)
}

// AnnounceDailyTraining tells about Training Content
func AnnounceDailyTraining(current *State) string {
	if checkState(current) != nil {
		return "Dein Trainingszustand enthält einen Fehler. Wende dich bitte an den Administrator"
	}

	text := GetCurrentState(current)

	return text

}

//SwitchToNextTraining - switch to next Training Unit
func SwitchToNextTraining(last *State) bool {
	return switchToNextTrainingLinear(last)
}

type check struct {
	val *int
	le  int
}

func recur(last *State, arr []check) bool {
	if len(arr) == 0 { // recursion exit fn
		return true // initial value for first check
	}
	r := recur(last, arr[1:])
	if r { // check only if from righter position set to zero
		if *arr[0].val < arr[0].le-1 { // if checked value < max value
			*arr[0].val++      // increment value +1
			if len(arr) == 4 { // if on outer position (copy value)
				last.Level = trainingLevel(*arr[0].val)
			}
			return false // the left side position should not checkt
		}
		*arr[0].val = 0 // checkd value == max value > set to zero
		return true     // the left side position should be checked
	}
	return false // no check
}

func switchToNextTrainingRecursive(last *State) bool {
	tLevel := int(last.Level)
	return recur(last, []check{
		{&tLevel, len(trainings[last.Week].TrainingDays[last.Day].Exercises)},
		{&last.Week, len(trainings)},
		{&last.Day, len(trainings[last.Week].TrainingDays)},
		{&last.Unit, len(trainings[last.Week].TrainingDays[last.Day].Exercises[last.Level])},
	})
}

func lt(val *int, aLen int) bool {
	if *val < aLen-1 {
		*val++
		return true
	}
	*val = 0
	return false
}

func switchToNextTrainingLinear(last *State) bool {
	if lt(&last.Unit, len(trainings[last.Week].TrainingDays[last.Day].Exercises[last.Level])) {
		return false
	}
	if lt(&last.Day, len(trainings[last.Week].TrainingDays)) {
		return false
	}
	if lt(&last.Week, len(trainings)) {
		return false
	}
	tLevel := int(last.Level)
	if lt(&tLevel, len(trainings[last.Week].TrainingDays[last.Day].Exercises)) {
		last.Level = trainingLevel(tLevel)
		return false
	}
	return true
}

func timeAsStr(sec int) string {
	var res string
	min := sec / 60
	s := sec % 60
	if min > 1 {
		res += fmt.Sprintf("%d Minuten", min)
	} else if min > 0 {
		res += fmt.Sprintf("%d Minute", min)
	}
	if s > 0 {
		if min > 0 {
			res += " und "
		}
		res += fmt.Sprintf("%d Sekunden", s)
	}
	if sec == 0 {
		res = "Ende"
	}
	return res
}

func count(word string) string {
	var res string
	br := "<break time=\"500ms\"/>"
	for i, r := range []string{"Drei", "Zwei", "Eins"} {
		res += r
		if i < 2 {
			res += br
		}
	}
	return res + br + word
}

func breakFor(ms int) string {
	divider := 10000
	var res string
	br := `<break time="%dms"/>`
	for ms >= divider {
		ms -= divider
		res += fmt.Sprintf(br, divider)
	}
	if ms > 0 {
		res += fmt.Sprintf(br, ms)
	}
	return res
}

func addTimeInfo(sec, half, breakTime, breakTimeSub int) (string, int) {
	waitTime := timeAsStr(sec)
	if sec == half {
		waitTime += ". Du hast bereits die hälfte geschafft"
	}
	breakTimeSub += 700 // für Context Wechsel
	breakTime -= breakTimeSub
	res := fmt.Sprintf("%snoch %s.", breakFor(breakTime), waitTime)
	if len(waitTime) < 11 {
		breakTimeSub = 1080
	} else if len(waitTime) < 25 {
		breakTimeSub = 1350
	} else if len(waitTime) < 48 {
		breakTimeSub = 2450
	} else if len(waitTime) < 64 {
		breakTimeSub = 3500
	} else {
		breakTimeSub = 5233
	}
	return res, breakTimeSub
}

func timeText(sec int) string {
	res := "Du musst die Übung " + timeAsStr(sec) + " durchhalten. Wir starten jetzt. "
	res += count("start")
	half := sec / 2
	count("start")
	intervall := 30
	var breakTimeSub int
	var breakTime int
	var line string
	for sec > intervall {
		sec--
		breakTime += 1000
		if sec%intervall == 0 || sec == half {
			line, breakTimeSub = addTimeInfo(sec, half, breakTime, breakTimeSub)
			res += line
			breakTime = 0
		}
	}
	res += breakFor(intervall*1000 - 4000 - breakTimeSub)
	res += count("stop")
	return res
}
