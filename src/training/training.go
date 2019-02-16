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
	day, exes, _ := s.getDayExesAndUnit()

	text := "Du absolvierst zur Zeit das "
	switch s.Level {
	case basisProgram:
		text += s.Level.name()
	default:
		text += fmt.Sprintf(`<lang xml:lang="en-US">%s</lang>`, s.Level.name())
	}

	text += fmt.Sprintf(" und bist in der %d. Trainingswoche beim %d. Übungstag angelangt. ", s.Week+1, s.Day+1)
	text += fmt.Sprintf(`Das Training steht in dieser Woche unter dem Motto "%s". `, week.Description)
	text += fmt.Sprintf("Das derzeitige Trainingsprogramm enthält insgesamt %d verschiedene Übungen und ist mit %s durchzuführen. ", len(exes), day.Method.name())

	if day.Method == stufenIntervall {
		text += "Die Intervalle dauern "
		text += timeAsStr(getLevelTimeStufenIntervall(s.Level)) + ". "
	}

	text += "Folgende Übungen sind noch durchzuführen: "
	for i := s.Unit; i < len(exes); i++ {
		text += fmt.Sprintf("\n%d. Übung: %s. %s", i+1, exes[i].Exercise.Name, addNote(exes[i]))
	}
	return text
}

// ShortProgress describes Progress in short
func (s State) ShortProgress() string {
	return fmt.Sprintf("%d. Woche, %d. Tag", s.Week+1, s.Day+1)
}

// DayShortDescription of Trainings a day
func (s State) DayShortDescription() string {
	day, exes, _ := s.getDayExesAndUnit()
	text := fmt.Sprintf("Trainingsmethode: %s\n", day.Method.name())
	for i, ex := range exes {
		text += fmt.Sprintf("%d. %s %s\n", i+1, ex.Exercise.Name, addNote(ex))
	}
	return text
}

func (e tExercise) getShortInfo() string {
	ex := e.Exercise
	text := fmt.Sprintf("%s\n", ex.Name)
	if e.Note.text() != "" {
		text += fmt.Sprintf("%s\n", e.Note)
	}
	text += fmt.Sprintf("Seite %d", ex.Page)
	return text
}

// UnitShortDescription describes the training in short
func (s State) UnitShortDescription() string {
	day, exes, unit := s.getDayExesAndUnit()
	text := fmt.Sprintf("%s\n", day.Method.name())
	switch day.Method {
	case stufenIntervall, hochIntensitaetsSatz:
		text += unit.getShortInfo()
	case intervallSatz, zirkelIntervall:
		for i, e := range exes {
			text += fmt.Sprintf("%d. %s\n", i+1, e.getShortInfo())
		}
	case superSatz:
		for i := 0; i < len(exes)/2; i++ {
			text += fmt.Sprintf("%d. Supersatz\n", i+1)
			text += fmt.Sprintf("%s\n", exes[i*2].getShortInfo())
			text += fmt.Sprintf("%s\n\n", exes[i*2+1].getShortInfo())
		}
	}
	return text
}

// ExplainExercise explains the next Exercise
func (s State) ExplainExercise() string {
	// TODO: wenn Supersatz, Intervallsatz und Zirkelintervall müssen alle Übungen angesagt werden
	_, exes, _ := s.getDayExesAndUnit()
	ex := exes[s.Unit].Exercise
	note := addNote(exes[s.Unit])
	return fmt.Sprintf(`Als nächste %s steht mit Schwierigkeitsgrad %d an: %s. %sGenauere Infos findest Du auf Seite %d im Buch. `,
		ex.Type.name(), ex.Difficulty, ex.Name, note, ex.Page)
}

// WasLastUnit returns true, if
func (s State) WasLastUnit() bool {
	return s.Unit == 0
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
	text := "Stufenintervalle mit einer Dauer von "
	sec := getLevelTimeStufenIntervall(s.Level)
	text += timeAsStr(sec) + ". "
	_, _, unit := s.getDayExesAndUnit()
	text += "Wir starten mit: "
	text += unit.Exercise.Name + ". "
	text += addNote(unit)

	s.switchToNextTraining()

	return text + timeText(sec, 30, false, true, true, true, true)
}

