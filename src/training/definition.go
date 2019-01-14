package training

type exercise struct {
	Number     int          `json:"number"`
	Name       string       `json:"name"`
	Type       trainingType `json:"type"`
	Page       int          `json:"page"`
	Comment    string       `json:"comment"`
	Difficulty int          `json:"difficulty"`
}

type tExercise struct {
	Exercise exerciseEnum `json:"exercise"`
	Note     string       `json:"note"`
}

type trainingDay struct {
	Method        trainingMethod `json:"method"`
	Exercises     [][]tExercise  `json:"exercies"`
	MinutesNeeded int            `json:"minutesNeeded"`
}

type trainingWeek struct {
	TrainingDays []trainingDay `json:"TrainingDays"`
	Description  string        `json:"description"`
}

//TrainingState contains every State of the training
type TrainingState struct {
	Level trainingLevel `json:"level"`
	Week  int           `json:"week"`
	Day   int           `json:"day"`
	Unit  int           `json:"unit"`
}

func (t trainingMethod) name() string {
	return trainingMethods[t]
}

// gehört zur Übung
type trainingType int

const (
	druecken trainingType = iota
	ziehen
	beineUndGesaess
	core
	ganzKoerper
)

var trainingTypes = []string{
	"Drücken",
	"Ziehen",
	"Beine und Gesäß",
	"Core",
	"Ganzkörperlich",
}

func (t trainingType) name() string {
	return trainingTypes[t]
}

type trainingLevel int

const (
	basisProgram trainingLevel = iota
	firstClass
	masterClass
	chiefClass
)

var trainingLevels = []string{
	"Basisprogramm für Einsteiger",
	"First Class",
	"Master Class",
	"Chief Class",
}

func (t trainingLevel) name() string {
	return trainingLevels[t]
}

type trainingMethod int

const (
	stufenIntervall trainingMethod = iota
	intervallSatz
	superSatz
	zirkelIntervall
	hochIntensitaetsSatz
)

var trainingMethods = []string{
	"Stufenintervalle",
	"Intervallsätze",
	"den Supersatz",
	"Zirkelintervalle",
	"Hochintensitätssätze",
}

type exerciseEnum int

// Varianten werden als eigene Übung eingetragen
const (
	liegestuetzMitErhoehtenHaenden exerciseEnum = iota
	tuerziehen
	Seestern
	umgekehrtesBankdruecken
	ausfallschrittNachHintenImWechsel
	rumaenischesKreuzhebenAufEinemBeinImWechsel
	engeKniebeuge
	schwimmer
	seitlicherAusfallschritt
	schraegerCrunch
	klassischerLiegestuetz
	militaryPressMitErhoehtenHaenden
	engerLiegestuetzMitErhoehtenHaenden
	kniebeugeImAusfallschritt
	curlMitHandtuch
	beinheber
	kaefer
	russischerTwist
	crunch
	liegestuetzMitErhoehtenFuessen
	liegestuetzMitAbstossen
	militaryPress
	baerenGang
	engerLiegestuetz
	gesprungeneKniebeuge
	ausfallschrittNachVorneImWechsel
	klimmzugMitUnterstuetzung
	umgekehrtesBankdrueckenImUntergriff
	tuerziehenImUntergriff
	vUp
	kaeferIpsilateral
	crunchItUp
	haengendesBeinheben
	pogoSprung
	hueftTwist
	knieheberImStehen
	bodyRock
	vierPhasenLiegestuetz
	goodMorning
	kaeferKontralateral
	seitlichesHueftheben
	aufstehenAusDemEinbeinigenKniestand
	ausfallSchritt
	fahrradfahren
	kreuzheben
	kaeferUnilateral
	trizepsdipMitUnterstuetzung
	militaryPressMitErhoehtenFuessen
	sturzflug
	engerLiegestuetzMitErhoehtenFuessen
	trizepsdip
	einbeinigeKniebeugeMitUnterstuetzung
	kistenSprung
	klimmzug
	beinTwist
	tiefeKniebeuge
	ironMike
	seitensprung
	einbeinigerHueftstrecker
	breiterSturzflug
	gestreckterBeinTwist
	einarmigerLiegestuetzMitErhoehtenHaenden
	hueftStrecker
	nackentrainerInBauchlage
	einarmigerLiegestuetz
	federnderLiegestuetz
	erhoehterTrizepsstrecker
	pistole
	knieenderBeinwechsel
	einarmigesTuerziehen
	seitlichesKnieoeffnenImStand
	schraegerVUp
	einbeinigeKniebeuge
	kreuzschritt
	gestrecktesHaengendesBeinheben
	storch
	handstandLiegestuetz
	klappmesser
	einarmigerLiegestuetzMitErhoehtenFuessen
	bergsteiger
	gekreuzterBergsteiger
	pointer
	vorgebeugtesSeitlichesSchulterheben
)

func (n exerciseEnum) Name() string {
	return exes[n].Name
}

