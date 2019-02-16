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

// func TestAlleNotes(t *testing.T) {
// 	set := make(map[string]struct{})
// 	for _, week := range trainings {
// 		for _, day := range week.TrainingDays {
// 			for _, level := range day.Exercises {
// 				for _, unit := range level {
// 					if unit.Note != "" {
// 						if _, ok := set[unit.Note]; !ok {
// 							set[unit.Note] = struct{}{}
// 						}
// 					}
// 				}
// 			}

// 		}
// 	}

//     keys := make([]string, 0, len(set))
//     for k := range set {
//         keys = append(keys, k)
//     }
// 	sort.Strings(keys)
// 	for _, n := range keys {
// 		t.Logf("= \"%s\"",n)
// 	}
// }

// import (
// 	"fmt"
// 	"sort"
// 	"testing"
// )

// type ExCollection []exercise

// func (e ExCollection) Len() int {
// 	return len(e)
// }

// func (e ExCollection) Swap(i, j int) {
// 	e[i], e[j] = e[j], e[i]
// }

// func (e ExCollection) Less(i, j int) bool {
// 	return e[i].Page < e[j].Page
// }

// var varNames = []string{
// 	"liegestuetzMitErhoehtenHaenden",
// 	"tuerziehen",
// 	"Seestern",
// 	"umgekehrtesBankdruecken",
// 	"ausfallschrittNachHintenImWechsel",
// 	"rumaenischesKreuzhebenAufEinemBeinImWechsel",
// 	"engeKniebeuge",
// 	"schwimmer",
// 	"seitlicherAusfallschritt",
// 	"schraegerCrunch",
// 	"klassischerLiegestuetz",
// 	"militaryPressMitErhoehtenHaenden",
// 	"engerLiegestuetzMitErhoehtenHaenden",
// 	"kniebeugeImAusfallschritt",
// 	"curlMitHandtuch",
// 	"beinheber",
// 	"kaefer",
// 	"russischerTwist",
// 	"crunch",
// 	"liegestuetzMitErhoehtenFuessen",
// 	"liegestuetzMitAbstossen",
// 	"militaryPress",
// 	"baerenGang",
// 	"engerLiegestuetz",
// 	"gesprungeneKniebeuge",
// 	"ausfallschrittNachVorneImWechsel",
// 	"klimmzugMitUnterstuetzung",
// 	"umgekehrtesBankdrueckenImUntergriff",
// 	"tuerziehenImUntergriff",
// 	"vUp",
// 	"kaeferIpsilateral",
// 	"crunchItUp",
// 	"haengendesBeinheben",
// 	"pogoSprung",
// 	"hueftTwist",
// 	"knieheberImStehen",
// 	"bodyRock",
// 	"vierPhasenLiegestuetz",
// 	"goodMorning",
// 	"kaeferKontralateral",
// 	"seitlichesHueftheben",
// 	"aufstehenAusDemEinbeinigenKniestand",
// 	"ausfallSchritt",
// 	"fahrradfahren",
// 	"kreuzheben",
// 	"kaeferUnilateral",
// 	"trizepsdipMitUnterstuetzung",
// 	"militaryPressMitErhoehtenFuessen",
// 	"sturzflug",
// 	"engerLiegestuetzMitErhoehtenFuessen",
// 	"trizepsdip",
// 	"einbeinigeKniebeugeMitUnterstuetzung",
// 	"kistenSprung",
// 	"klimmzug",
// 	"beinTwist",
// 	"tiefeKniebeuge",
// 	"ironMike",
// 	"seitensprung",
// 	"einbeinigerHueftstrecker",
// 	"breiterSturzflug",
// 	"gestreckterBeinTwist",
// 	"einarmigerLiegestuetzMitErhoehtenHaenden",
// 	"hueftStrecker",
// 	"nackentrainerInBauchlage",
// 	"einarmigerLiegestuetz",
// 	"federnderLiegestuetz",
// 	"erhoehterTrizepsstrecker",
// 	"pistole",
// 	"knieenderBeinwechsel",
// 	"einarmigesTuerziehen",
// 	"seitlichesKnieoeffnenImStand",
// 	"schraegerVUp",
// 	"einbeinigeKniebeuge",
// 	"kreuzschritt",
// 	"gestrecktesHaengendesBeinheben",
// 	"storch",
// 	"handstandLiegestuetz",
// 	"klappmesser",
// 	"einarmigerLiegestuetzMitErhoehtenFuessen",
// 	"bergsteiger",
// 	"gekreuzterBergsteiger",
// 	"pointer",
// 	"vorgebeugtesSeitlichesSchulterheben",
// }

// type OutputL []Output
// type Output struct {
// 	page int
// 	str  string
// }

// func (o OutputL) Len() int           { return len(o) }
// func (o OutputL) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
// func (o OutputL) Less(i, j int) bool { return o[i].page < o[j].page }

// func TestTrainingEnum(t *testing.T) {
// 	// sort.Sort(ExCollection(exesList))
// 	var varlist []Output
// 	for i, e := range exesList {
// 		o := Output{e.Page, fmt.Sprintf(" %s = exercise{\"%s\",%d,%d,%d}", varNames[i], e.Name, e.Type, e.Page, e.Difficulty)}
// 		varlist = append(varlist, o)
// 	}
// 	sort.Sort(OutputL(varlist))
// 	for _, e := range varlist {
// 		t.Log(e.str)
// 	}
// }
