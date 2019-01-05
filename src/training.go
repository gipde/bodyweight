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
						TExercise{Exercise: TRIZEPSDIP_MIT_UNTERSTUETZUNG,Note:"Beine nach hinten anwinklen und Fußspitzen auf einen Stuhl absetzen um das Hochdrücken zu erleichtern"},
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
							Note:"mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT,
							Note:"mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen"},
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
						TExercise{Exercise: TRIZEPSDIP_MIT_UNTERSTUETZUNG,Note:"Beine nach hinten anwinklen und Fußspitzen auf einen Stuhl absetzen um das Hochdrücken zu erleichtern"},
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
							Note:"mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: SEITLICHER_AUSFALLSCHRITT,
							Note:"mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: GESPRUNGENE_KNIEBEUGE,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						TExercise{Exercise: RUMAENISCHES_KREUZHEBEN_AUF_EINEM_BEIN_IM_WECHSEL,
							Note: "auf einem Kissen"},
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
						TExercise{Exercise: LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN,Note:"mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						TExercise{Exercise: LIEGESTUETZ_MIT_ABSTOSSEN},
						TExercise{Exercise: MILITARY_PRESS_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: STURZFLUG},
						TExercise{Exercise: ENGER_LIEGESTUETZ_MIT_ERHOEHTEN_FUESSEN},
						TExercise{Exercise: TRIZEPSDIP},
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
				},
			},
			TrainingDay{
				Method: HOCHINTENSITAETSSATZ,
				Exercises: [][]TExercise{
					[]TExercise{
						TExercise{Exercise: AUFSTEHEN_AUS_DEM_EINBINIGEN_KNIESTAND, Note: "Areme in Vorhlate"},
						TExercise{Exercise: AUSFALLSCHRITT},
						TExercise{Exercise: GOOD_MORNING},
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
				},
			},
		},
	},
}