var exes = []exercise{
	exercise{
		Name:       "Liegestütz mit erhöhten Händen",
		Type:       druecken,
		Page:       112,
		Difficulty: 1,
	},
	exercise{
		Name:       "Tür ziehen",
		Type:       ziehen,
		Page:       145,
		Difficulty: 1,
	},

	exercise{
		Name:       "Seestern",
		Type:       core,
		Page:       207,
		Difficulty: 1,
	},
	exercise{
		Name:       "Umgekehrtes Bankdrücken",
		Type:       ziehen,
		Page:       147,
		Difficulty: 2,
	},
	exercise{
		Name:       "Ausfallschritt nach hinten im Wechsel",
		Type:       beineUndGesaess,
		Page:       178,
		Difficulty: 1,
	},
	exercise{
		Name:       "Rumänisches Kruezheben auf einem Bein im Wechsel",
		Type:       beineUndGesaess,
		Page:       176,
		Difficulty: 2,
	},
	exercise{
		Name:       "Enge Kniebeuge",
		Type:       beineUndGesaess,
		Page:       183,
		Difficulty: 1,
	},
	exercise{
		Name:       "Schwimmer",
		Type:       beineUndGesaess,
		Page:       174,
		Difficulty: 2,
	},
	exercise{
		Name:       "Seitlicher Ausfallschritt",
		Type:       beineUndGesaess,
		Page:       179,
		Difficulty: 2,
	},
	exercise{
		Name:       "Schräger Crunch",
		Type:       beineUndGesaess,
		Page:       215,
		Difficulty: 1,
	},
	exercise{
		Name:       "Klassischer Liegestütz",
		Type:       druecken,
		Page:       112,
		Difficulty: 2,
	},
	exercise{
		Name:       "Military Press mit erhöhten Händen",
		Type:       druecken,
		Page:       135,
		Difficulty: 2,
	},
	exercise{
		Name:       "Enger Liegestütz mit erhöhten Händen",
		Type:       druecken,
		Page:       130,
		Difficulty: 2,
	},
	exercise{
		Name:       "Kniebeuge im Ausfallschritt",
		Type:       beineUndGesaess,
		Page:       178,
		Difficulty: 2,
	},
	exercise{
		Name:       "Curl mit Handtuch",
		Type:       ziehen,
		Page:       159,
		Difficulty: 2,
	},
	exercise{
		Name:       "Beinheber",
		Type:       core,
		Page:       217,
		Difficulty: 1,
	},
	exercise{
		Name:       "Käfer",
		Type:       core,
		Page:       204,
		Difficulty: 2,
	},
	exercise{
		Name:       "Russischer Twist",
		Type:       core,
		Page:       213,
		Difficulty: 1,
	},
	exercise{
		Name:       "Crunch",
		Type:       core,
		Page:       215,
		Difficulty: 1,
	},
	exercise{
		Name:       "Liegestütz mit erhöhten Füßen",
		Type:       druecken,
		Page:       113,
		Difficulty: 3,
	},
	exercise{
		Name:       "Liegestütz mit Abstoßen",
		Type:       druecken,
		Page:       117,
		Difficulty: 3,
	},
	exercise{
		Name:       "Military Press",
		Type:       druecken,
		Page:       134,
		Difficulty: 3,
	},
	exercise{
		Name:       "Bärengang",
		Type:       druecken,
		Page:       110,
		Difficulty: 1,
	},
	exercise{
		Name:       "Enger Liegestütz",
		Type:       druecken,
		Page:       130,
		Difficulty: 3,
	},
	exercise{
		Name:       "Gesprungene Kniebeuge",
		Type:       beineUndGesaess,
		Page:       191,
		Difficulty: 1,
	},
	exercise{
		Name:       "Ausfallschritt nach vorn im Wechsel",
		Type:       beineUndGesaess,
		Page:       177,
		Difficulty: 1,
	},
	exercise{
		Name:       "Klimmzug mit Unterstützung",
		Type:       ziehen,
		Page:       152,
		Difficulty: 2,
	},
	exercise{
		Name:       "Umgekehrtes Bankdrücken im Untergriff",
		Type:       ziehen,
		Page:       148,
		Difficulty: 3,
	},
	exercise{
		Name:       "Türziehen",
		Type:       ziehen,
		Page:       145,
		Difficulty: 3,
	},
	exercise{
		Name:       "V-Up",
		Type:       core,
		Page:       219,
		Difficulty: 2,
	},
	exercise{
		Name:       "Käfer ipsilateral",
		Type:       core,
		Page:       204,
		Difficulty: 2,
	},
	exercise{
		Name:       "Crunch It Up",
		Type:       core,
		Page:       214,
		Difficulty: 1,
	},
	exercise{
		Name:       "Hängendes Beinheben",
		Type:       core,
		Page:       223,
		Difficulty: 3,
	},
	exercise{
		Name:       "Pogo Sprung",
		Type:       beineUndGesaess,
		Page:       199,
		Difficulty: 2,
	},
	exercise{
		Name:       "Hüfttwist",
		Type:       core,
		Page:       206,
		Difficulty: 1,
	},
	exercise{
		Name:       "Knieheber im Stehen",
		Type:       ganzKoerper,
		Page:       228,
		Difficulty: 1,
	},
	exercise{
		Name:       "Bodyrock",
		Type:       core,
		Page:       208,
		Difficulty: 1,
	},
	exercise{
		Name:       "4 Phasen Liegestütz",
		Type:       ganzKoerper,
		Page:       229,
		Difficulty: 2,
	},
	exercise{
		Name:       "Good Morning",
		Type:       beineUndGesaess,
		Page:       164,
		Difficulty: 2,
	},
	exercise{
		Name:       "Käfer kontralateral",
		Type:       core,
		Page:       204,
		Difficulty: 2,
	},
	exercise{
		Name:       "Seitliches Hüftheben",
		Type:       core,
		Page:       204,
		Difficulty: 3,
	},
	exercise{
		Name:       "Aufstehen aus dem einbeinigen Kniestand",
		Type:       beineUndGesaess,
		Page:       169,
		Difficulty: 1,
	},
	exercise{
		Name:       "Ausfallschritt",
		Type:       beineUndGesaess,
		Page:       177,
		Difficulty: 2,
	},
	exercise{
		Name:       "Fahrradfahren",
		Type:       core,
		Page:       218,
		Difficulty: 2,
	},
	exercise{
		Name:       "Kreuzheben",
		Type:       ziehen,
		Page:       168,
		Difficulty: 1,
	},
	exercise{
		Name:       "Käfer unilateral",
		Type:       core,
		Page:       204,
		Difficulty: 2,
	},
	exercise{
		Name:       "Trizepsdip mit Unterstützung",
		Type:       druecken,
		Page:       132,
		Difficulty: 2,
	},
	exercise{
		Name:       "Military Press mit erhöhten Füßen",
		Type:       druecken,
		Page:       135,
		Difficulty: 3,
	},
	exercise{
		Name:       "Sturzflug",
		Type:       druecken,
		Page:       120,
		Difficulty: 3,
	},
	exercise{
		Name:       "Enger Liegestütz mit erhöhten Füßen",
		Type:       druecken,
		Page:       130,
		Difficulty: 2,
	},
	exercise{
		Name:       "Trizepsdip",
		Type:       druecken,
		Page:       130,
		Difficulty: 4,
	},
	exercise{
		Name:       "Einbeinige Kniebeuge mit Unterstützung im Wechsel",
		Type:       beineUndGesaess,
		Page:       187,
		Difficulty: 4,
	},
	exercise{
		Name:       "Kistensprung",
		Type:       beineUndGesaess,
		Page:       192,
		Difficulty: 1,
	},
	exercise{
		Name:       "Klimmzug",
		Type:       ziehen,
		Page:       150,
		Difficulty: 3,
	},
	exercise{
		Name:       "Beintwist",
		Type:       core,
		Page:       221,
		Difficulty: 2,
	},
	exercise{
		Name:       "Tiefe Kniebeuge",
		Type:       beineUndGesaess,
		Page:       221,
		Difficulty: 3,
	},
	exercise{
		Name:       "Iron Mike",
		Type:       beineUndGesaess,
		Page:       195,
		Difficulty: 3,
	},
	exercise{
		Name:       "Seitsprung",
		Type:       beineUndGesaess,
		Page:       194,
		Difficulty: 2,
	},
	exercise{
		Name:       "Einbeiniger Hüftstrecker",
		Type:       beineUndGesaess,
		Page:       194,
		Difficulty: 2,
	},
	exercise{
		Name:       "Breiter Sturzflug",
		Type:       beineUndGesaess,
		Page:       121,
		Difficulty: 3,
	},
	exercise{
		Name:       "Gestreckter Beintwist",
		Type:       core,
		Page:       221,
		Difficulty: 3,
	},
	exercise{
		Name:       "Einarmiger Liegestütz mit erhöhten Händen im Wechsel",
		Type:       druecken,
		Page:       125,
		Difficulty: 4,
	},
	exercise{
		Name:       "Hüftstrecker",
		Type:       beineUndGesaess,
		Page:       181,
		Difficulty: 2,
	},
	exercise{
		Name:       "Nackentrainer in Bauchlage",
		Type:       core,
		Page:       227,
		Difficulty: 2,
	},
	exercise{
		Name:       "Einarmiger Liegestütz",
		Type:       druecken,
		Page:       124,
		Difficulty: 4,
	},
	exercise{
		Name:       "Federnder Liegestütz",
		Type:       druecken,
		Page:       117,
		Difficulty: 3,
	},
	exercise{
		Name:       "Erhöhter Trizepsstrecker",
		Type:       druecken,
		Page:       131,
		Difficulty: 4,
	},
	exercise{
		Name:       "Pistole",
		Type:       beineUndGesaess,
		Page:       188,
		Difficulty: 4,
	},
	exercise{
		Name:       "Kniender Beinwechsel",
		Type:       beineUndGesaess,
		Page:       175,
		Difficulty: 2,
	},
	exercise{
		Name:       "Einarmiges Türziehen",
		Type:       ziehen,
		Page:       146,
		Difficulty: 2,
	},
	exercise{
		Name:       "Seitliches Knieöffnen im Stand",
		Type:       beineUndGesaess,
		Page:       166,
		Difficulty: 1,
	},
	exercise{
		Name:       "Schräger V-Up",
		Type:       core,
		Page:       220,
		Difficulty: 3,
	},
	exercise{
		Name:       "Einbeinige Kniebeuge im Wechsel",
		Type:       beineUndGesaess,
		Page:       188,
		Difficulty: 4,
	},
	exercise{
		Name:       "Kreuzschritt",
		Type:       beineUndGesaess,
		Page:       167,
		Difficulty: 1,
	},
	exercise{
		Name:       "Gestrecktes hängendes Beinheben",
		Type:       core,
		Page:       223,
		Difficulty: 1,
	},
	exercise{
		Name:       "Storch",
		Type:       beineUndGesaess,
		Page:       165,
		Difficulty: 1,
	},
	exercise{
		Name:       "Handstandliegestütz",
		Type:       beineUndGesaess,
		Page:       138,
		Difficulty: 1,
	},
	exercise{
		Name:       "Klappmesser",
		Type:       core,
		Page:       222,
		Difficulty: 3,
	},
	exercise{
		Name:       "Einarmiger Liegestütz mit erhöhten Füßen",
		Type:       druecken,
		Page:       126,
		Difficulty: 4,
	},
	exercise{
		Name:       "Bergsteiger",
		Type:       core,
		Page:       202,
		Difficulty: 2,
	},
	exercise{
		Name:       "gekreuzter Bergsteiger",
		Type:       core,
		Page:       203,
		Difficulty: 2,
	},
	exercise{
		Name:       "Pointer",
		Type:       beineUndGesaess,
		Page:       174,
		Difficulty: 2,
	},
	exercise{
		Name:       "Vorgebeugtes seitliches Schulterheben",
		Type:       beineUndGesaess,
		Page:       156,
		Difficulty: 1,
	},
}

