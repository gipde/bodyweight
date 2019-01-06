package main

type Exercise struct {
	Number     int          `json:"number"`
	Name       string       `json:"name"`
	Type       TrainingType `json:"type"`
	Page       int          `json:"page"`
	Comment    string       `json:"comment"`
	Difficulty int          `json:"difficulty`
}

type TExercise struct {
	Exercise ExerciseEnum `json:"exercise"`
	Note     string       `json:"note"`
}

type TrainingDay struct {
	Method        TrainingMethod `json:"method"`
	Exercises     [][]TExercise  `json:"exercies"`
	MinutesNeeded int            `json:"minutesNeeded"`
}

type TrainingWeek struct {
	TrainingDays []TrainingDay `json:"trainingDays"`
	Description  string        `json:"description"`
}

type TrainingMethod int

const (
	STUFENINTERVALL TrainingMethod = iota
	INTERVALLSATZ
	SUPERSATZ
	ZIRKELINTERVALL
	HOCHINTENSITAETSSATZ
)

var trainingMethods = []string{
	"Stufenintervall",
	"Intervallsatz",
	"Supersatz",
	"Zirkelintervall",
	"Hochintensitätssatz",
}

func (t TrainingMethod) name() string {
	return trainingMethods[t]
}

// gehört zur Übung
type TrainingType int

const (
	DRUECKEN TrainingType = iota
	ZIEHEN
	BEINE_UND_GESAESS
	CORE
	GANZKOERPER
)

var trainingTypes = []string{
	"Drücken",
	"Ziehen",
	"Beine und Gesäß",
	"Core",
	"Ganzkörperlich",
}

func (t TrainingType) name() string {
	return trainingTypes[t]
}

type TrainingLevel int

const (
	BASISPROGRAM TrainingLevel = iota
	FIRSTCLASS
	MASTERCLASSC
	CHIEFCLASS
)

var trainingLevels = []string{
	"Basisprogramm für Einsteiger",
	"First Class",
	"Master Class",
	"Chief Class",
}

func (t TrainingLevel) name() string {
	return trainingLevels[t]
}

type ExerciseEnum int

// Varianten werden als eigene Übung eingetragen

const (
	LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN ExerciseEnum = iota
	TUERZIEHEN
	SEESTERN
	UMGEKEHRTES_BANKDRUECKEN
	AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL
	RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL
	ENGE_KNIEBEUGE
	SCHWIMMER
	SEITLICHER_AUSFALLSCHRITT
	SCHRAEGER_CRUNCH
	KLASSISCHER_LIEGESTUETZ
	MILITARY_PRESS_MIT_ERHOEHTEN_HAENDEN
	ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN
	KNIEBEUGE_IM_AUSFALLSCHRITT
	CURL_MIT_HANDTUCH
	BEINHEBER
	KAEFER
	RUSSISCHER_TWIST
	CRUNCH
	LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN
	LIEGESTUETZ_MIT_ABSTOSSEN
	MILITARY_PRESS
	BAERENGANG
	ENGER_LIEGESTUETZ
	GESPRUNGENE_KNIEBEUGE
	AUSFALLSCHRITT_NACH_VORNE_IM_WECHSEL
	KLIMMZUG_MIT_UNTERSTUEZUNG
	UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF
	TUERZIEHEN_IM_UNTERGRIFF
	V_UP
	KAEFER_IPSILATERAL
	CRUNCH_IT_UP
	HAENGENDES_BEINHEBEN
	POGO_SPRUNG
	HUEFTTWIST
	KNIEHEBER_IM_STEHEN
	BODYROCK
	VIER_PHASEN_LIEGESTUETZ
	GOOD_MORNING
	KAEFER_KONTRALATERAL
	SEITLICHES_HUEFTHEBEN
	AUFSTEHEN_AUS_DEM_EINBINIGEN_KNIESTAND
	AUSFALLSCHRITT
	FAHRRADFAHREN
	KREUZHEBEN
	KAEFER_UNILATERAL
	TRIZEPSDIP_MIT_UNTERSTUETZUNG
	MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN
	STURZFLUG
	ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN
	TRIZEPSDIP
	EINBINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG
	KISTENSPRUNG
	KLIMMZUG
	BEINTWIST
	TIEFE_KNIEBEUGE
	IRON_MIKE
	SEITSPRUNG
	EINBEINIGER_HUEFTSTRECKER
	BREITER_STURZFLUG
	GESTRECKTER_BEINTWIST
	EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN
	HUEFTSTRECKER
	NACKENTRAINER_IN_BAUCHLAGE
	EINARMIGER_LIEGESTUETZ
	FEDERNDER_LIEGESTUETZ
	ERHOEHTER_TRIZEPSSTRECKER
	PISTOLE
	KNIENDER_BEINWECHSEL
	EINARMIGES_TUERZIEHEN
	SEITLICHES_KNIEOEFFNEN_IM_STAND
	SCHRAEGER_V_UP
	EINBEINIGE_KNIEBEUGE
	KREUZSCHRITT
	GESTRECKTES_HAENGENDES_BEINHEBEN
	EINBEINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG
	STORCH
	HANDSTANDLIEGESTUETZ
	KLAPPMESSER
	EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN
	BERGSTEIGER
	GEKREUZTER_BERGSTEIGER
	POINTER
	VORGEBEUGTES_SEITLICHES_SCHULTERHEBEN
)

