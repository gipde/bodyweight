package training

import (
	"bodyweight/tools"
	"fmt"
)

// GetBeginningState initial Trainig State
func GetBeginningState() State {
	return State{
		Level: basisProgram,
		Week:  0,
		Day:   0,
		Unit:  0,
	}
}

// ExplainTraining explains the Training
func (s State) ExplainTraining() string {
	week := trainings[s.Week]
	day := trainings[s.Week].TrainingDays[s.Day]
	exes := day.Exercises[s.Level]
	text := "Du absolvierst zur Zeit das "
	switch s.Level {
	case basisProgram:
		text += s.Level.name()
	default:
		text += fmt.Sprintf(`<lang xml:lang="en-US">%s</lang>`, s.Level.name())
	}
	text += fmt.Sprintf(" und bist in der %d. Trainingswoche beim %d. Übungstag angelangt. ", s.Week+1, s.Day+1)
	text += fmt.Sprintf(`Das Training steht in dieser Woche unter dem Motto "%s". `, week.Description)
	text += fmt.Sprintf("Das derzeitige Trainingsprogramm enthält insgesagt %d verschiedene Übungen und ist mit %s durchzuführen. ", len(exes), day.Method.name())
	text += "Die Intervalle dauern "
	if day.Method == stufenIntervall {
		switch s.Level {
		case basisProgram:
			text += "4 Minuten und 30 Sekunden. "
		case firstClass:
			text += "5 Minuten und 30 Sekunden. "
		case masterClass:
			text += "6 Minuten und 30 Sekunden. "
		case chiefClass:
			text += "7 Minuten und 30 Sekunden. "
		}
	}
	text += "Folgende Übungen sind noch durchzuführen: "
	for i := s.Unit; i < len(exes); i++ {
		text += fmt.Sprintf("Die %d. Übung ist: %s. ", i+1, exes[i].Exercise.get().Name)
	}
	return text
}

func getExes(s State) []tExercise {
	return trainings[s.Week].TrainingDays[s.Day].Exercises[s.Level]
}

// ExplainExercise explains the next Exercise
func (s State) ExplainExercise() string {
	exes := getExes(s)

	// TODO: Logik wg. zusammenhängender Trainingseinheit
	/*
	- beim Stufenintervall: jede Übung z.B. 4:30
	- beim intervallsatz: komplette Zeit durchmachen mit allen 4 übungen; 3 sätze jew. 3 Minuten; letzter satz bei 33:00
	- beim Supersatz: immer zweier Block zusammenmachen;
		4 Minuten je supersatz, 2 Suerpsätze pro Übungspar
			1-5 wiederholungen beim ersten satz
			6-12 wiederholugen für die Übung direkt danach	
	- Hochintensität: 8 Sätze je 20 Sekunden + 10 Sekunden Pause	
	- beim Zirkelintervall: alle 3 Übungen im Wechsel insgesamt 20 Minuten
	*/
	ex := exes[s.Unit].Exercise.get()
	note := exes[s.Unit].Note
	return fmt.Sprintf("Als nächste %s steht an: %s. %s. Der Schwierigkeitsgrad ist: %d. Genauere Infos findest Du auf Seite %d im Buch. ",
		ex.Type.name(), ex.Name, note, ex.Difficulty, ex.Page)
}

//StartTraining starting Training
func (s State) StartTraining() string {

	// TODO: Logik wg. zusammenhängender Trainingseinheit
	exes := getExes(s)
	ex := exes[s.Unit].Exercise.get()
	note := exes[s.Unit].Note
	r := fmt.Sprintf("Wir starten mit %s. %s. ", ex.Name, note)
	return r + timeText(4*60+30)
}

// IsLastUnit returns true, if
func (s State) IsLastUnit() bool {
	exes := getExes(s)
	if s.Unit+1 == len(exes) {
		return true
	}
	return false
}

//SwitchToNextTraining - switch to next Training Unit
func (s State) SwitchToNextTraining() State {
	switchToNextTrainingLinear(&s)
	return s
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
	br := "<break time=\"750ms\"/>"
	for i, r := range []string{"Drei", "Zwei", "Eins"} {
		res += r + ". "
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

	if tools.IsDebug() {
		return fmt.Sprintf(br, 100)
	}

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
		waitTime += ". Du hast bereits die Hälfte geschafft"
	}
	breakTimeSub += 700 // für Context Wechsel
	breakTime -= breakTimeSub
	res := fmt.Sprintf("%sNoch %s. ", breakFor(breakTime), waitTime)
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
	res := fmt.Sprintf("Die Übung dauert %s. ", timeAsStr(sec))
	res += count("Start. ")
	half := sec / 2
	// count("start")
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
	res += count("Stop. ")
	res += breakFor(2000)
	return res
}