var trainings = []trainingWeek{
	trainingWeek{
		Description: "Muskuläre Ausdauer",
		TrainingDays: []trainingDay{
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: Seestern},
						tExercise{Exercise: umgekehrtesBankdruecken},
					},
					[]tExercise{
						tExercise{Exercise: klassischerLiegestuetz},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: tuerziehen},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: klimmzugMitUnterstuetzung},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: umgekehrtesBankdruecken},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: klimmzugMitUnterstuetzung},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "Füße erhöht"},
					},
				},
			},
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: engeKniebeuge, Note: "Arme in Vorhalte"},
						tExercise{Exercise: schwimmer},
					},
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: gesprungeneKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						tExercise{Exercise: ausfallschrittNachHintenImWechsel, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: hueftStrecker},
						tExercise{Exercise: kaeferUnilateral, Note: "Beine gebeugt"},
					},
					[]tExercise{
						tExercise{Exercise: pistole},
						tExercise{Exercise: ironMike},
						tExercise{Exercise: einbeinigerHueftstrecker},
						tExercise{Exercise: nackentrainerInBauchlage},
					},
				},
			},
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: Seestern},
						tExercise{Exercise: umgekehrtesBankdruecken},
					},
					[]tExercise{
						tExercise{Exercise: klassischerLiegestuetz},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: tuerziehen},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: tuerziehen},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: einarmigesTuerziehen},
					},
				},
			},
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: seitlicherAusfallschritt},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: engeKniebeuge},
						tExercise{Exercise: schraegerCrunch},
					},
					[]tExercise{
						tExercise{Exercise: seitlicherAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: engeKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: russischerTwist},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						tExercise{Exercise: seitlicherAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: beinTwist},
					},
					[]tExercise{
						tExercise{Exercise: knieenderBeinwechsel},
						tExercise{Exercise: kistenSprung},
						tExercise{Exercise: storch},
						tExercise{Exercise: beinTwist},
					},
				},
			},
		},
	},
	trainingWeek{
		Description: "Muskuläre Ausdauer",
		TrainingDays: []trainingDay{
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: Seestern},
						tExercise{Exercise: umgekehrtesBankdruecken},
					},
					[]tExercise{
						tExercise{Exercise: klassischerLiegestuetz},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: tuerziehen},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: klimmzugMitUnterstuetzung},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: umgekehrtesBankdruecken},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: klimmzugMitUnterstuetzung},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "Füße erhöht"},
					},
				},
			},
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: engeKniebeuge, Note: "Arme in Vorhalte"},
						tExercise{Exercise: schwimmer},
					},
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: gesprungeneKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						tExercise{Exercise: ausfallschrittNachHintenImWechsel, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: hueftStrecker},
						tExercise{Exercise: kaeferUnilateral, Note: "Beine gebeugt"},
					},
					[]tExercise{
						tExercise{Exercise: pistole},
						tExercise{Exercise: ironMike},
						tExercise{Exercise: einbeinigerHueftstrecker},
						tExercise{Exercise: nackentrainerInBauchlage},
					},
				},
			},
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: Seestern},
						tExercise{Exercise: umgekehrtesBankdruecken},
					},
					[]tExercise{
						tExercise{Exercise: klassischerLiegestuetz},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: tuerziehen},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: tuerziehen},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: einarmigesTuerziehen},
					},
				},
			},
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: seitlicherAusfallschritt},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: engeKniebeuge},
						tExercise{Exercise: schraegerCrunch},
					},
					[]tExercise{
						tExercise{Exercise: seitlicherAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: engeKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: russischerTwist},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						tExercise{Exercise: seitlicherAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: beinTwist},
					},
					[]tExercise{
						tExercise{Exercise: knieenderBeinwechsel},
						tExercise{Exercise: kistenSprung},
						tExercise{Exercise: storch},
						tExercise{Exercise: beinTwist},
					},
				},
			},
		},
	},
	trainingWeek{
		Description: "Kraft",
		TrainingDays: []trainingDay{
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: klassischerLiegestuetz},
						tExercise{Exercise: militaryPressMitErhoehtenHaenden},
						tExercise{Exercise: engerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: Seestern},
					},
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: engerLiegestuetz},
						tExercise{Exercise: trizepsdipMitUnterstuetzung, Note: "Beine nach hinten anwinklen und Fußspitzen auf einen Stuhl absetzen um das Hochdrücken zu erleichtern"},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: trizepsdipMitUnterstuetzung},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetz},
						tExercise{Exercise: handstandLiegestuetz},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: erhoehterTrizepsstrecker, Note: "Hände etwa hüfthoch"},
					},
				},
			},
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: kniebeugeImAusfallschritt},
						tExercise{Exercise: seitlicherAusfallschritt},
						tExercise{Exercise: engeKniebeuge,
							Note: "Arme in T-Halte mit 1-3 Sek. Haltezeit am tiefsten Punkt",
						},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
					},
					[]tExercise{
						tExercise{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: gesprungeneKniebeuge,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						tExercise{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						tExercise{Exercise: einbeinigerHueftstrecker,
							Note: "auf einem Kissen"},
					},
					[]tExercise{
						tExercise{Exercise: knieenderBeinwechsel},
						tExercise{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 1-3 Sek. Pause",
						},
						tExercise{Exercise: ironMike},
					},
				},
			},
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehen, Note: "Im Untergriff"},
						tExercise{Exercise: curlMitHandtuch},
					},
					[]tExercise{
						tExercise{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehen, Note: "Im Untergriff"},
						tExercise{Exercise: curlMitHandtuch},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: tuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug, Note: "mit 1-3 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: einarmigesTuerziehen},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						tExercise{Exercise: tuerziehen, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: beinheber},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
						tExercise{Exercise: russischerTwist},
						tExercise{Exercise: crunch},
					},
					[]tExercise{
						tExercise{Exercise: beinheber, Note: "mit gekreuzten Armen auf der Brust"},
						tExercise{Exercise: kaeferIpsilateral, Note: "iBeine gestreckt"},
						tExercise{Exercise: fahrradfahren},
						tExercise{Exercise: kaeferKontralateral},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben, Note: "Beine gestreckt und parallel zum Boden"},
						tExercise{Exercise: nackentrainerInBauchlage},
						tExercise{Exercise: vUp},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben, Note: "Beine gestreckt bis ganz nach oben"},
						tExercise{Exercise: nackentrainerInBauchlage},
						tExercise{Exercise: klappmesser},
						tExercise{Exercise: kaeferKontralateral, Note: "Beine gestreckt"},
					},
				},
			},
		},
	},
	trainingWeek{
		Description: "Kraft",
		TrainingDays: []trainingDay{
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: klassischerLiegestuetz},
						tExercise{Exercise: militaryPressMitErhoehtenHaenden},
						tExercise{Exercise: engerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: Seestern},
					},
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: engerLiegestuetz},
						tExercise{Exercise: trizepsdipMitUnterstuetzung, Note: "Beine nach hinten anwinklen und Fußspitzen auf einen Stuhl absetzen um das Hochdrücken zu erleichtern"},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: trizepsdipMitUnterstuetzung},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetz},
						tExercise{Exercise: handstandLiegestuetz},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: erhoehterTrizepsstrecker, Note: "Hände etwa hüfthoch"},
					},
				},
			},
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: kniebeugeImAusfallschritt},
						tExercise{Exercise: seitlicherAusfallschritt},
						tExercise{Exercise: engeKniebeuge,
							Note: "Arme in T-Halte mit 1-3 Sek. Haltezeit am tiefsten Punkt",
						},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
					},
					[]tExercise{
						tExercise{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: gesprungeneKniebeuge,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						tExercise{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						tExercise{Exercise: einbeinigerHueftstrecker,
							Note: "auf einem Kissen"},
					},
					[]tExercise{
						tExercise{Exercise: knieenderBeinwechsel},
						tExercise{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 1-3 Sek. Pause",
						},
						tExercise{Exercise: ironMike},
					},
				},
			},
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehen, Note: "Im Untergriff"},
						tExercise{Exercise: curlMitHandtuch},
					},
					[]tExercise{
						tExercise{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehen, Note: "Im Untergriff"},
						tExercise{Exercise: curlMitHandtuch},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: tuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug, Note: "mit 1-3 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: einarmigesTuerziehen},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						tExercise{Exercise: tuerziehen, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: beinheber},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
						tExercise{Exercise: russischerTwist},
						tExercise{Exercise: crunch},
					},
					[]tExercise{
						tExercise{Exercise: beinheber, Note: "mit gekreuzten Armen auf der Brust"},
						tExercise{Exercise: kaeferIpsilateral, Note: "iBeine gestreckt"},
						tExercise{Exercise: fahrradfahren},
						tExercise{Exercise: kaeferKontralateral},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben, Note: "Beine gestreckt und parallel zum Boden"},
						tExercise{Exercise: nackentrainerInBauchlage},
						tExercise{Exercise: vUp},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben, Note: "Beine gestreckt bis ganz nach oben"},
						tExercise{Exercise: nackentrainerInBauchlage},
						tExercise{Exercise: klappmesser},
						tExercise{Exercise: kaeferKontralateral, Note: "Beine gestreckt"},
					},
				},
			},
		},
	},
	trainingWeek{
		Description: "Powerblock",
		TrainingDays: []trainingDay{
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: liegestuetzMitAbstossen},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: baerenGang},
						tExercise{Exercise: engerLiegestuetz},
						tExercise{Exercise: Seestern},
					},
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenFuessen, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: liegestuetzMitAbstossen},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: engerLiegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: trizepsdip},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetz},
						tExercise{Exercise: federnderLiegestuetz},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: erhoehterTrizepsstrecker},
						tExercise{Exercise: liegestuetzMitAbstossen},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: federnderLiegestuetz},
						tExercise{Exercise: handstandLiegestuetz, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: erhoehterTrizepsstrecker, Note: "Hände etwa hüfthoch"},
						tExercise{Exercise: liegestuetzMitAbstossen},
					},
				},
			},
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: gesprungeneKniebeuge},
						tExercise{Exercise: ausfallschrittNachVorneImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
						tExercise{Exercise: engeKniebeuge,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						tExercise{Exercise: gesprungeneKniebeuge,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: kistenSprung},
					},
					[]tExercise{
						tExercise{Exercise: pistole},
						tExercise{Exercise: kistenSprung},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: ironMike},
						tExercise{Exercise: knieenderBeinwechsel},
						tExercise{Exercise: seitensprung},
					},
					[]tExercise{
						tExercise{Exercise: pistole, Note: "Mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: kistenSprung},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: ironMike},
						tExercise{Exercise: knieenderBeinwechsel},
						tExercise{Exercise: seitensprung},
					},
				},
			},
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: tuerziehen,
							Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: curlMitHandtuch},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
						tExercise{Exercise: tuerziehenImUntergriff,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: tuerziehen,
							Note: "mit 4-6 Sek. Haltezeit bei der Kontraktionam höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						tExercise{Exercise: tuerziehenImUntergriff},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: einarmigesTuerziehen},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
						tExercise{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: curlMitHandtuch},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "Füße erhöht"},
						tExercise{Exercise: einarmigesTuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der kontraktion am Höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						tExercise{Exercise: klimmzug, Note: "Bis zum Brustbein hochziehen"},
						tExercise{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: vUp},
						tExercise{Exercise: russischerTwist},
						tExercise{Exercise: kaeferIpsilateral,
							Note: "Beine gestreckt"},
						tExercise{Exercise: crunchItUp},
						tExercise{Exercise: haengendesBeinheben},
						tExercise{Exercise: beinheber},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben},
						tExercise{Exercise: beinTwist},
						tExercise{Exercise: hueftTwist},
						tExercise{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
						tExercise{Exercise: vUp},
						tExercise{Exercise: russischerTwist},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben},
						tExercise{Exercise: fahrradfahren},
						tExercise{Exercise: nackentrainerInBauchlage},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
						tExercise{Exercise: vUp},
						tExercise{Exercise: beinTwist},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt, Beine gestreckt bis ganz nach oben"},
						tExercise{Exercise: fahrradfahren, Note: "langsam, 2 Sek. für das Heranziehen jedes Knies"},
						tExercise{Exercise: nackentrainerInBauchlage, Note: "mit 4-6 Sek. Pause"},
						tExercise{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
						tExercise{Exercise: klappmesser},
						tExercise{Exercise: beinTwist},
					},
				},
			},
		},
	},
	trainingWeek{
		Description: "Powerblock",
		TrainingDays: []trainingDay{
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: liegestuetzMitAbstossen},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: baerenGang},
						tExercise{Exercise: engerLiegestuetz},
						tExercise{Exercise: Seestern},
					},
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenFuessen, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: liegestuetzMitAbstossen},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: engerLiegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: trizepsdip},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetz},
						tExercise{Exercise: federnderLiegestuetz},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: erhoehterTrizepsstrecker},
						tExercise{Exercise: liegestuetzMitAbstossen},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: federnderLiegestuetz},
						tExercise{Exercise: handstandLiegestuetz, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: erhoehterTrizepsstrecker, Note: "Hände etwa hüfthoch"},
						tExercise{Exercise: liegestuetzMitAbstossen},
					},
				},
			},
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: gesprungeneKniebeuge},
						tExercise{Exercise: ausfallschrittNachVorneImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
						tExercise{Exercise: engeKniebeuge,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						tExercise{Exercise: gesprungeneKniebeuge,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: kistenSprung},
					},
					[]tExercise{
						tExercise{Exercise: pistole},
						tExercise{Exercise: kistenSprung},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: ironMike},
						tExercise{Exercise: knieenderBeinwechsel},
						tExercise{Exercise: seitensprung},
					},
					[]tExercise{
						tExercise{Exercise: pistole, Note: "Mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: kistenSprung},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: ironMike},
						tExercise{Exercise: knieenderBeinwechsel},
						tExercise{Exercise: seitensprung},
					},
				},
			},
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: tuerziehen,
							Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: curlMitHandtuch},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
						tExercise{Exercise: tuerziehenImUntergriff,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: tuerziehen,
							Note: "mit 4-6 Sek. Haltezeit bei der Kontraktionam höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						tExercise{Exercise: tuerziehenImUntergriff},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: einarmigesTuerziehen},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
						tExercise{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: curlMitHandtuch},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "Füße erhöht"},
						tExercise{Exercise: einarmigesTuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der kontraktion am Höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						tExercise{Exercise: klimmzug, Note: "Bis zum Brustbein hochziehen"},
						tExercise{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: vUp},
						tExercise{Exercise: russischerTwist},
						tExercise{Exercise: kaeferIpsilateral,
							Note: "Beine gestreckt"},
						tExercise{Exercise: crunchItUp},
						tExercise{Exercise: haengendesBeinheben},
						tExercise{Exercise: beinheber},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben},
						tExercise{Exercise: beinTwist},
						tExercise{Exercise: hueftTwist},
						tExercise{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
						tExercise{Exercise: vUp},
						tExercise{Exercise: russischerTwist},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben},
						tExercise{Exercise: fahrradfahren},
						tExercise{Exercise: nackentrainerInBauchlage},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
						tExercise{Exercise: vUp},
						tExercise{Exercise: beinTwist},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt, Beine gestreckt bis ganz nach oben"},
						tExercise{Exercise: fahrradfahren, Note: "langsam, 2 Sek. für das Heranziehen jedes Knies"},
						tExercise{Exercise: nackentrainerInBauchlage, Note: "mit 4-6 Sek. Pause"},
						tExercise{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
						tExercise{Exercise: klappmesser},
						tExercise{Exercise: beinTwist},
					},
				},
			},
		},
	},
	trainingWeek{
		Description: "Wechselblock",
		TrainingDays: []trainingDay{
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: militaryPressMitErhoehtenHaenden},
						tExercise{Exercise: liegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: engerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: Seestern},
					},
					[]tExercise{
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: klassischerLiegestuetz},
						tExercise{Exercise: engerLiegestuetz},
						tExercise{Exercise: Seestern},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: trizepsdip},
						tExercise{Exercise: Seestern},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetz},
						tExercise{Exercise: handstandLiegestuetz},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: trizepsdip, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
					},
				},
			},
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: gesprungeneKniebeuge},
						tExercise{Exercise: ausfallschrittNachVorneImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen und mit 1-3 Sek. Haltezeit in der Mitte"},
						tExercise{Exercise: pogoSprung},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						tExercise{Exercise: gesprungeneKniebeuge,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen und mit 1-3 Sek. Haltezeit in der Mitte"},
						tExercise{Exercise: kistenSprung},
					},
					[]tExercise{
						tExercise{Exercise: pistole},
						tExercise{Exercise: seitlichesKnieoeffnenImStand},
						tExercise{Exercise: kistenSprung},
						tExercise{Exercise: ironMike},
						tExercise{Exercise: knieenderBeinwechsel},
						tExercise{Exercise: seitensprung},
					},
					[]tExercise{
						tExercise{Exercise: pistole, Note: "im Wechsel mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: kistenSprung},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: ironMike},
						tExercise{Exercise: knieenderBeinwechsel},
						tExercise{Exercise: seitensprung},
					},
				},
			},
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehenImUntergriff},
						tExercise{Exercise: curlMitHandtuch},
					},
					[]tExercise{
						tExercise{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: curlMitHandtuch},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: tuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am Höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug, Note: "mit 1-3 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: einarmigesTuerziehen},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						tExercise{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			trainingDay{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: russischerTwist},
						tExercise{Exercise: hueftTwist},
						tExercise{Exercise: knieheberImStehen},
					},
					[]tExercise{
						tExercise{Exercise: russischerTwist},
						tExercise{Exercise: bodyRock},
						tExercise{Exercise: tiefeKniebeuge, Note: "Hände hinter dem Kopf"},
					},
					[]tExercise{
						tExercise{Exercise: schraegerVUp, Note: "4 Sätze pro Seite"},
						tExercise{Exercise: bodyRock},
						tExercise{Exercise: tiefeKniebeuge, Note: "Arme in Streamline Position"},
					},
					[]tExercise{
						tExercise{Exercise: vUp},
						tExercise{Exercise: schraegerVUp, Note: "nach jedem Satz Seitenwechsel, insgesamt 4 Sätze"},
						tExercise{Exercise: russischerTwist},
						tExercise{Exercise: bergsteiger},
					},
				},
			},
			trainingDay{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "10 Wiederholungen pro Seite"},
						tExercise{Exercise: tuerziehen,
							Note: "8 Wiederholungen"},
						tExercise{Exercise: klassischerLiegestuetz,
							Note: "6 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: umgekehrtesBankdruecken,
							Note: "6 Wiederholungen"},
						tExercise{Exercise: seitlicherAusfallschritt,
							Note: "12 Wiederholungen"},
						tExercise{Exercise: klassischerLiegestuetz,
							Note: "8 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeuge,
							Note: "12 Wiederholungen"},
						tExercise{Exercise: sturzflug,
							Note: "6 Wiederholungen"},
						tExercise{Exercise: umgekehrtesBankdruecken,
							Note: "8 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: pistole,
							Note: "12 Wiederholungen"},
						tExercise{Exercise: handstandLiegestuetz,
							Note: "6 Wiederholungen"},
						tExercise{Exercise: klimmzug,
							Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
	trainingWeek{
		Description: "Wechselblock",
		TrainingDays: []trainingDay{
			trainingDay{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenHaenden, Note: "brusthoch abgestützt"},
						tExercise{Exercise: bodyRock},
						tExercise{Exercise: vierPhasenLiegestuetz},
					},
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: Seestern},
						tExercise{Exercise: tiefeKniebeuge},
					},
					[]tExercise{
						tExercise{Exercise: klassischerLiegestuetz},
						tExercise{Exercise: liegestuetzMitAbstossen},
						tExercise{Exercise: tiefeKniebeuge, Note: "Arme in T-Halte"},
					},
					[]tExercise{
						tExercise{Exercise: klassischerLiegestuetz},
						tExercise{Exercise: liegestuetzMitAbstossen},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
					},
				},
			},
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						tExercise{Exercise: engeKniebeuge, Note: "Arme in Vorhalte mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: goodMorning, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
					},
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel},
						tExercise{Exercise: seitlicherAusfallschritt},
						tExercise{Exercise: gesprungeneKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeuge},
						tExercise{Exercise: kreuzschritt},
						tExercise{Exercise: seitlicherAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: hueftStrecker},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeuge},
						tExercise{Exercise: seitlichesKnieoeffnenImStand},
						tExercise{Exercise: ironMike},
						tExercise{Exercise: einbeinigerHueftstrecker},
					},
				},
			},
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: klimmzugMitUnterstuetzung},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: tuerziehen, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "mit gebeugten Knien"},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Beine gestreckt"},
						tExercise{Exercise: tuerziehenImUntergriff},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: tuerziehen, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
						tExercise{Exercise: tuerziehenImUntergriff},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: einarmigesTuerziehen},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						tExercise{Exercise: tuerziehenImUntergriff},
						tExercise{Exercise: curlMitHandtuch},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "Füße erhöht"},
						tExercise{Exercise: einarmigesTuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						tExercise{Exercise: klimmzug, Note: "bis zum Brustbein hochziehen"},
						tExercise{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: beinheber},
						tExercise{Exercise: kaeferKontralateral, Note: "Beine gebeugt"},
						tExercise{Exercise: russischerTwist},
						tExercise{Exercise: seitlichesHueftheben},
					},
					[]tExercise{
						tExercise{Exercise: beinheber, Note: "mit gekreuzten Armen auf der Brust"},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
						tExercise{Exercise: fahrradfahren},
						tExercise{Exercise: kaeferKontralateral, Note: "Beine gestreckt"},
					},
					[]tExercise{
						tExercise{Exercise: gestrecktesHaengendesBeinheben},
						tExercise{Exercise: nackentrainerInBauchlage},
						tExercise{Exercise: vUp},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben, Note: "Beine gestreckt, Füße bis zu den Händen"},
						tExercise{Exercise: nackentrainerInBauchlage, Note: "mit 1-3 Sek. Pause"},
						tExercise{Exercise: gekreuzterBergsteiger},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
					},
				},
			},
			trainingDay{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel, Note: "10 Wiederholungen pro Seite"},
						tExercise{Exercise: tuerziehen, Note: "8 Wiederholungen"},
						tExercise{Exercise: klassischerLiegestuetz, Note: "6 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "6 Wiederholungen"},
						tExercise{Exercise: seitlicherAusfallschritt, Note: "12 Wiederholungen"},
						tExercise{Exercise: klassischerLiegestuetz, Note: "8 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeuge, Note: "12 Wiederholungen pro Seite"},
						tExercise{Exercise: sturzflug, Note: "6 Wiederholungen"},
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "8 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: pistole, Note: "6 Wiederholungen pro Seite"},
						tExercise{Exercise: handstandLiegestuetz, Note: "6 Wiederholungen"},
						tExercise{Exercise: klimmzug, Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
	trainingWeek{
		Description: "Wechselblock",
		TrainingDays: []trainingDay{
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: klassischerLiegestuetz},
						tExercise{Exercise: militaryPressMitErhoehtenHaenden},
						tExercise{Exercise: engerLiegestuetzMitErhoehtenHaenden},
					},
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: militaryPressMitErhoehtenHaenden},
						tExercise{Exercise: engerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: Seestern},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: trizepsdip},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetz},
						tExercise{Exercise: handstandLiegestuetz},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: erhoehterTrizepsstrecker, Note: "Hände etwa hüfthoch"},
					},
				},
			},
			trainingDay{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: aufstehenAusDemEinbeinigenKniestand, Note: "Arme in Vorhalte"},
						tExercise{Exercise: ausfallSchritt},
						tExercise{Exercise: goodMorning},
					},
					[]tExercise{
						tExercise{Exercise: ironMike},
						tExercise{Exercise: seitensprung},
						tExercise{Exercise: tiefeKniebeuge, Note: "Arme in T-Halte"},
					},
					[]tExercise{
						tExercise{Exercise: ironMike},
						tExercise{Exercise: seitensprung},
						tExercise{Exercise: tiefeKniebeuge, Note: "Hände an den Hüften"},
					},
					[]tExercise{
						tExercise{Exercise: ironMike},
						tExercise{Exercise: seitensprung},
						tExercise{Exercise: pointer},
						tExercise{Exercise: engeKniebeuge, Note: "Hände hinter dem Kopf"},
					},
				},
			},
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
						tExercise{Exercise: tuerziehenImUntergriff},
					},
					[]tExercise{
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
						tExercise{Exercise: tuerziehenImUntergriff},
					},
					[]tExercise{
						tExercise{Exercise: klimmzugMitUnterstuetzung},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
						tExercise{Exercise: tuerziehen},
					},
					[]tExercise{
						tExercise{Exercise: klimmzug},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: umgekehrtesBankdrueckenImUntergriff},
						tExercise{Exercise: einarmigesTuerziehen},
					},
				},
			},
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: vUp},
						tExercise{Exercise: russischerTwist},
						tExercise{Exercise: kaeferKontralateral, Note: "Beine gestreckt"},
						tExercise{Exercise: seitlichesHueftheben},
						tExercise{Exercise: fahrradfahren},
						tExercise{Exercise: beinheber},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben},
						tExercise{Exercise: beinTwist},
						tExercise{Exercise: einbeinigerHueftstrecker},
						tExercise{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
						tExercise{Exercise: vUp},
						tExercise{Exercise: russischerTwist},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben},
						tExercise{Exercise: fahrradfahren},
						tExercise{Exercise: nackentrainerInBauchlage, Note: "mit 4-6 Sek. Pause"},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
						tExercise{Exercise: vUp},
						tExercise{Exercise: beinTwist},
					},
					[]tExercise{
						tExercise{Exercise: haengendesBeinheben, Note: "Beine gestreckt, Füße bis zu den Händen mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						tExercise{Exercise: fahrradfahren, Note: "12 langsame Wdh. pro Seite)"},
						tExercise{Exercise: nackentrainerInBauchlage, Note: "mit 4-6 Sek. Pause"},
						tExercise{Exercise: kaeferKontralateral, Note: "Beine gebeugt"},
						tExercise{Exercise: klappmesser},
						tExercise{Exercise: beinTwist, Note: "6 langsame Wiederholungen pro Seite"},
					},
				},
			},
			trainingDay{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel, Note: "10 Wiederholungen pro Seite"},
						tExercise{Exercise: tuerziehen, Note: "8 Wiederholungen"},
						tExercise{Exercise: klassischerLiegestuetz, Note: "6 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "6 Wiederholungen pro Seite"},
						tExercise{Exercise: seitlicherAusfallschritt, Note: "12 Wiederholungen"},
						tExercise{Exercise: klassischerLiegestuetz, Note: "8 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung, Note: "12 Wiederholungen pro Seite"},
						tExercise{Exercise: sturzflug, Note: "6 Wiederholungen"},
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "8 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: pistole, Note: "6 Wiederholungen pro Seite"},
						tExercise{Exercise: handstandLiegestuetz, Note: "6 Wiederholungen"},
						tExercise{Exercise: klimmzug, Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
	trainingWeek{
		Description: "Wechselblock",
		TrainingDays: []trainingDay{
			trainingDay{
				Method: superSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: liegestuetzMitAbstossen},
						tExercise{Exercise: militaryPress},
						tExercise{Exercise: baerenGang},
						tExercise{Exercise: engerLiegestuetz},
						tExercise{Exercise: Seestern},
					},
					[]tExercise{
						tExercise{Exercise: liegestuetzMitErhoehtenFuessen, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: liegestuetzMitAbstossen},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: breiterSturzflug},
						tExercise{Exercise: engerLiegestuetzMitErhoehtenFuessen},
						tExercise{Exercise: Seestern},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetz},
						tExercise{Exercise: federnderLiegestuetz},
						tExercise{Exercise: militaryPressMitErhoehtenFuessen},
						tExercise{Exercise: sturzflug},
						tExercise{Exercise: erhoehterTrizepsstrecker},
						tExercise{Exercise: liegestuetzMitAbstossen},
					},
					[]tExercise{
						tExercise{Exercise: einarmigerLiegestuetz},
						tExercise{Exercise: federnderLiegestuetz},
						tExercise{Exercise: handstandLiegestuetz},
						tExercise{Exercise: sturzflug, Note: "mit Haltezeit wenn die Brust am tiefsten Punkt der Kontraktion zwischen den Händen ist"},
						tExercise{Exercise: erhoehterTrizepsstrecker, Note: "Trizepsstrecker, hände kniehoch"},
						tExercise{Exercise: liegestuetzMitAbstossen},
					},
				},
			},
			trainingDay{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: kniebeugeImAusfallschritt},
						tExercise{Exercise: seitlicherAusfallschritt},
						tExercise{Exercise: engeKniebeuge, Note: "Arme in T-Halte, mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel, Note: "auf einem Kissen"},
					},
					[]tExercise{
						tExercise{Exercise: kniebeugeImAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: gesprungeneKniebeuge, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel, Note: "auf einem Kissen"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeuge},
						tExercise{Exercise: kniebeugeImAusfallschritt, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: seitlicherAusfallschritt, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						tExercise{Exercise: einbeinigerHueftstrecker},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeuge, Note: "knieend"},
						tExercise{Exercise: kniebeugeImAusfallschritt, Note: "dabei am tiefsten Punkt einen Rucksack über den Kopf stemen"},
						tExercise{Exercise: ironMike},
						tExercise{Exercise: knieenderBeinwechsel},
					},
				},
			},
			trainingDay{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: tuerziehen, Note: "Füße sind hinter den Händen platziert, gehen Sie dafür einem Schritt zurürck"},
						tExercise{Exercise: kreuzheben},
						tExercise{Exercise: curlMitHandtuch},
					},
					[]tExercise{
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: tiefeKniebeuge, Note: "Arme in Streamline Position"},
					},
					[]tExercise{
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: tiefeKniebeuge, Note: "Hände hinter dem Kopf"},
					},
					[]tExercise{
						tExercise{Exercise: klimmzugMitUnterstuetzung},
						tExercise{Exercise: umgekehrtesBankdruecken},
						tExercise{Exercise: tuerziehen},
						tExercise{Exercise: vorgebeugtesSeitlichesSchulterheben},
					},
				},
			},
			trainingDay{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: crunchItUp},
						tExercise{Exercise: kaeferIpsilateral},
						tExercise{Exercise: beinheber},
						tExercise{Exercise: kaeferKontralateral},
					},
					[]tExercise{
						tExercise{Exercise: fahrradfahren},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
						tExercise{Exercise: bodyRock, Note: "mit vorwärtsgreifen"},
						tExercise{Exercise: gestreckterBeinTwist},
					},
					[]tExercise{
						tExercise{Exercise: schraegerVUp, Note: "gebeugten Beinen"},
						tExercise{Exercise: kreuzheben},
						tExercise{Exercise: beinheber, Note: "mit gekreuzten Armen auf der Brust"},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
					},
					[]tExercise{
						tExercise{Exercise: klappmesser},
						tExercise{Exercise: nackentrainerInBauchlage},
						tExercise{Exercise: schraegerVUp, Note: "mit gekreuzten Armen auf der Brust"},
						tExercise{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
					},
				},
			},
			trainingDay{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					[]tExercise{
						tExercise{Exercise: ausfallschrittNachHintenImWechsel, Note: "10 Wiederholungen pro Seite"},
						tExercise{Exercise: tuerziehen, Note: "8 Wiederholungen"},
						tExercise{Exercise: klassischerLiegestuetz, Note: "6 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "6 Wiederholungen pro Seite"},
						tExercise{Exercise: seitlicherAusfallschritt, Note: "12 Wiederholungen"},
						tExercise{Exercise: klassischerLiegestuetz, Note: "8 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: einbeinigeKniebeugeMitUnterstuetzung, Note: "12 Wiederholungen pro Seite"},
						tExercise{Exercise: sturzflug, Note: "6 Wiederholungen"},
						tExercise{Exercise: umgekehrtesBankdruecken, Note: "8 Wiederholungen"},
					},
					[]tExercise{
						tExercise{Exercise: pistole, Note: "6 Wiederholungen pro Seite"},
						tExercise{Exercise: handstandLiegestuetz, Note: "6 Wiederholungen"},
						tExercise{Exercise: klimmzug, Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
}