var Exes []Exercise = []Exercise{
	Exercise{
		Name:       "Liegestütz mit eröhten Händen",
		Type:       DRUECKEN,
		Page:       112,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Türziehen",
		Type:       ZIEHEN,
		Page:       145,
		Difficulty: 1,
	},

	Exercise{
		Name:       "Seestern",
		Type:       CORE,
		Page:       207,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Umgekehrtes Bankdrücken",
		Type:       ZIEHEN,
		Page:       147,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Ausfallschritt nach hinten im Wechsel",
		Type:       BEINE_UND_GESAESS,
		Page:       178,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Rumänisches Kruezheben auf einem Bein im Wechsel",
		Type:       BEINE_UND_GESAESS,
		Page:       176,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Enge Kniebeuge",
		Type:       BEINE_UND_GESAESS,
		Page:       183,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Schwimmer",
		Type:       BEINE_UND_GESAESS,
		Page:       174,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Seitlicher Ausfallschritt",
		Type:       BEINE_UND_GESAESS,
		Page:       179,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Schräger Crunch",
		Type:       BEINE_UND_GESAESS,
		Page:       215,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Klassischer Liegestütz",
		Type:       DRUECKEN,
		Page:       112,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Military Press mit erhöhten Händen",
		Type:       DRUECKEN,
		Page:       135,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Enger Liegestütz mit erhöhten Händen",
		Type:       DRUECKEN,
		Page:       130,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Kniebeuge im Ausfallschritt",
		Type:       BEINE_UND_GESAESS,
		Page:       178,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Curl mit Handtuch",
		Type:       ZIEHEN,
		Page:       159,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Beinheber",
		Type:       CORE,
		Page:       217,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Käfer",
		Type:       CORE,
		Page:       204,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Russischer Twist",
		Type:       CORE,
		Page:       213,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Crunch",
		Type:       CORE,
		Page:       215,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Liegestütz mit erhöhten Füßen",
		Type:       DRUECKEN,
		Page:       113,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Liegestütz mit Abstoßen",
		Type:       DRUECKEN,
		Page:       117,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Military Press",
		Type:       DRUECKEN,
		Page:       134,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Bärengang",
		Type:       DRUECKEN,
		Page:       110,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Enger Liegestütz",
		Type:       DRUECKEN,
		Page:       130,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Gesprungene Kniebeuge",
		Type:       BEINE_UND_GESAESS,
		Page:       191,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Ausfallschritt nach vorn im Wechsel",
		Type:       BEINE_UND_GESAESS,
		Page:       177,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Klimmzug mit Unterstützung",
		Type:       ZIEHEN,
		Page:       152,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Umgekehrtes Bankdrücken im Untergriff",
		Type:       ZIEHEN,
		Page:       148,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Türziehen",
		Type:       ZIEHEN,
		Page:       145,
		Difficulty: 3,
	},
	Exercise{
		Name:       "V-Up",
		Type:       CORE,
		Page:       219,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Käfer ipsilateral",
		Type:       CORE,
		Page:       204,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Crunch It Up",
		Type:       CORE,
		Page:       214,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Hängendes Beinheben",
		Type:       CORE,
		Page:       223,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Pogo Sprung",
		Type:       BEINE_UND_GESAESS,
		Page:       199,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Hüfttwist",
		Type:       CORE,
		Page:       206,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Knieheber im Stehen",
		Type:       GANZKOERPER,
		Page:       228,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Bodyrock",
		Type:       CORE,
		Page:       208,
		Difficulty: 1,
	},
	Exercise{
		Name:       "4 Phasen Liegestütz",
		Type:       GANZKOERPER,
		Page:       229,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Good Morning",
		Type:       BEINE_UND_GESAESS,
		Page:       164,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Käfer kontralateral",
		Type:       CORE,
		Page:       204,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Seitliches Hüftheben",
		Type:       CORE,
		Page:       204,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Aufstehen aus dem einbeinigen Kniestand",
		Type:       BEINE_UND_GESAESS,
		Page:       169,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Ausfallschritt",
		Type:       BEINE_UND_GESAESS,
		Page:       177,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Fahrradfahren",
		Type:       CORE,
		Page:       218,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Kreuzheben",
		Type:       ZIEHEN,
		Page:       168,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Käfer unilateral",
		Type:       CORE,
		Page:       204,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Trizepsdip mit Unterstützung",
		Type:       DRUECKEN,
		Page:       132,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Military Press mit erhöhten Füßen",
		Type:       DRUECKEN,
		Page:       135,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Sturzflug",
		Type:       DRUECKEN,
		Page:       120,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Enger Liegestütz mit erhöhten Füßen",
		Type:       DRUECKEN,
		Page:       130,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Trizepsdip",
		Type:       DRUECKEN,
		Page:       130,
		Difficulty: 4,
	},
	Exercise{
		Name:       "Einbeinige Kniebeuge mit Unterstützung im Wechsel",
		Type:       BEINE_UND_GESAESS,
		Page:       187,
		Difficulty: 4,
	},
	Exercise{
		Name:       "Kistensprung",
		Type:       BEINE_UND_GESAESS,
		Page:       192,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Klimmzug",
		Type:       ZIEHEN,
		Page:       150,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Beintwist",
		Type:       CORE,
		Page:       221,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Tiefe Kniebeuge",
		Type:       BEINE_UND_GESAESS,
		Page:       221,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Iron Mike",
		Type:       BEINE_UND_GESAESS,
		Page:       195,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Seitsprung",
		Type:       BEINE_UND_GESAESS,
		Page:       194,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Einbeiniger Hüftstrecker",
		Type:       BEINE_UND_GESAESS,
		Page:       194,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Breiter Sturzflug",
		Type:       BEINE_UND_GESAESS,
		Page:       121,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Gestreckter Beintwist",
		Type:       CORE,
		Page:       221,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Einarmiger Liegestütz mit erhöhten Händen im Wechsel",
		Type:       DRUECKEN,
		Page:       125,
		Difficulty: 4,
	},
	Exercise{
		Name:       "Hüftstrecker",
		Type:       BEINE_UND_GESAESS,
		Page:       181,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Nackentrainer in Bauchlage",
		Type:       CORE,
		Page:       227,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Einarmiger Liegestütz",
		Type:       DRUECKEN,
		Page:       124,
		Difficulty: 4,
	},
	Exercise{
		Name:       "Federnder Liegestütz",
		Type:       DRUECKEN,
		Page:       117,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Erhöhter Trizepsstrecker",
		Type:       DRUECKEN,
		Page:       131,
		Difficulty: 4,
	},
	Exercise{
		Name:       "Pistole",
		Type:       BEINE_UND_GESAESS,
		Page:       188,
		Difficulty: 4,
	},
	Exercise{
		Name:       "Kniender Beinwechsel",
		Type:       BEINE_UND_GESAESS,
		Page:       175,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Einarmiges Türziehen",
		Type:       ZIEHEN,
		Page:       146,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Seitliches Knieöffnen im Stand",
		Type:       BEINE_UND_GESAESS,
		Page:       166,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Schräger V-Up",
		Type:       CORE,
		Page:       220,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Einbeinige Kniebeuge im Wechsel",
		Type:       BEINE_UND_GESAESS,
		Page:       188,
		Difficulty: 4,
	},
	Exercise{
		Name:       "Kreuzschritt",
		Type:       BEINE_UND_GESAESS,
		Page:       167,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Gestrecktes hängendes Beinheben",
		Type:       CORE,
		Page:       223,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Einbeinige Kniebeuge mit Unterstützung im Wechsel",
		Type:       BEINE_UND_GESAESS,
		Page:       187,
		Difficulty: 4,
	},
	Exercise{
		Name:       "Storch",
		Type:       BEINE_UND_GESAESS,
		Page:       165,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Handstandliegestütz",
		Type:       BEINE_UND_GESAESS,
		Page:       138,
		Difficulty: 1,
	},
	Exercise{
		Name:       "Klappmesser",
		Type:       CORE,
		Page:       222,
		Difficulty: 3,
	},
	Exercise{
		Name:       "Einarmiger Liegestütz mit erhöhten Füßen",
		Type:       DRUECKEN,
		Page:       126,
		Difficulty: 4,
	},
	Exercise{
		Name:       "Bergsteiger",
		Type:       CORE,
		Page:       202,
		Difficulty: 2,
	},
	Exercise{
		Name:       "gekreuzter Bergsteiger",
		Type:       CORE,
		Page:       203,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Pointer",
		Type:       BEINE_UND_GESAESS,
		Page:       174,
		Difficulty: 2,
	},
	Exercise{
		Name:       "Vorgebeugtes seitliches Schulterheben",
		Type:       BEINE_UND_GESAESS,
		Page:       156,
		Difficulty: 1,
	},
}