func (s *State) intervallSatzText() string {

	_, exes, _ := s.getDayExesAndUnit()

	text := fmt.Sprintf("Bei den Intervallsätzen sind heute insgesamt %d Übungen mit je 3 Sätzen zu je 3 Minuten auszuführen. ", len(exes))
	text += "Pro Satz machst Du bis zu 12 Wiederholungen. "

	for i := 0; i < len(exes); i++ {
		sUnit := exes[s.Unit].Exercise
		for j := 0; j < 3; j++ {
			text += fmt.Sprintf("\n%d. Satz: %s. ", (i*3)+j+1, sUnit.Name)
			text += addNote(exes[s.Unit])
			text += countDown("Start. ")
			text += timeText(3*60, 60, false, false, true, false, false)
		}

		s.switchToNextTraining()
	}
	return text
}

func (s *State) superSatzText() string {
	text := "Heute sind insgesamt 6 Supersätze je 4 Minuten hintereinander auszuführen. Jeder Supersatz besteht aus 2 Übungen. \n"

	_, exes, _ := s.getDayExesAndUnit()
	for i := 0; i < len(exes); i++ {

		text += fmt.Sprintf("%d. Satz, ", i+1)
		for j := 0; j < 2; j++ {

			exercise := exes[(i/2)*2+j]
			text += "\n"
			text += fmt.Sprintf(`<say-as interpret-as="ordinal">%d</say-as> Übung mit bis zu %d Wiederholungen: `, j+1, j*6+6)
			text += fmt.Sprintf(`%s. `, exercise.Exercise.Name)
			text += addNote(exercise)
		}
		s.switchToNextTraining()
		text += timeText(3*60+40, 60, false, false, true, true, false)
		text += "\n"
	}
	text += ""
	return text
}

func (s *State) hochIntensitaetsSatzText() string {

	_, exes, _ := s.getDayExesAndUnit()
	ex := exes[s.Unit]

	text := `Es sind 8 Hochintensitätssätze mit je 20 Sekunden Training gefolgt von 10 Sekunden Pause durchzuführen. 
	Die Anzahl der Wiederholungen sollte dabei in etwa gleich sein. `
	text += "\nWir starten mit: " + ex.Exercise.Name + "."
	text += addNote(ex)
	text += "\n"

	for i := 0; i < 8; i++ {
		text += countDown("Start. ")
		text += breakFor(20 * 1000)
		text += "Pause. "
		text += breakFor(6 * 1000)
		text += "\n"
	}
	text += countDown("Ende")
	s.switchToNextTraining()
	return text
}

func (s *State) zirkelIntervallText() string {
	text := "Heute steht ein Zirkeltraining an. Wechsel zwischen folgenden drei Übungen und versuche soviele Sätze wie möglich. "

	_, exes, _ := s.getDayExesAndUnit()
	for i := 0; i < len(exes); i++ {
		ex := exes[s.Unit]
		exText := ex.Exercise.Name + ". "
		exText += addNote(ex)

		text += fmt.Sprintf("\n"+`<say-as interpret-as="ordinal">%d</say-as> Übung: %s`, i+1, exText)
		s.switchToNextTraining()
	}

	text += "\n" + timeText(20*60, 5*60, true, true, true, true, true)
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

func (s *State) switchToNextTraining() bool {
	// if s Unit < Units per day {s Unit +=1 }
	if lt(&s.Unit, len(trainings[s.Week].TrainingDays[s.Day].Exercises[s.Level])) {
		return false
	}
	// if s Day < Trainingsday per week { s Day += 1}
	if lt(&s.Day, len(trainings[s.Week].TrainingDays)) {
		return false
	}
	// if s Week < Trainingsweeks { s Week += 1}
	if lt(&s.Week, len(trainings)) {
		return false
	}
	tLevel := int(s.Level)
	// if s Level < Levels (Trainings) { s Level}
	if lt(&tLevel, len(trainings[s.Week].TrainingDays[s.Day].Exercises)) {
		s.Level = trainingLevel(tLevel)
		return false
	}

	// TODO: How to handle Good-Level
	s.Level = 4
	s.Week = 0
	s.Day = 0
	s.Unit = 0
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

func getLevelTimeStufenIntervall(level trainingLevel) int {
	var sec int
	switch level {
	case basisProgram:
		sec = 4*60 + 30
	case firstClass:
		sec = 5*60 + 30
	case masterClass:
		sec = 6*60 + 30
	case chiefClass:
		sec = 7*60 + 30
	}
	return sec
}

func addNote(ex tExercise) string {
	if ex.Note.text() != "" {
		return fmt.Sprintf("Anmerkung: %s. ", ex.Note)
	}
	return ""
}
