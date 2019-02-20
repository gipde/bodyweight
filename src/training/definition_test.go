package training

import (
	"testing"
)

func TestAllTrainings(t *testing.T) {
	count := 0
	for level := 0; level < 4; level++ {
		t.Logf("+++ Level %d", level+1)
		for iWeek, week := range trainings {
			t.Logf("=== %d Woche: ", iWeek+1)
			for iDay, day := range week.TrainingDays {
				t.Logf("--- %d Tag: ", iDay+1)
				for i, unit := range day.Exercises[level] {
					t.Logf("%d |%d. %s, Seite: %d", count, i+1, unit.Exercise.Name, unit.Exercise.Page)
					if unit.Note != nil {
						t.Logf("       %s", unit.Note.text())
					}
					count++
				}
				t.Log()
			}
			t.Log()
		}
		t.Log()
	}
}