var Trainings = []TrainingWeek{
	TrainingWeek{
		Description: "Muskuläre Ausdauer",
		TrainingDays: []TrainingDay{
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: SEESTERN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
					},
					[]TExercise{
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: TUERZIEHEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "Füße erhöht"},
					},
				},
			},
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: ENGE_KNIEBEUGE, Note: "Arme in Vorhalte"},
						TExercise{Exercise: SCHWIMMER},
					},
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: KAEFER_UNILATERAL, Note: "Beine gestreckt"},
					},
					[]TExercise{
						TExercise{Exercise: EINBINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG},
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: HUEFTSTRECKER},
						TExercise{Exercise: KAEFER_UNILATERAL, Note: "Beine gebeugt"},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE},
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: EINBEINIGER_HUEFTSTRECKER},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE},
					},
				},
			},
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: SEESTERN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
					},
					[]TExercise{
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: TUERZIEHEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG,
							Note: "ohne Unterstützung in der Negativphase"},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: TUERZIEHEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN},
					},
				},
			},
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: ENGE_KNIEBEUGE},
						TExercise{Exercise: SCHRAEGER_CRUNCH},
					},
					[]TExercise{
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: ENGE_KNIEBEUGE, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUSSISCHER_TWIST},
					},
					[]TExercise{
						TExercise{Exercise: EINBINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: BEINTWIST},
					},
					[]TExercise{
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
						TExercise{Exercise: KISTENSPRUNG},
						TExercise{Exercise: STORCH},
						TExercise{Exercise: BEINTWIST},
					},
				},
			},
		},
	},
	TrainingWeek{
		Description: "Muskuläre Ausdauer",
		TrainingDays: []TrainingDay{
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: SEESTERN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
					},
					[]TExercise{
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: TUERZIEHEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "Füße erhöht"},
					},
				},
			},
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: ENGE_KNIEBEUGE, Note: "Arme in Vorhalte"},
						TExercise{Exercise: SCHWIMMER},
					},
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: KAEFER_UNILATERAL, Note: "Beine gestreckt"},
					},
					[]TExercise{
						TExercise{Exercise: EINBINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG},
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: HUEFTSTRECKER},
						TExercise{Exercise: KAEFER_UNILATERAL, Note: "Beine gebeugt"},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE},
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: EINBEINIGER_HUEFTSTRECKER},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE},
					},
				},
			},
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: SEESTERN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
					},
					[]TExercise{
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: TUERZIEHEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG,
							Note: "ohne Unterstützung in der Negativphase"},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: TUERZIEHEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN},
					},
				},
			},
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: ENGE_KNIEBEUGE},
						TExercise{Exercise: SCHRAEGER_CRUNCH},
					},
					[]TExercise{
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: ENGE_KNIEBEUGE, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUSSISCHER_TWIST},
					},
					[]TExercise{
						TExercise{Exercise: EINBINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: BEINTWIST},
					},
					[]TExercise{
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
						TExercise{Exercise: KISTENSPRUNG},
						TExercise{Exercise: STORCH},
						TExercise{Exercise: BEINTWIST},
					},
				},
			},
		},
	},
	TrainingWeek{
		Description: "Kraft",
		TrainingDays: []TrainingDay{
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: SEESTERN},
					},
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: ENGER_LIEGESTUETZ},
						TExercise{Exercise: TRIZEPSDIP_MIT_UNTERSTUETZUNG, Note: "Beine nach hinten anwinklen und Fußspitzen auf einen Stuhl absetzen um das Hochdrücken zu erleichtern"},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: TRIZEPSDIP_MIT_UNTERSTUETZUNG},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ERHOEHTER_TRIZEPSSTRECKER, Note: "Hände etwa hüfthoch"},
					},
				},
			},
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT},
						TExercise{Exercise: ENGE_KNIEBEUGE,
							Note: "Arme in T-Halte mit 1-3 Sek. Haltezeit am tiefsten Punkt",
						},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen"},
					},
					[]TExercise{
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen"},
					},
					[]TExercise{
						TExercise{Exercise: EINBINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG},
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						TExercise{Exercise: EINBEINIGER_HUEFTSTRECKER,
							Note: "auf einem Kissen"},
					},
					[]TExercise{
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "mit 1-3 Sek. Pause",
						},
						TExercise{Exercise: IRON_MIKE},
					},
				},
			},
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN, Note: "Im Untergriff"},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG,
							Note: "ohne Unterstützung in der Negativphase"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN, Note: "Im Untergriff"},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: TUERZIEHEN, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG, Note: "mit 1-3 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF, Note: "Füße erhöht"},
						TExercise{Exercise: TUERZIEHEN, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: BEINHEBER},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gebeugt"},
						TExercise{Exercise: RUSSISCHER_TWIST},
						TExercise{Exercise: CRUNCH},
					},
					[]TExercise{
						TExercise{Exercise: BEINHEBER, Note: "mit gekreuzten Armen auf der Brust"},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "iBeine gestreckt"},
						TExercise{Exercise: FAHRRADFAHREN},
						TExercise{Exercise: KAEFER_KONTRALATERAL},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN, Note: "Beine gestreckt und parallel zum Boden"},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE},
						TExercise{Exercise: V_UP},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gestreckt"},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN, Note: "Beine gestreckt bis ganz nach oben"},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE},
						TExercise{Exercise: KLAPPMESSER},
						TExercise{Exercise: KAEFER_KONTRALATERAL, Note: "Beine gestreckt"},
					},
				},
			},
		},
	},
	TrainingWeek{
		Description: "Kraft",
		TrainingDays: []TrainingDay{
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: SEESTERN},
					},
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: ENGER_LIEGESTUETZ},
						TExercise{Exercise: TRIZEPSDIP_MIT_UNTERSTUETZUNG, Note: "Beine nach hinten anwinklen und Fußspitzen auf einen Stuhl absetzen um das Hochdrücken zu erleichtern"},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: TRIZEPSDIP_MIT_UNTERSTUETZUNG},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ERHOEHTER_TRIZEPSSTRECKER, Note: "Hände etwa hüfthoch"},
					},
				},
			},
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT},
						TExercise{Exercise: ENGE_KNIEBEUGE,
							Note: "Arme in T-Halte mit 1-3 Sek. Haltezeit am tiefsten Punkt",
						},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen"},
					},
					[]TExercise{
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen"},
					},
					[]TExercise{
						TExercise{Exercise: EINBINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG},
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						TExercise{Exercise: EINBEINIGER_HUEFTSTRECKER,
							Note: "auf einem Kissen"},
					},
					[]TExercise{
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "mit 1-3 Sek. Pause",
						},
						TExercise{Exercise: IRON_MIKE},
					},
				},
			},
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN, Note: "Im Untergriff"},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG,
							Note: "ohne Unterstützung in der Negativphase"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN, Note: "Im Untergriff"},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: TUERZIEHEN, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG, Note: "mit 1-3 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF, Note: "Füße erhöht"},
						TExercise{Exercise: TUERZIEHEN, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: BEINHEBER},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gebeugt"},
						TExercise{Exercise: RUSSISCHER_TWIST},
						TExercise{Exercise: CRUNCH},
					},
					[]TExercise{
						TExercise{Exercise: BEINHEBER, Note: "mit gekreuzten Armen auf der Brust"},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "iBeine gestreckt"},
						TExercise{Exercise: FAHRRADFAHREN},
						TExercise{Exercise: KAEFER_KONTRALATERAL},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN, Note: "Beine gestreckt und parallel zum Boden"},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE},
						TExercise{Exercise: V_UP},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gestreckt"},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN, Note: "Beine gestreckt bis ganz nach oben"},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE},
						TExercise{Exercise: KLAPPMESSER},
						TExercise{Exercise: KAEFER_KONTRALATERAL, Note: "Beine gestreckt"},
					},
				},
			},
		},
	},
	TrainingWeek{
		Description: "Powerblock",
		TrainingDays: []TrainingDay{
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: BAERENGANG},
						TExercise{Exercise: ENGER_LIEGESTUETZ},
						TExercise{Exercise: SEESTERN},
					},
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: TRIZEPSDIP},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ},
						TExercise{Exercise: FEDERNDER_LIEGESTUETZ},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ERHOEHTER_TRIZEPSSTRECKER},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: FEDERNDER_LIEGESTUETZ},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ERHOEHTER_TRIZEPSSTRECKER, Note: "Hände etwa hüfthoch"},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
					},
				},
			},
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE},
						TExercise{Exercise: AUSFALLSCHRITT_NACH_VORNE_IM_WECHSEL,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen"},
						TExercise{Exercise: ENGE_KNIEBEUGE,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					[]TExercise{
						TExercise{Exercise: EINBINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: KISTENSPRUNG},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE},
						TExercise{Exercise: KISTENSPRUNG},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
						TExercise{Exercise: SEITSPRUNG},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE, Note: "Mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: KISTENSPRUNG},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
						TExercise{Exercise: SEITSPRUNG},
					},
				},
			},
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG,
							Note: "ohne Unterstützung in der Negativphase"},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: TUERZIEHEN,
							Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: TUERZIEHEN,
							Note: "mit 4-6 Sek. Haltezeit bei der Kontraktionam höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF, Note: "Füße erhöht"},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "Füße erhöht"},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN, Note: "mit 1-3 Sek. Haltezeit bei der kontraktion am Höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF, Note: "Füße erhöht"},
						TExercise{Exercise: KLIMMZUG, Note: "Bis zum Brustbein hochziehen"},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: V_UP},
						TExercise{Exercise: RUSSISCHER_TWIST},
						TExercise{Exercise: KAEFER_IPSILATERAL,
							Note: "Beine gestreckt"},
						TExercise{Exercise: CRUNCH_IT_UP},
						TExercise{Exercise: HAENGENDES_BEINHEBEN},
						TExercise{Exercise: BEINHEBER},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN},
						TExercise{Exercise: BEINTWIST},
						TExercise{Exercise: HUEFTTWIST},
						TExercise{Exercise: KAEFER_UNILATERAL, Note: "Beine gestreckt"},
						TExercise{Exercise: V_UP},
						TExercise{Exercise: RUSSISCHER_TWIST},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN},
						TExercise{Exercise: FAHRRADFAHREN},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gebeugt"},
						TExercise{Exercise: V_UP},
						TExercise{Exercise: BEINTWIST},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt, Beine gestreckt bis ganz nach oben"},
						TExercise{Exercise: FAHRRADFAHREN, Note: "langsam, 2 Sek. für das Heranziehen jedes Knies"},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE, Note: "mit 4-6 Sek. Pause"},
						TExercise{Exercise: KAEFER_UNILATERAL, Note: "Beine gestreckt"},
						TExercise{Exercise: KLAPPMESSER},
						TExercise{Exercise: BEINTWIST},
					},
				},
			},
		},
	},
	TrainingWeek{
		Description: "Powerblock",
		TrainingDays: []TrainingDay{
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: BAERENGANG},
						TExercise{Exercise: ENGER_LIEGESTUETZ},
						TExercise{Exercise: SEESTERN},
					},
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: TRIZEPSDIP},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ},
						TExercise{Exercise: FEDERNDER_LIEGESTUETZ},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ERHOEHTER_TRIZEPSSTRECKER},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: FEDERNDER_LIEGESTUETZ},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ERHOEHTER_TRIZEPSSTRECKER, Note: "Hände etwa hüfthoch"},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
					},
				},
			},
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE},
						TExercise{Exercise: AUSFALLSCHRITT_NACH_VORNE_IM_WECHSEL,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen"},
						TExercise{Exercise: ENGE_KNIEBEUGE,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					[]TExercise{
						TExercise{Exercise: EINBINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: KISTENSPRUNG},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE},
						TExercise{Exercise: KISTENSPRUNG},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
						TExercise{Exercise: SEITSPRUNG},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE, Note: "Mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: KISTENSPRUNG},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
						TExercise{Exercise: SEITSPRUNG},
					},
				},
			},
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG,
							Note: "ohne Unterstützung in der Negativphase"},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: TUERZIEHEN,
							Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: TUERZIEHEN,
							Note: "mit 4-6 Sek. Haltezeit bei der Kontraktionam höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF, Note: "Füße erhöht"},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "Füße erhöht"},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN, Note: "mit 1-3 Sek. Haltezeit bei der kontraktion am Höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF, Note: "Füße erhöht"},
						TExercise{Exercise: KLIMMZUG, Note: "Bis zum Brustbein hochziehen"},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: V_UP},
						TExercise{Exercise: RUSSISCHER_TWIST},
						TExercise{Exercise: KAEFER_IPSILATERAL,
							Note: "Beine gestreckt"},
						TExercise{Exercise: CRUNCH_IT_UP},
						TExercise{Exercise: HAENGENDES_BEINHEBEN},
						TExercise{Exercise: BEINHEBER},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN},
						TExercise{Exercise: BEINTWIST},
						TExercise{Exercise: HUEFTTWIST},
						TExercise{Exercise: KAEFER_UNILATERAL, Note: "Beine gestreckt"},
						TExercise{Exercise: V_UP},
						TExercise{Exercise: RUSSISCHER_TWIST},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN},
						TExercise{Exercise: FAHRRADFAHREN},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gebeugt"},
						TExercise{Exercise: V_UP},
						TExercise{Exercise: BEINTWIST},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt, Beine gestreckt bis ganz nach oben"},
						TExercise{Exercise: FAHRRADFAHREN, Note: "langsam, 2 Sek. für das Heranziehen jedes Knies"},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE, Note: "mit 4-6 Sek. Pause"},
						TExercise{Exercise: KAEFER_UNILATERAL, Note: "Beine gestreckt"},
						TExercise{Exercise: KLAPPMESSER},
						TExercise{Exercise: BEINTWIST},
					},
				},
			},
		},
	},
	TrainingWeek{
		Description: "Wechselblock",
		TrainingDays: []TrainingDay{
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: SEESTERN},
					},
					[]TExercise{
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ},
						TExercise{Exercise: ENGER_LIEGESTUETZ},
						TExercise{Exercise: SEESTERN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: TRIZEPSDIP},
						TExercise{Exercise: SEESTERN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: TRIZEPSDIP, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
					},
				},
			},
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE},
						TExercise{Exercise: AUSFALLSCHRITT_NACH_VORNE_IM_WECHSEL,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen und mit 1-3 Sek. Haltezeit in der Mitte"},
						TExercise{Exercise: POGO_SPRUNG},
					},
					[]TExercise{
						TExercise{Exercise: EINBINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen und mit 1-3 Sek. Haltezeit in der Mitte"},
						TExercise{Exercise: KISTENSPRUNG},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE},
						TExercise{Exercise: SEITLICHES_KNIEOEFFNEN_IM_STAND},
						TExercise{Exercise: KISTENSPRUNG},
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
						TExercise{Exercise: SEITSPRUNG},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE, Note: "im Wechsel mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: KISTENSPRUNG},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
						TExercise{Exercise: SEITSPRUNG},
					},
				},
			},
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG,
							Note: "ohne Unterstützung in der Negativphase"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: TUERZIEHEN, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am Höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG, Note: "mit 1-3 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF, Note: "Füße erhöht"},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			TrainingDay{
				Method: HOCHINTENSITAETSSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: RUSSISCHER_TWIST},
						TExercise{Exercise: HUEFTTWIST},
						TExercise{Exercise: KNIEHEBER_IM_STEHEN},
					},
					[]TExercise{
						TExercise{Exercise: RUSSISCHER_TWIST},
						TExercise{Exercise: BODYROCK},
						TExercise{Exercise: TIEFE_KNIEBEUGE, Note: "Hände hinter dem Kopf"},
					},
					[]TExercise{
						TExercise{Exercise: SCHRAEGER_V_UP, Note: "4 Sätze pro Seite"},
						TExercise{Exercise: BODYROCK},
						TExercise{Exercise: TIEFE_KNIEBEUGE, Note: "Arme in Streamline Position"},
					},
					[]TExercise{
						TExercise{Exercise: V_UP},
						TExercise{Exercise: SCHRAEGER_V_UP, Note: "nach jedem Satz Seitenwechsel, insgesamt 4 Sätze"},
						TExercise{Exercise: RUSSISCHER_TWIST},
						TExercise{Exercise: BERGSTEIGER},
					},
				},
			},
			TrainingDay{
				Method: ZIRKELINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL,
							Note: "10 Wiederholungen pro Seite"},
						TExercise{Exercise: TUERZIEHEN,
							Note: "8 Wiederholungen"},
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ,
							Note: "6 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN,
							Note: "6 Wiederholungen"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT,
							Note: "12 Wiederholungen"},
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ,
							Note: "8 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: EINBEINIGE_KNIEBEUGE,
							Note: "12 Wiederholungen"},
						TExercise{Exercise: STURZFLUG,
							Note: "6 Wiederholungen"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN,
							Note: "8 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE,
							Note: "12 Wiederholungen"},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ,
							Note: "6 Wiederholungen"},
						TExercise{Exercise: KLIMMZUG,
							Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
	TrainingWeek{
		Description: "Wechselblock",
		TrainingDays: []TrainingDay{
			TrainingDay{
				Method: HOCHINTENSITAETSSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN, Note: "brusthoch abgestützt"},
						TExercise{Exercise: BODYROCK},
						TExercise{Exercise: VIER_PHASEN_LIEGESTUETZ},
					},
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: SEESTERN},
						TExercise{Exercise: TIEFE_KNIEBEUGE},
					},
					[]TExercise{
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
						TExercise{Exercise: TIEFE_KNIEBEUGE, Note: "Arme in T-Halte"},
					},
					[]TExercise{
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
					},
				},
			},
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
						TExercise{Exercise: ENGE_KNIEBEUGE, Note: "Arme in Vorhalte mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: GOOD_MORNING, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
					},
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL},
					},
					[]TExercise{
						TExercise{Exercise: EINBEINIGE_KNIEBEUGE},
						TExercise{Exercise: KREUZSCHRITT},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: HUEFTSTRECKER},
					},
					[]TExercise{
						TExercise{Exercise: EINBEINIGE_KNIEBEUGE},
						TExercise{Exercise: SEITLICHES_KNIEOEFFNEN_IM_STAND},
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: EINBEINIGER_HUEFTSTRECKER},
					},
				},
			},
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: TUERZIEHEN, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "mit gebeugten Knien"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF, Note: "Beine gestreckt"},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: TUERZIEHEN, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF, Note: "Füße erhöht"},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "Füße erhöht"},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF, Note: "Füße erhöht"},
						TExercise{Exercise: KLIMMZUG, Note: "bis zum Brustbein hochziehen"},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: BEINHEBER},
						TExercise{Exercise: KAEFER_KONTRALATERAL, Note: "Beine gebeugt"},
						TExercise{Exercise: RUSSISCHER_TWIST},
						TExercise{Exercise: SEITLICHES_HUEFTHEBEN},
					},
					[]TExercise{
						TExercise{Exercise: BEINHEBER, Note: "mit gekreuzten Armen auf der Brust"},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gestreckt"},
						TExercise{Exercise: FAHRRADFAHREN},
						TExercise{Exercise: KAEFER_KONTRALATERAL, Note: "Beine gestreckt"},
					},
					[]TExercise{
						TExercise{Exercise: GESTRECKTES_HAENGENDES_BEINHEBEN},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE},
						TExercise{Exercise: V_UP},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gestreckt"},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN, Note: "Beine gestreckt, Füße bis zu den Händen"},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE, Note: "mit 1-3 Sek. Pause"},
						TExercise{Exercise: GEKREUZTER_BERGSTEIGER},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gebeugt"},
					},
				},
			},
			TrainingDay{
				Method: ZIRKELINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL, Note: "10 Wiederholungen pro Seite"},
						TExercise{Exercise: TUERZIEHEN, Note: "8 Wiederholungen"},
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ, Note: "6 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "6 Wiederholungen"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT, Note: "12 Wiederholungen"},
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ, Note: "8 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: EINBEINIGE_KNIEBEUGE, Note: "12 Wiederholungen pro Seite"},
						TExercise{Exercise: STURZFLUG, Note: "6 Wiederholungen"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "8 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE, Note: "6 Wiederholungen pro Seite"},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ, Note: "6 Wiederholungen"},
						TExercise{Exercise: KLIMMZUG, Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
	TrainingWeek{
		Description: "Wechselblock",
		TrainingDays: []TrainingDay{
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
					},
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: SEESTERN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ_MIT_ERHOEHTEN_HAENDEN},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: TRIZEPSDIP},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ERHOEHTER_TRIZEPSSTRECKER, Note: "Hände etwa hüfthoch"},
					},
				},
			},
			TrainingDay{
				Method: HOCHINTENSITAETSSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUFSTEHEN_AUS_DEM_EINBINIGEN_KNIESTAND, Note: "Arme in Vorhalte"},
						TExercise{Exercise: AUSFALLSCHRITT},
						TExercise{Exercise: GOOD_MORNING},
					},
					[]TExercise{
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: SEITSPRUNG},
						TExercise{Exercise: TIEFE_KNIEBEUGE, Note: "Arme in T-Halte"},
					},
					[]TExercise{
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: SEITSPRUNG},
						TExercise{Exercise: TIEFE_KNIEBEUGE, Note: "Hände an den Hüften"},
					},
					[]TExercise{
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: SEITSPRUNG},
						TExercise{Exercise: POINTER},
						TExercise{Exercise: ENGE_KNIEBEUGE, Note: "Hände hinter dem Kopf"},
					},
				},
			},
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF},
					},
					[]TExercise{
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
						TExercise{Exercise: TUERZIEHEN_IM_UNTERGRIFF},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
						TExercise{Exercise: TUERZIEHEN},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN_IM_UNTERGRIFF},
						TExercise{Exercise: EINARMIGES_TUERZIEHEN},
					},
				},
			},
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: V_UP},
						TExercise{Exercise: RUSSISCHER_TWIST},
						TExercise{Exercise: KAEFER_KONTRALATERAL, Note: "Beine gestreckt"},
						TExercise{Exercise: SEITLICHES_HUEFTHEBEN},
						TExercise{Exercise: FAHRRADFAHREN},
						TExercise{Exercise: BEINHEBER},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN},
						TExercise{Exercise: BEINTWIST},
						TExercise{Exercise: EINBEINIGER_HUEFTSTRECKER},
						TExercise{Exercise: KAEFER_UNILATERAL, Note: "Beine gestreckt"},
						TExercise{Exercise: V_UP},
						TExercise{Exercise: RUSSISCHER_TWIST},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN},
						TExercise{Exercise: FAHRRADFAHREN},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE, Note: "mit 4-6 Sek. Pause"},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gebeugt"},
						TExercise{Exercise: V_UP},
						TExercise{Exercise: BEINTWIST},
					},
					[]TExercise{
						TExercise{Exercise: HAENGENDES_BEINHEBEN, Note: "Beine gestreckt, Füße bis zu den Händen mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						TExercise{Exercise: FAHRRADFAHREN, Note: "12 langsame Wdh. pro Seite)"},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE, Note: "mit 4-6 Sek. Pause"},
						TExercise{Exercise: KAEFER_KONTRALATERAL, Note: "Beine gebeugt"},
						TExercise{Exercise: KLAPPMESSER},
						TExercise{Exercise: BEINTWIST, Note: "6 langsame Wiederholungen pro Seite"},
					},
				},
			},
			TrainingDay{
				Method: ZIRKELINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL, Note: "10 Wiederholungen pro Seite"},
						TExercise{Exercise: TUERZIEHEN, Note: "8 Wiederholungen"},
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ, Note: "6 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "6 Wiederholungen pro Seite"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT, Note: "12 Wiederholungen"},
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ, Note: "8 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: EINBEINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG, Note: "12 Wiederholungen pro Seite"},
						TExercise{Exercise: STURZFLUG, Note: "6 Wiederholungen"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "8 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE, Note: "6 Wiederholungen pro Seite"},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ, Note: "6 Wiederholungen"},
						TExercise{Exercise: KLIMMZUG, Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
	TrainingWeek{
		Description: "Wechselblock",
		TrainingDays: []TrainingDay{
			TrainingDay{
				Method: SUPERSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
						TExercise{Exercise: MILITARY_PRESS},
						TExercise{Exercise: BAERENGANG},
						TExercise{Exercise: ENGER_LIEGESTUETZ},
						TExercise{Exercise: SEESTERN},
					},
					[]TExercise{
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: BREITER_STURZFLUG},
						TExercise{Exercise: ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: SEESTERN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ},
						TExercise{Exercise: FEDERNDER_LIEGESTUETZ},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ERHOEHTER_TRIZEPSSTRECKER},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
					},
					[]TExercise{
						TExercise{Exercise: EINARMIGER_LIEGESTUETZ},
						TExercise{Exercise: FEDERNDER_LIEGESTUETZ},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ},
						TExercise{Exercise: STURZFLUG, Note: "mit Haltezeit wenn die Brust am tiefsten Punkt der Kontraktion zwischen den Händen ist"},
						TExercise{Exercise: ERHOEHTER_TRIZEPSSTRECKER, Note: "Trizepsstrecker, hände kniehoch"},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
					},
				},
			},
			TrainingDay{
				Method: INTERVALLSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT},
						TExercise{Exercise: ENGE_KNIEBEUGE, Note: "Arme in T-Halte, mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL, Note: "auf einem Kissen"},
					},
					[]TExercise{
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL, Note: "auf einem Kissen"},
					},
					[]TExercise{
						TExercise{Exercise: EINBEINIGE_KNIEBEUGE},
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: EINBEINIGER_HUEFTSTRECKER},
					},
					[]TExercise{
						TExercise{Exercise: EINBEINIGE_KNIEBEUGE, Note: "knieend"},
						TExercise{Exercise: KNIEBEUGE_IM_AUSFALLSCHRITT, Note: "dabei am tiefsten Punkt einen Rucksack über den Kopf stemen"},
						TExercise{Exercise: IRON_MIKE},
						TExercise{Exercise: KNIENDER_BEINWECHSEL},
					},
				},
			},
			TrainingDay{
				Method: HOCHINTENSITAETSSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: TUERZIEHEN, Note: "Füße sind hinter den Händen platziert, gehen Sie dafür einem Schritt zurürck"},
						TExercise{Exercise: KREUZHEBEN},
						TExercise{Exercise: CURL_MIT_HANDTUCH},
					},
					[]TExercise{
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: TIEFE_KNIEBEUGE, Note: "Arme in Streamline Position"},
					},
					[]TExercise{
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: TIEFE_KNIEBEUGE, Note: "Hände hinter dem Kopf"},
					},
					[]TExercise{
						TExercise{Exercise: KLIMMZUG_MIT_UNTERSTUEZUNG},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN},
						TExercise{Exercise: TUERZIEHEN},
						TExercise{Exercise: VORGEBEUGTES_SEITLICHES_SCHULTERHEBEN},
					},
				},
			},
			TrainingDay{
				Method: STUFENINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: CRUNCH_IT_UP},
						TExercise{Exercise: KAEFER_IPSILATERAL},
						TExercise{Exercise: BEINHEBER},
						TExercise{Exercise: KAEFER_KONTRALATERAL},
					},
					[]TExercise{
						TExercise{Exercise: FAHRRADFAHREN},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gestreckt"},
						TExercise{Exercise: BODYROCK, Note: "mit vorwärtsgreifen"},
						TExercise{Exercise: GESTRECKTER_BEINTWIST},
					},
					[]TExercise{
						TExercise{Exercise: SCHRAEGER_V_UP, Note: "gebeugten Beinen"},
						TExercise{Exercise: KREUZHEBEN},
						TExercise{Exercise: BEINHEBER, Note: "mit gekreuzten Armen auf der Brust"},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gebeugt"},
					},
					[]TExercise{
						TExercise{Exercise: KLAPPMESSER},
						TExercise{Exercise: NACKENTRAINER_IN_BAUCHLAGE},
						TExercise{Exercise: SCHRAEGER_V_UP, Note: "mit gekreuzten Armen auf der Brust"},
						TExercise{Exercise: KAEFER_IPSILATERAL, Note: "Beine gestreckt"},
					},
				},
			},
			TrainingDay{
				Method: ZIRKELINTERVALL,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUSFALLSCHRITT_NACH_HINTEN_IM_WECHSEL, Note: "10 Wiederholungen pro Seite"},
						TExercise{Exercise: TUERZIEHEN, Note: "8 Wiederholungen"},
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ, Note: "6 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "6 Wiederholungen pro Seite"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT, Note: "12 Wiederholungen"},
						TExercise{Exercise: KLASSISCHER_LIEGESTUETZ, Note: "8 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: EINBEINIGE_KNIEBEUGE_MIT_UNTERSTUETZUNG, Note: "12 Wiederholungen pro Seite"},
						TExercise{Exercise: STURZFLUG, Note: "6 Wiederholungen"},
						TExercise{Exercise: UMGEKEHRTES_BANKDRUECKEN, Note: "8 Wiederholungen"},
					},
					[]TExercise{
						TExercise{Exercise: PISTOLE, Note: "6 Wiederholungen pro Seite"},
						TExercise{Exercise: HANDSTANDLIEGESTUETZ, Note: "6 Wiederholungen"},
						TExercise{Exercise: KLIMMZUG, Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
}
