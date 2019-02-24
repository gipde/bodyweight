package training

import (
	"bodyweight/tools"
	"fmt"
)

const (
	DNL = "\n\n"
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

	day := s.getDay()
	exes := s.getExes(day)

	text := "Du absolvierst zur Zeit das "
	switch s.Level {
	case basisProgram:
		text += s.Level.name()
	default:
		text += fmt.Sprintf(`<lang xml:lang="en-US">%s</lang>`, s.Level.name())
	}

	text += fmt.Sprintf(" und bist in der %d. Trainingswoche beim %d. Übungstag angelangt. ", s.Week+1, s.Day+1)
	text += fmt.Sprintf(`Das Training steht in dieser Woche unter dem Motto "%s". `, week.Description)
	text += fmt.Sprintf("Das derzeitige Trainingsprogramm enthält insgesamt "+
		"%d verschiedene Übungen und ist mit %s durchzuführen. ", len(exes), day.Method.name())

	if day.Method == StufenIntervall {
		text += "Die Intervalle dauern "
		text += timeAsStr(getLevelTimeStufenIntervall(s.Level)) + ". "
	}

	return text
}

// ExplainExercise explains the next Exercise
func (s State) ExplainExercise() string {

	day := s.getDay()
	exes := s.getExes(day)
	unit := exes[s.Unit]

	var text string
	switch day.Method {
	case StufenIntervall, HochIntensitaetsSatz:
		text = fmt.Sprintf("Als nächste %s steht mit %s an: ", unit.Exercise.Type.name(), day.Method.name())
		text += fmt.Sprintf("%s. ", unit.Exercise.Name)
		text += addNote(unit)
		text += fmt.Sprintf("Nähere Infos findest Du auf Seite %d im Buch.", unit.Exercise.Page)
	case IntervallSatz, ZirkelIntervall:
		text = fmt.Sprintf("Die nächsten Übungen sind mit %s zu absolvieren:\n", day.Method.name())
		if day.Method == IntervallSatz {
			text += "Es sind 6 bis 12 Wiederholungen durchzuführen. "
		}
		for i, e := range exes {
			text += fmt.Sprintf("%s Übung: %s. %sSeite %d im Buch.\n",
				getOrdinal(i+1), e.Exercise.Name, addNote(e), e.Exercise.Page)
		}
	case SuperSatz:
		text = "Die nächsten Übungen sind mit Supersätzen zu absolvieren. " +
			"Bei der ersten Übung sind 1 bis 5 Wiederholungen, bei der zweiten Übung 6 bis 12 Wiederholungen durchzuführen:\n"
		for i := 0; i < len(exes)/2; i++ {
			e1 := exes[i*2]
			e2 := exes[i*2+1]
			text += fmt.Sprintf("%ss Übungspaar:\n", getOrdinal(i+1))
			text += fmt.Sprintf("%s. %sSeite %d im Buch.\n", e1.Exercise.Name, addNote(e1), e1.Exercise.Page)
			text += fmt.Sprintf("%s. %sSeite %d im Buch.\n", e2.Exercise.Name, addNote(e2), e2.Exercise.Page)
		}
	}
	return text
}

// ShortProgress describes Progress in short
func (s State) ShortProgress() string {
	return fmt.Sprintf("%d. Woche, %d. Tag", s.Week+1, s.Day+1)
}

// CardDayDescription of Trainings a day for Card
func (s State) CardDayDescription() string {
	day := s.getDay()
	exes := s.getExes(day)

	switch day.Method {
	case SuperSatz, IntervallSatz, ZirkelIntervall:
		return s.CardUnitDescription()
	}

	text := s.cardHeader(false)
	for i := range exes {
		n := s
		n.Unit = i
		text += fmt.Sprintf("%d. %s", i+1, n.cardUnitDescription(false))
		if i+1 != len(exes) {
			text += DNL
		}
	}
	return text

}

// CardUnitDescription describes the training in short for Card
func (s State) CardUnitDescription() string {
	return s.cardUnitDescription(true)
}

// WasLastUnit returns true, if
func (s State) WasLastUnit() bool {
	return s.Unit == 0
}

// InstructTraining return's the training instructions per unit
func (s *State) InstructTraining() string {
	switch s.getDay().Method {
	case StufenIntervall:
		return s.StufenIntervallText()
	case IntervallSatz:
		return s.IntervallSatzText()
	case SuperSatz:
		return s.superSatzText()
	case HochIntensitaetsSatz:
		return s.HochIntensitaetsSatzText()
	case ZirkelIntervall:
		return s.ZirkelIntervallText()
	}
	return "Fehler: ungültige Trainingsmethode"
}

func (s *State) getDay() trainingDay {
	return trainings[s.Week].TrainingDays[s.Day]
}

func (s *State) getExesD() []tExercise {
	return s.getDay().Exercises[s.Level]
}

func (s *State) getExes(d trainingDay) []tExercise {
	return d.Exercises[s.Level]
}

func (s *State) StufenIntervallText() string {
	text := "Stufenintervalle mit einer Dauer von "
	sec := getLevelTimeStufenIntervall(s.Level)
	text += timeAsStr(sec) + ". "
	unit := s.getExesD()[s.Unit]
	text += "Wir starten mit: "
	text += unit.Exercise.Name + ". "
	text += addNote(unit)

	s.switchToNextTraining()

	return text + timeText(sec, 30, false, true, true, true)
}

