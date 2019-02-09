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

// IsLastUnit returns true, if
func (s State) IsLastUnit() bool {
	exes := trainings[s.Week].TrainingDays[s.Day].Exercises[s.Level]
	if s.Unit+1 == len(exes) {
		return true
	}
	return false
}

// InstructTraining return's the training instructions per unit
func (s *State) InstructTraining() string {
	day, _, _ := s.getDayExesAndUnit()

	switch day.Method {
	case stufenIntervall:
		return s.stufenIntervallText()
	case intervallSatz:
		return s.intervallSatzText()
	case superSatz:
		return s.superSatzText()
	case hochIntensitaetsSatz:
		return s.hochIntensitaetsSatzText()
	case zirkelIntervall:
		return s.zirkelIntervallText()
	}
	return "Fehler: ungültige Trainingsmethode"
}

func (s *State) getDayExesAndUnit() (trainingDay, []tExercise, tExercise) {
	day := trainings[s.Week].TrainingDays[s.Day]
	exes := day.Exercises[s.Level]
	unit := exes[s.Unit]
	return day, exes, unit
}

func (s *State) stufenIntervallText() string {
	var sec int
	switch s.Level {
	case basisProgram:
		sec = 4*60 + 30
	case firstClass:
		sec = 5*60 + 30
	case masterClass:
		sec = 6*60 + 30
	case chiefClass:
		sec = 7*60 + 30
	}

	_, _, unit := s.getDayExesAndUnit()
	text := "Als nächste Übung steht an: "
	text += unit.Exercise.get().Name + ". "
	if unit.Note != "" {
		text += "Anmerkung: " + unit.Note + ". "
	}

	s.switchToNextTraining()

	return text + timeText(sec, 30, true, true, true, true, true)
}

func (s *State) intervallSatzText() string {

	_, exes, _ := s.getDayExesAndUnit()

	text := fmt.Sprintf("Bei den Intervallsätzen sind heute insgesamt %d Übungen mit je 3 Sätzen zu je 3 Minuten auszuführen. ", len(exes))
	text += "Pro Satz machen Sie bis zu 12 Wiederholugen. "

	for i := 0; i < len(exes); i++ {
		currentUnit := exes[s.Unit].Exercise.get()
		for j := 0; j < 3; j++ {
			text += fmt.Sprintf("\n%d. Satz: %s.  ", j+1, currentUnit.Name)
			if i == 0 && j == 0 {
				text += countDown("Start. ")
			}
			text += timeText(3*60, 60, false, false, true, false, false)
		}

		s.switchToNextTraining()
	}
	return text
}

func (s *State) superSatzText() string {
	text := "Es sind insgesamt 6 Supersätze je 4 Minuten hintereinander auszuführen. Jeder Supersatz besteht aus 2 Übungen. \n"

	_, exes, _ := s.getDayExesAndUnit()
	for i := 0; i < len(exes); i++ {

		text += fmt.Sprintf("%d. Satz, ", i+1)
		for j := 0; j < 2; j++ {

			exercise := exes[(i/2)*2+j]
			text += fmt.Sprintf("%d. Übung mit bis zu %d Wiederholungen: ", j+1, j*6+6)
			text += fmt.Sprintf("%s. ", exercise.Exercise.get().Name)
			if exercise.Note != "" {
				text += fmt.Sprintf("%s. \n", exercise.Note)
			} else {
				text += "\n"
			}
			s.switchToNextTraining()
		}
		text += timeText(3*60+40, 60, false, false, false, true, false)
		text += "\n"
	}
	text += ""
	return text
}

func (s *State) hochIntensitaetsSatzText() string {

	_, exes, _ := s.getDayExesAndUnit()
	ex := exes[s.Unit]

	text := "Als nächste Übung steht an: " + ex.Exercise.get().Name + ". "
	if ex.Note != "" {
		text += ex.Note + ". "
	}
	text += "Es sind 8 Sätze mit je 20 Sekunden Training gefolgt von 10 Sekunden Pause durchzuführen. Die Anzahl der Wiederholungen sollte dabei in etwa gleich sein. "

	for i := 0; i < 8; i++ {
		text += countDown("Start. ")
		text += breakFor(20 * 1000)
		text += "Pause. "
		text += breakFor(6 * 1000)
	}
	s.switchToNextTraining()
	return text
}

