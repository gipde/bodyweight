package training

import (
	"fmt"
)

func getCurrentState(current *TrainingState) string {
	return ""
}

func checkState(state *TrainingState) error {
	return nil
}

func AnnounceDailyTraining(current *TrainingState) string {
	if checkState(current) != nil {
		return "Dein Trainingszustand enthält einen Fehler. Wende dich bitte an den Administrator"
	}
	week := Trainings[current.Week]
	day := Trainings[current.Week].TrainingDays[current.Day]
	exes := day.Exercises[current.Level]

	text := fmt.Sprintf("Du bist derzeit in der %d. Trainingswoche und beim %d. Übungstag angelangt. ", current.Week+1, current.Day+1)

	text += fmt.Sprintf(`Das Training steht in dieser Woche unter dem 
	Motto "%s" und ist heute mit %s durchzuführen. `, week.Description, day.Method.name())

	if day.Method == STUFENINTERVALL {
		switch current.Level {
		case BASISPROGRAM:
			text += fmt.Sprintf("Im %s dauern die Intervalle 4 Minuten und 30 Sekunden. ", current.Level.name())
		case FIRSTCLASS:
			text += fmt.Sprintf(`In der <lang xml:lang="en-US">%s</lang> dauern die Intervalle 5 Minuten und 30 Sekunden. `, current.Level.name())
		case MASTERCLASSC:
			text += fmt.Sprintf(`In der <lang xml:lang="en-US">%s</lang> dauern die Intervalle 6 Minuten und 30 Sekunden. `, current.Level.name())
		case CHIEFCLASS:
			text += fmt.Sprintf(`In der <lang xml:lang="en-US">%s</lang> dauern die Intervalle 7 Minuten und 30 Sekunden. `, current.Level.name())
		}
	}

	text += fmt.Sprintf("Insgesamt musst Du %d verschiedene Übungen durchführen. ", len(exes))

	text += fmt.Sprintf("Die erste Übung ist: %s. ", exes[0].Exercise.Name())
	text += fmt.Sprintf("Die zweite Übung ist: %s. ", exes[1].Exercise.Name())
	text += fmt.Sprintf("Die dritte Übung ist: %s. ", exes[2].Exercise.Name())
	text += fmt.Sprintf("Die vierte Übung ist: %s. ", exes[3].Exercise.Name())

	text += "Wenn Du beginnen möchtest, sage: ich bin bereit."

	return text

}

func switchToNextTraining(last *TrainingState) *TrainingState {
	// inc values
	// return getCurrent
	return nil
}

func timeAsStr(sec int) string {
	var res string
	min := sec / 60
	s := sec % 60
	if min > 0 {
		res += fmt.Sprintf("%d Minuten", min)
	}
	if s > 0 {
		if min > 0 {
			res += " und "
		}
		res += fmt.Sprintf("%d Sekunden", s)
	}
	return res
}

func count(word string) string {
	var res string
	br := "<break time=\"500ms\"/>"
	for _, r := range []string{"drei", "zwei", "eins"} {
		res += fmt.Sprintf("%s%s", br, r)
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
	res := "Du musst die Übung " + timeAsStr(sec) + " durchhalten."
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