func (s *State) IntervallSatzText() string {

	exes := s.getExesD()

	text := fmt.Sprintf("Bei Intervallsätzen sind %d Übungen mit je 3 Sätzen zu je 3 Minuten auszuführen. ",
		len(exes))
	text += "Pro Satz machst Du 6 bis 12 Wiederholungen. Wechsle die Seite nach einem Satz, " +
		"sofern nichts anderes angegen ist."

	for i := 0; i < len(exes); i++ {
		sUnit := exes[s.Unit].Exercise
		for j := 0; j < 3; j++ {
			text += fmt.Sprintf("\n%d. Satz: %s. ", j+1, sUnit.Name)
			text += addNote(exes[s.Unit])
			text += countDown("Start. ")
			text += timeText(3*60, 60, false, false, false, false)
		}

		s.switchToNextTraining()
	}
	return text
}

func (s *State) superSatzText() string {
	text := "Heute sind insgesamt 6 Supersätze je 4 Minuten hintereinander auszuführen. "
	text += "Jeder Supersatz besteht aus 2 Übungen. Wechsle die Seite nach jeder Wiederholung, " +
		"sofern nichts anderes angegeben ist.\n"

	exes := s.getExesD()
	for i := 0; i < len(exes); i++ {

		text += fmt.Sprintf("%d. Supersatz, %d. Satz, ", i/2+1, i%2+1)
		for j := 0; j < 2; j++ {

			exercise := exes[(i/2)*2+j]
			text += fmt.Sprintf(`%s Übung mit %d bis %d Wiederholungen: `, getOrdinal(j+1), j*5+1, j*6+5+j)
			text += fmt.Sprintf(`%s. `, exercise.Exercise.Name)
			text += addNote(exercise)
		}
		s.switchToNextTraining()
		text += timeText(3*60+40, 60, false, false, true, false)
		text += "\n"
	}
	text += ""
	return text
}

func (s *State) HochIntensitaetsSatzText() string {

	exes := s.getExesD()
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

func (s *State) ZirkelIntervallText() string {
	text := "Heute steht ein Zirkeltraining an. Wechsel zwischen folgenden drei " +
		"Übungen und versuche soviele Sätze wie möglich. "

	exes := s.getExesD()
	for i := 0; i < len(exes); i++ {
		ex := exes[s.Unit]
		exText := ex.Exercise.Name + ". "
		exText += addNote(ex)

		text += fmt.Sprintf("\n"+`%s Übung: %s`, getOrdinal(i+1), exText)
		s.switchToNextTraining()
	}

	text += "\n" + timeText(20*60, 5*60, true, true, true, true)
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

func (s *State) switchToNextTraining() {
	// if s Unit < Units per day {s Unit +=1 }
	if lt(&s.Unit, len(trainings[s.Week].TrainingDays[s.Day].Exercises[s.Level])) {
		return
	}
	// if s Day < Trainingsday per week { s Day += 1}
	if lt(&s.Day, len(trainings[s.Week].TrainingDays)) {
		return
	}
	// if s Week < Trainingsweeks { s Week += 1}
	if lt(&s.Week, len(trainings)) {
		return
	}
	tLevel := int(s.Level)
	if lt(&tLevel, len(trainings[s.Week].TrainingDays[s.Day].Exercises)) {
		s.Level = trainingLevel(tLevel)
		return
	}

	// if reached God-Level :) -> switch back to Week 7
	s.Level = 3
	s.Week = 6
	s.Day = 0
	s.Unit = 0
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

func timeText(sec, interval int, begin, half, countin, countout bool) string {
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
	ret += "Ende. "

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
	n := ex.Note.text()
	if n != "" {
		return fmt.Sprintf("Anmerkung: %s. ", n)
	}
	return ""
}

func getOrdinal(i int) string {
	switch i {
	case 1:
		return "erste"
	case 2:
		return "zweite"
	case 3:
		return "dritte"
	case 4:
		return "vierte"
	default:
		return "Fehler"
	}
}

func (e tExercise) getCardInfo() string {
	ex := e.Exercise
	text := fmt.Sprintf("%s\n", ex.Name)

	if n := e.Note.text(); n != "" {
		text += n + "\n"
	}
	text += fmt.Sprintf("Seite %d", ex.Page)
	return text
}

func (s State) cardUnitDescription(displayHeader bool) string {
	day := s.getDay()
	exes := s.getExes(day)
	var text string

	switch day.Method {
	case StufenIntervall, HochIntensitaetsSatz:
		if displayHeader {
			text = s.cardHeader(true)
		}
		text += exes[s.Unit].getCardInfo()
	case IntervallSatz, ZirkelIntervall:
		if displayHeader {
			text = s.cardHeader(false)
		}
		for i, e := range exes {
			text += fmt.Sprintf("%d. %s", i+1, e.getCardInfo())
			if i+1 < len(exes) {
				text += DNL
			}
		}
	case SuperSatz:
		if displayHeader {
			text = s.cardHeader(false)
		}
		for i := 0; i < len(exes)/2; i++ {
			text += fmt.Sprintf("%d.1. %s\n", i+1, exes[i*2].getCardInfo())
			text += fmt.Sprintf("%d.2. %s", i+1, exes[i*2+1].getCardInfo())
			if i+1 < len(exes)/2 {
				text += DNL
			}
		}
	}
	return text
}

func (s State) cardHeader(withUnit bool) string {
	day := s.getDay()
	var uTxt string
	if withUnit {
		uTxt = fmt.Sprintf(", %d. Übung", s.Unit+1)
	}
	var repeats string
	switch day.Method {
	case SuperSatz:
		repeats = ". 1. Übung 1-5 Wdh., 2. Übung 6-12 Wdh."
	case IntervallSatz:
		repeats = ". 6 bis 12 Wdh."
	}
	return fmt.Sprintf("%s%s\nTrainingsmethode: %s%s\n\n",
		s.ShortProgress(),
		uTxt,
		day.Method.name(),
		repeats)
}