func (s *State) zirkelIntervallText() string {
	text := "Heute steht ein Zirkeltraining an. Wechsel zwischen folgenden drei Übungen und versuche soviele Sätze wie möglich. "

	_, exes, _ := s.getDayExesAndUnit()
	for i := 0; i < len(exes); i++ {
		ex := exes[s.Unit]
		exText := ex.Exercise.get().Name + ". "
		if ex.Note != "" {
			exText += ex.Note + ". "
		}
		text += fmt.Sprintf("%d. Übung: %s", i+1, exText)
		s.switchToNextTraining()
	}

	text += timeText(20*60, 5*60, true, true, true, true, true)
	return text
}

// increment val if it is smaller then aLen
func lt(val *int, aLen int) bool {
	if *val < aLen-1 {
		*val++
		return true
	}
	*val = 0
	return false
}

func (current *State) switchToNextTraining() bool {
	// if current Unit < Units per day {current Unit +=1 }
	if lt(&current.Unit, len(trainings[current.Week].TrainingDays[current.Day].Exercises[current.Level])) {
		return false
	}
	// if current Day < Trainingsday per week { current Day += 1}
	if lt(&current.Day, len(trainings[current.Week].TrainingDays)) {
		return false
	}
	// if current Week < Trainingsweeks { current Week += 1}
	if lt(&current.Week, len(trainings)) {
		return false
	}
	tLevel := int(current.Level)
	// if current Level < Levels (Trainings) { current Level}
	if lt(&tLevel, len(trainings[current.Week].TrainingDays[current.Day].Exercises)) {
		current.Level = trainingLevel(tLevel)
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

func countDown(word string) string {
	var res string
	for _, r := range []string{"Drei", "Zwei", "Eins"} {
		res += r + ". "
	}
	return res + word
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

func timeText(sec, interval int, begin, half, end, countin, countout bool) string {
	halfSec := sec / 2
	secondsGone := 0

	var ret string
	// Begin
	if begin {
		ret += fmt.Sprintf("Noch %s. ", timeAsStr(sec))
		secondsGone -= 3
	}
	if countin {
		ret += countDown("Start. ")
		secondsGone = -4
	}
	for sec > 1 {
		secondsGone++
		sec--
		// Intervall reached
		if sec%interval == 0 {
			ret += breakFor(secondsGone * 1000)
			secondsGone = -2
			ret += fmt.Sprintf("Noch %s. ", timeAsStr(sec))
		}
		// half reached
		if half && sec == halfSec {
			ret += breakFor(secondsGone * 1000)
			secondsGone = -2
			ret += "Die Hälfte ist vorbei. "
		}
	}
	ret += breakFor(secondsGone * 1000)

	if countout {
		ret += countDown("")
	}
	// Ende
	if end {
		ret += "Ende. "
	}

	return ret
}

// // ExplainTraining explains the Training
// func (s State) ExplainTraining() string {
// 	week := trainings[s.Week]
// 	day := trainings[s.Week].TrainingDays[s.Day]
// 	exes := day.Exercises[s.Level]
// 	text := "Du absolvierst zur Zeit das "
// 	switch s.Level {
// 	case basisProgram:
// 		text += s.Level.name()
// 	default:
// 		text += fmt.Sprintf(`<lang xml:lang="en-US">%s</lang>`, s.Level.name())
// 	}
// 	text += fmt.Sprintf(" und bist in der %d. Trainingswoche beim %d. Übungstag angelangt. ", s.Week+1, s.Day+1)
// 	text += fmt.Sprintf(`Das Training steht in dieser Woche unter dem Motto "%s". `, week.Description)
// 	text += fmt.Sprintf("Das derzeitige Trainingsprogramm enthält insgesagt %d verschiedene Übungen und ist mit %s durchzuführen. ", len(exes), day.Method.name())
// 	text += "Die Intervalle dauern "
// 	if day.Method == stufenIntervall {
// 		switch s.Level {
// 		case basisProgram:
// 			text += "4 Minuten und 30 Sekunden. "
// 		case firstClass:
// 			text += "5 Minuten und 30 Sekunden. "
// 		case masterClass:
// 			text += "6 Minuten und 30 Sekunden. "
// 		case chiefClass:
// 			text += "7 Minuten und 30 Sekunden. "
// 		}
// 	}
// 	text += "Folgende Übungen sind noch durchzuführen: "
// 	for i := s.Unit; i < len(exes); i++ {
// 		text += fmt.Sprintf("Die %d. Übung ist: %s. ", i+1, exes[i].Exercise.get().Name)
// 	}
// 	return text
// }

// // ExplainExercise explains the next Exercise
// func (s State) ExplainExercise() string {
// 	exes := getExes(s)

// 	// TODO: Logik wg. zusammenhängender Trainingseinheit
// 	/*
// 		- beim Stufenintervall: jede Übung z.B. 4:30
// 		- beim intervallsatz: komplette Zeit durchmachen mit allen 4 übungen; 3 sätze jew. 3 Minuten; letzter satz bei 33:00
// 		- beim Supersatz: immer zweier Block zusammenmachen;
// 			4 Minuten je supersatz, 2 Suerpsätze pro Übungspar
// 				1-5 wiederholungen beim ersten satz
// 				6-12 wiederholugen für die Übung direkt danach
// 		- Hochintensität: 8 Sätze je 20 Sekunden + 10 Sekunden Pause
// 		- beim Zirkelintervall: alle 3 Übungen im Wechsel insgesamt 20 Minuten
// 	*/
// 	ex := exes[s.Unit].Exercise.get()
// 	note := exes[s.Unit].Note
// 	return fmt.Sprintf("Als nächste %s steht an: %s. %s. Der Schwierigkeitsgrad ist: %d. Genauere Infos findest Du auf Seite %d im Buch. ",
// 		ex.Type.name(), ex.Name, note, ex.Difficulty, ex.Page)
// }

// //StartTraining starting Training
// func (s State) StartTraining() (string, State) {

// 	// TODO: Logik wg. zusammenhängender Trainingseinheit
// 	exes := getExes(s)
// 	ex := exes[s.Unit].Exercise.get()
// 	note := exes[s.Unit].Note
// 	r := fmt.Sprintf("Wir starten mit %s. %s. ", ex.Name, note)

// 	// bitte mit Pointer
// 	// newState := s.SwitchToNextTraining()

// 	return r + timeText(4*60+30), s
// 	// return r + timeText(4*60+30), newState
// }

// func addTimeInfo(sec, half, breakTime, breakTimeSub int) (string, int) {
// 	waitTime := timeAsStr(sec)
// 	if sec == half {
// 		waitTime += ". Du hast bereits die Hälfte geschafft"
// 	}
// 	breakTimeSub += 700 // für Context Wechsel
// 	breakTime -= breakTimeSub
// 	res := fmt.Sprintf("%sNoch %s. ", breakFor(breakTime), waitTime)
// 	if len(waitTime) < 11 {
// 		breakTimeSub = 1080
// 	} else if len(waitTime) < 25 {
// 		breakTimeSub = 1350
// 	} else if len(waitTime) < 48 {
// 		breakTimeSub = 2450
// 	} else if len(waitTime) < 64 {
// 		breakTimeSub = 3500
// 	} else {
// 		breakTimeSub = 5233
// 	}
// 	return res, breakTimeSub
// }

// func timeText(sec int) string {
// 	res := fmt.Sprintf("Die Übung dauert %s. ", timeAsStr(sec))
// 	res += countDown("Start. ")
// 	half := sec / 2
// 	intervall := 30
// 	var breakTimeSub int
// 	var breakTime int
// 	var line string
// 	for sec > intervall {
// 		sec--
// 		breakTime += 1000
// 		if sec%intervall == 0 || sec == half {
// 			line, breakTimeSub = addTimeInfo(sec, half, breakTime, breakTimeSub)
// 			res += line
// 			breakTime = 0
// 		}
// 	}
// 	res += breakFor(intervall*1000 - 4000 - breakTimeSub)
// 	res += countDown("Stop. ")
// 	res += breakFor(2000)
// 	return res
// }

//SwitchToNextTraining - switch to next Training Unit
// func (s *State) SwitchToNextTraining() {
// 	switchToNextTrainingLinear(s)
// }

// type check struct {
// 	val *int
// 	le  int
// }

// func recur(last *State, arr []check) bool {
// 	if len(arr) == 0 { // recursion exit fn
// 		return true // initial value for first check
// 	}
// 	r := recur(last, arr[1:])
// 	if r { // check only if from righter position set to zero
// 		if *arr[0].val < arr[0].le-1 { // if checked value < max value
// 			*arr[0].val++      // increment value +1
// 			if len(arr) == 4 { // if on outer position (copy value)
// 				last.Level = trainingLevel(*arr[0].val)
// 			}
// 			return false // the left side position should not checkt
// 		}
// 		*arr[0].val = 0 // checkd value == max value > set to zero
// 		return true     // the left side position should be checked
// 	}
// 	return false // no check
// }

// func switchToNextTrainingRecursive(last *State) bool {
// 	tLevel := int(last.Level)
// 	return recur(last, []check{
// 		{&tLevel, len(trainings[last.Week].TrainingDays[last.Day].Exercises)},
// 		{&last.Week, len(trainings)},
// 		{&last.Day, len(trainings[last.Week].TrainingDays)},
// 		{&last.Unit, len(trainings[last.Week].TrainingDays[last.Day].Exercises[last.Level])},
// 	})
// }
