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

//State contains every State of the training
type State struct {
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
	"Druckübung",
	"Zugübung",
	"Übung für Beine und Gesäß",
	"Übung für die Rumpfmuskulator",
	"ganzkörperliche Übung",
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

// func getLevel(level int) trainingLevel {
// 	return trainingLevel(level)
// }

// func (t trainingLevel) int() int {
// 	return int(t)
// }

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

func (n exerciseEnum) get() exercise {
	return exesList[n]
}

var exesList = []exercise{
	{
		Name:       "Liegestütz mit erhöhten Händen",
		Type:       druecken,
		Page:       112,
		Difficulty: 1,
	},
	{
		Name:       "Tür ziehen",
		Type:       ziehen,
		Page:       145,
		Difficulty: 1,
	},

	{
		Name:       "Seestern",
		Type:       core,
		Page:       207,
		Difficulty: 1,
	},
	{
		Name:       "Umgekehrtes Bankdrücken",
		Type:       ziehen,
		Page:       147,
		Difficulty: 2,
	},
	{
		Name:       "Ausfallschritt nach hinten im Wechsel",
		Type:       beineUndGesaess,
		Page:       178,
		Difficulty: 1,
	},
	{
		Name:       "Rumänisches Kreuzheben auf einem Bein im Wechsel",
		Type:       beineUndGesaess,
		Page:       176,
		Difficulty: 2,
	},
	{
		Name:       "Enge Kniebeuge",
		Type:       beineUndGesaess,
		Page:       183,
		Difficulty: 1,
	},
	{
		Name:       "Schwimmer",
		Type:       beineUndGesaess,
		Page:       174,
		Difficulty: 2,
	},
	{
		Name:       "Seitlicher Ausfallschritt",
		Type:       beineUndGesaess,
		Page:       179,
		Difficulty: 2,
	},
	{
		Name:       "Schräger Crunch",
		Type:       beineUndGesaess,
		Page:       215,
		Difficulty: 1,
	},
	{
		Name:       "Klassischer Liegestütz",
		Type:       druecken,
		Page:       112,
		Difficulty: 2,
	},
	{
		Name:       "Military Press mit erhöhten Händen",
		Type:       druecken,
		Page:       135,
		Difficulty: 2,
	},
	{
		Name:       "Enger Liegestütz mit erhöhten Händen",
		Type:       druecken,
		Page:       130,
		Difficulty: 2,
	},
	{
		Name:       "Kniebeuge im Ausfallschritt",
		Type:       beineUndGesaess,
		Page:       178,
		Difficulty: 2,
	},
	{
		Name:       "Curl mit Handtuch",
		Type:       ziehen,
		Page:       159,
		Difficulty: 2,
	},
	{
		Name:       "Beinheber",
		Type:       core,
		Page:       217,
		Difficulty: 1,
	},
	{
		Name:       "Käfer",
		Type:       core,
		Page:       204,
		Difficulty: 2,
	},
	{
		Name:       "Russischer Twist",
		Type:       core,
		Page:       213,
		Difficulty: 1,
	},
	{
		Name:       "Crunch",
		Type:       core,
		Page:       215,
		Difficulty: 1,
	},
	{
		Name:       "Liegestütz mit erhöhten Füßen",
		Type:       druecken,
		Page:       113,
		Difficulty: 3,
	},
	{
		Name:       "Liegestütz mit Abstoßen",
		Type:       druecken,
		Page:       117,
		Difficulty: 3,
	},
	{
		Name:       "Military Press",
		Type:       druecken,
		Page:       134,
		Difficulty: 3,
	},
	{
		Name:       "Bärengang",
		Type:       druecken,
		Page:       110,
		Difficulty: 1,
	},
	{
		Name:       "Enger Liegestütz",
		Type:       druecken,
		Page:       130,
		Difficulty: 3,
	},
	{
		Name:       "Gesprungene Kniebeuge",
		Type:       beineUndGesaess,
		Page:       191,
		Difficulty: 1,
	},
	{
		Name:       "Ausfallschritt nach vorn im Wechsel",
		Type:       beineUndGesaess,
		Page:       177,
		Difficulty: 1,
	},
	{
		Name:       "Klimmzug mit Unterstützung",
		Type:       ziehen,
		Page:       152,
		Difficulty: 2,
	},
	{
		Name:       "Umgekehrtes Bankdrücken im Untergriff",
		Type:       ziehen,
		Page:       148,
		Difficulty: 3,
	},
	{
		Name:       "Türziehen",
		Type:       ziehen,
		Page:       145,
		Difficulty: 3,
	},
	{
		Name:       "V-Up",
		Type:       core,
		Page:       219,
		Difficulty: 2,
	},
	{
		Name:       "Käfer ipsilateral",
		Type:       core,
		Page:       204,
		Difficulty: 2,
	},
	{
		Name:       "Crunch It Up",
		Type:       core,
		Page:       214,
		Difficulty: 1,
	},
	{
		Name:       "Hängendes Beinheben",
		Type:       core,
		Page:       223,
		Difficulty: 3,
	},
	{
		Name:       "Pogo Sprung",
		Type:       beineUndGesaess,
		Page:       199,
		Difficulty: 2,
	},
	{
		Name:       "Hüfttwist",
		Type:       core,
		Page:       206,
		Difficulty: 1,
	},
	{
		Name:       "Knieheber im Stehen",
		Type:       ganzKoerper,
		Page:       228,
		Difficulty: 1,
	},
	{
		Name:       "Bodyrock",
		Type:       core,
		Page:       208,
		Difficulty: 1,
	},
	{
		Name:       "4 Phasen Liegestütz",
		Type:       ganzKoerper,
		Page:       229,
		Difficulty: 2,
	},
	{
		Name:       "Good Morning",
		Type:       beineUndGesaess,
		Page:       164,
		Difficulty: 2,
	},
	{
		Name:       "Käfer kontralateral",
		Type:       core,
		Page:       204,
		Difficulty: 2,
	},
	{
		Name:       "Seitliches Hüftheben",
		Type:       core,
		Page:       204,
		Difficulty: 3,
	},
	{
		Name:       "Aufstehen aus dem einbeinigen Kniestand",
		Type:       beineUndGesaess,
		Page:       169,
		Difficulty: 1,
	},
	{
		Name:       "Ausfallschritt",
		Type:       beineUndGesaess,
		Page:       177,
		Difficulty: 2,
	},
	{
		Name:       "Fahrradfahren",
		Type:       core,
		Page:       218,
		Difficulty: 2,
	},
	{
		Name:       "Kreuzheben",
		Type:       ziehen,
		Page:       168,
		Difficulty: 1,
	},
	{
		Name:       "Käfer unilateral",
		Type:       core,
		Page:       204,
		Difficulty: 2,
	},
	{
		Name:       "Trizepsdip mit Unterstützung",
		Type:       druecken,
		Page:       132,
		Difficulty: 2,
	},
	{
		Name:       "Military Press mit erhöhten Füßen",
		Type:       druecken,
		Page:       135,
		Difficulty: 3,
	},
	{
		Name:       "Sturzflug",
		Type:       druecken,
		Page:       120,
		Difficulty: 3,
	},
	{
		Name:       "Enger Liegestütz mit erhöhten Füßen",
		Type:       druecken,
		Page:       130,
		Difficulty: 2,
	},
	{
		Name:       "Trizepsdip",
		Type:       druecken,
		Page:       130,
		Difficulty: 4,
	},
	{
		Name:       "Einbeinige Kniebeuge mit Unterstützung im Wechsel",
		Type:       beineUndGesaess,
		Page:       187,
		Difficulty: 4,
	},
	{
		Name:       "Kistensprung",
		Type:       beineUndGesaess,
		Page:       192,
		Difficulty: 1,
	},
	{
		Name:       "Klimmzug",
		Type:       ziehen,
		Page:       150,
		Difficulty: 3,
	},
	{
		Name:       "Beintwist",
		Type:       core,
		Page:       221,
		Difficulty: 2,
	},
	{
		Name:       "Tiefe Kniebeuge",
		Type:       beineUndGesaess,
		Page:       221,
		Difficulty: 3,
	},
	{
		Name:       "Iron Mike",
		Type:       beineUndGesaess,
		Page:       195,
		Difficulty: 3,
	},
	{
		Name:       "Seitsprung",
		Type:       beineUndGesaess,
		Page:       194,
		Difficulty: 2,
	},
	{
		Name:       "Einbeiniger Hüftstrecker",
		Type:       beineUndGesaess,
		Page:       194,
		Difficulty: 2,
	},
	{
		Name:       "Breiter Sturzflug",
		Type:       beineUndGesaess,
		Page:       121,
		Difficulty: 3,
	},
	{
		Name:       "Gestreckter Beintwist",
		Type:       core,
		Page:       221,
		Difficulty: 3,
	},
	{
		Name:       "Einarmiger Liegestütz mit erhöhten Händen im Wechsel",
		Type:       druecken,
		Page:       125,
		Difficulty: 4,
	},
	{
		Name:       "Hüftstrecker",
		Type:       beineUndGesaess,
		Page:       181,
		Difficulty: 2,
	},
	{
		Name:       "Nackentrainer in Bauchlage",
		Type:       core,
		Page:       227,
		Difficulty: 2,
	},
	{
		Name:       "Einarmiger Liegestütz",
		Type:       druecken,
		Page:       124,
		Difficulty: 4,
	},
	{
		Name:       "Federnder Liegestütz",
		Type:       druecken,
		Page:       117,
		Difficulty: 3,
	},
	{
		Name:       "Erhöhter Trizepsstrecker",
		Type:       druecken,
		Page:       131,
		Difficulty: 4,
	},
	{
		Name:       "Pistole",
		Type:       beineUndGesaess,
		Page:       188,
		Difficulty: 4,
	},
	{
		Name:       "Kniender Beinwechsel",
		Type:       beineUndGesaess,
		Page:       175,
		Difficulty: 2,
	},
	{
		Name:       "Einarmiges Türziehen",
		Type:       ziehen,
		Page:       146,
		Difficulty: 2,
	},
	{
		Name:       "Seitliches Knieöffnen im Stand",
		Type:       beineUndGesaess,
		Page:       166,
		Difficulty: 1,
	},
	{
		Name:       "Schräger V-Up",
		Type:       core,
		Page:       220,
		Difficulty: 3,
	},
	{
		Name:       "Einbeinige Kniebeuge im Wechsel",
		Type:       beineUndGesaess,
		Page:       188,
		Difficulty: 4,
	},
	{
		Name:       "Kreuzschritt",
		Type:       beineUndGesaess,
		Page:       167,
		Difficulty: 1,
	},
	{
		Name:       "Gestrecktes hängendes Beinheben",
		Type:       core,
		Page:       223,
		Difficulty: 1,
	},
	{
		Name:       "Storch",
		Type:       beineUndGesaess,
		Page:       165,
		Difficulty: 1,
	},
	{
		Name:       "Handstandliegestütz",
		Type:       beineUndGesaess,
		Page:       138,
		Difficulty: 1,
	},
	{
		Name:       "Klappmesser",
		Type:       core,
		Page:       222,
		Difficulty: 3,
	},
	{
		Name:       "Einarmiger Liegestütz mit erhöhten Füßen",
		Type:       druecken,
		Page:       126,
		Difficulty: 4,
	},
	{
		Name:       "Bergsteiger",
		Type:       core,
		Page:       202,
		Difficulty: 2,
	},
	{
		Name:       "gekreuzter Bergsteiger",
		Type:       core,
		Page:       203,
		Difficulty: 2,
	},
	{
		Name:       "Pointer",
		Type:       beineUndGesaess,
		Page:       174,
		Difficulty: 2,
	},
	{
		Name:       "Vorgebeugtes seitliches Schulterheben",
		Type:       beineUndGesaess,
		Page:       156,
		Difficulty: 1,
	},
}

var trainings = []trainingWeek{
	{
		Description: "Muskuläre Ausdauer",
		TrainingDays: []trainingDay{
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetzMitErhoehtenHaenden},
						{Exercise: tuerziehen},
						{Exercise: Seestern},
						{Exercise: umgekehrtesBankdruecken},
					},
					{
						{Exercise: klassischerLiegestuetz},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: militaryPress},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: klimmzugMitUnterstuetzung},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: umgekehrtesBankdruecken},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: klimmzugMitUnterstuetzung},
						{Exercise: sturzflug},
						{Exercise: umgekehrtesBankdruecken, Note: "Die Füße sind erhöht. "},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschrittNachHintenImWechsel},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: engeKniebeuge, Note: "Arme in Vorhalte"},
						{Exercise: schwimmer},
					},
					{
						{Exercise: ausfallschrittNachHintenImWechsel},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: gesprungeneKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						{Exercise: ausfallschrittNachHintenImWechsel, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: hueftStrecker},
						{Exercise: kaeferUnilateral, Note: "Beine gebeugt"},
					},
					{
						{Exercise: pistole},
						{Exercise: ironMike},
						{Exercise: einbeinigerHueftstrecker},
						{Exercise: nackentrainerInBauchlage},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetzMitErhoehtenHaenden},
						{Exercise: tuerziehen},
						{Exercise: Seestern},
						{Exercise: umgekehrtesBankdruecken},
					},
					{
						{Exercise: klassischerLiegestuetz},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: militaryPress},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						{Exercise: militaryPress},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: klimmzug},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: einarmigesTuerziehen},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: seitlicherAusfallschritt},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: engeKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: schraegerCrunch},
					},
					{
						{Exercise: seitlicherAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: engeKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: russischerTwist},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						{Exercise: seitlicherAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: beinTwist},
					},
					{
						{Exercise: knieenderBeinwechsel},
						{Exercise: kistenSprung},
						{Exercise: storch},
						{Exercise: beinTwist},
					},
				},
			},
		},
	},
	{
		Description: "Muskuläre Ausdauer",
		TrainingDays: []trainingDay{
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetzMitErhoehtenHaenden},
						{Exercise: tuerziehen},
						{Exercise: Seestern},
						{Exercise: umgekehrtesBankdruecken},
					},
					{
						{Exercise: klassischerLiegestuetz},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: militaryPress},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: klimmzugMitUnterstuetzung},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: umgekehrtesBankdruecken},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: klimmzugMitUnterstuetzung},
						{Exercise: sturzflug},
						{Exercise: umgekehrtesBankdruecken, Note: "Füße erhöht"},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschrittNachHintenImWechsel},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: engeKniebeuge, Note: "Arme in Vorhalte"},
						{Exercise: schwimmer},
					},
					{
						{Exercise: ausfallschrittNachHintenImWechsel},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: gesprungeneKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						{Exercise: ausfallschrittNachHintenImWechsel, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: hueftStrecker},
						{Exercise: kaeferUnilateral, Note: "Beine gebeugt"},
					},
					{
						{Exercise: pistole},
						{Exercise: ironMike},
						{Exercise: einbeinigerHueftstrecker},
						{Exercise: nackentrainerInBauchlage},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetzMitErhoehtenHaenden},
						{Exercise: tuerziehen},
						{Exercise: Seestern},
						{Exercise: umgekehrtesBankdruecken},
					},
					{
						{Exercise: klassischerLiegestuetz},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: militaryPress},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						{Exercise: militaryPress},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: klimmzug},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: einarmigesTuerziehen},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: seitlicherAusfallschritt},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: engeKniebeuge},
						{Exercise: schraegerCrunch},
					},
					{
						{Exercise: seitlicherAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: engeKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: russischerTwist},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						{Exercise: seitlicherAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: beinTwist},
					},
					{
						{Exercise: knieenderBeinwechsel},
						{Exercise: kistenSprung},
						{Exercise: storch},
						{Exercise: beinTwist},
					},
				},
			},
		},
	},
	{
		Description: "Kraft",
		TrainingDays: []trainingDay{
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: klassischerLiegestuetz},
						{Exercise: militaryPressMitErhoehtenHaenden},
						{Exercise: engerLiegestuetzMitErhoehtenHaenden},
						{Exercise: Seestern},
					},
					{
						{Exercise: liegestuetzMitErhoehtenFuessen},
						{Exercise: militaryPress},
						{Exercise: engerLiegestuetz},
						{Exercise: trizepsdipMitUnterstuetzung, Note: "Beine nach hinten anwinklen und Fußspitzen auf einen Stuhl absetzen um das Hochdrücken zu erleichtern"},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: sturzflug},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: trizepsdipMitUnterstuetzung},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: handstandLiegestuetz},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker, Note: "Hände etwa hüfthoch"},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: kniebeugeImAusfallschritt},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: engeKniebeuge,
							Note: "Arme in T-Halte mit 1-3 Sek. Haltezeit am tiefsten Punkt",
						},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
					},
					{
						{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: gesprungeneKniebeuge,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						{Exercise: einbeinigerHueftstrecker,
							Note: "auf einem Kissen"},
					},
					{
						{Exercise: knieenderBeinwechsel},
						{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 1-3 Sek. Pause",
						},
						{Exercise: ironMike},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: tuerziehen},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen, Note: "Im Untergriff"},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen, Note: "Im Untergriff"},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
					},
					{
						{Exercise: klimmzug, Note: "mit 1-3 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: einarmigesTuerziehen},
						{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						{Exercise: tuerziehen, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: beinheber},
						{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
						{Exercise: russischerTwist},
						{Exercise: crunch},
					},
					{
						{Exercise: beinheber, Note: "mit gekreuzten Armen auf der Brust"},
						{Exercise: kaeferIpsilateral, Note: "iBeine gestreckt"},
						{Exercise: fahrradfahren},
						{Exercise: kaeferKontralateral},
					},
					{
						{Exercise: haengendesBeinheben, Note: "Beine gestreckt und parallel zum Boden"},
						{Exercise: nackentrainerInBauchlage},
						{Exercise: vUp},
						{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
					},
					{
						{Exercise: haengendesBeinheben, Note: "Beine gestreckt bis ganz nach oben"},
						{Exercise: nackentrainerInBauchlage},
						{Exercise: klappmesser},
						{Exercise: kaeferKontralateral, Note: "Beine gestreckt"},
					},
				},
			},
		},
	},
	{
		Description: "Kraft",
		TrainingDays: []trainingDay{
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: klassischerLiegestuetz},
						{Exercise: militaryPressMitErhoehtenHaenden},
						{Exercise: engerLiegestuetzMitErhoehtenHaenden},
						{Exercise: Seestern},
					},
					{
						{Exercise: liegestuetzMitErhoehtenFuessen},
						{Exercise: militaryPress},
						{Exercise: engerLiegestuetz},
						{Exercise: trizepsdipMitUnterstuetzung, Note: "Beine nach hinten anwinklen und Fußspitzen auf einen Stuhl absetzen um das Hochdrücken zu erleichtern"},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: sturzflug},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: trizepsdipMitUnterstuetzung},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: handstandLiegestuetz},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker, Note: "Hände etwa hüfthoch"},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: kniebeugeImAusfallschritt},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: engeKniebeuge,
							Note: "Arme in T-Halte mit 1-3 Sek. Haltezeit am tiefsten Punkt",
						},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
					},
					{
						{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: gesprungeneKniebeuge,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
						},
						{Exercise: einbeinigerHueftstrecker,
							Note: "auf einem Kissen"},
					},
					{
						{Exercise: knieenderBeinwechsel},
						{Exercise: kniebeugeImAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 1-3 Sek. Pause",
						},
						{Exercise: ironMike},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: tuerziehen},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen, Note: "Im Untergriff"},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen, Note: "Im Untergriff"},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
					},
					{
						{Exercise: klimmzug, Note: "mit 1-3 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: einarmigesTuerziehen},
						{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						{Exercise: tuerziehen, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: beinheber},
						{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
						{Exercise: russischerTwist},
						{Exercise: crunch},
					},
					{
						{Exercise: beinheber, Note: "mit gekreuzten Armen auf der Brust"},
						{Exercise: kaeferIpsilateral, Note: "iBeine gestreckt"},
						{Exercise: fahrradfahren},
						{Exercise: kaeferKontralateral},
					},
					{
						{Exercise: haengendesBeinheben, Note: "Beine gestreckt und parallel zum Boden"},
						{Exercise: nackentrainerInBauchlage},
						{Exercise: vUp},
						{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
					},
					{
						{Exercise: haengendesBeinheben, Note: "Beine gestreckt bis ganz nach oben"},
						{Exercise: nackentrainerInBauchlage},
						{Exercise: klappmesser},
						{Exercise: kaeferKontralateral, Note: "Beine gestreckt"},
					},
				},
			},
		},
	},
	{
		Description: "Powerblock",
		TrainingDays: []trainingDay{
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetzMitErhoehtenFuessen},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPress},
						{Exercise: baerenGang},
						{Exercise: engerLiegestuetz},
						{Exercise: Seestern},
					},
					{
						{Exercise: liegestuetzMitErhoehtenFuessen, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: sturzflug},
						{Exercise: engerLiegestuetzMitErhoehtenFuessen},
						{Exercise: trizepsdip},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: federnderLiegestuetz},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker},
						{Exercise: liegestuetzMitAbstossen},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenFuessen},
						{Exercise: federnderLiegestuetz},
						{Exercise: handstandLiegestuetz, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker, Note: "Hände etwa hüfthoch"},
						{Exercise: liegestuetzMitAbstossen},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: gesprungeneKniebeuge},
						{Exercise: ausfallschrittNachVorneImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
						{Exercise: engeKniebeuge,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						{Exercise: gesprungeneKniebeuge,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: kistenSprung},
					},
					{
						{Exercise: pistole},
						{Exercise: kistenSprung},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
						{Exercise: seitensprung},
					},
					{
						{Exercise: pistole, Note: "Mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: kistenSprung},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
						{Exercise: seitensprung},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen,
							Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: curlMitHandtuch},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
						{Exercise: tuerziehenImUntergriff,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen,
							Note: "mit 4-6 Sek. Haltezeit bei der Kontraktionam höchsten Punkt"},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						{Exercise: tuerziehenImUntergriff},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen},
						{Exercise: einarmigesTuerziehen},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
						{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: umgekehrtesBankdruecken, Note: "Füße erhöht"},
						{Exercise: einarmigesTuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der kontraktion am Höchsten Punkt"},
						{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						{Exercise: klimmzug, Note: "Bis zum Brustbein hochziehen"},
						{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: vUp},
						{Exercise: russischerTwist},
						{Exercise: kaeferIpsilateral,
							Note: "Beine gestreckt"},
						{Exercise: crunchItUp},
						{Exercise: haengendesBeinheben},
						{Exercise: beinheber},
					},
					{
						{Exercise: haengendesBeinheben},
						{Exercise: beinTwist},
						{Exercise: hueftTwist},
						{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
						{Exercise: vUp},
						{Exercise: russischerTwist},
					},
					{
						{Exercise: haengendesBeinheben},
						{Exercise: fahrradfahren},
						{Exercise: nackentrainerInBauchlage},
						{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
						{Exercise: vUp},
						{Exercise: beinTwist},
					},
					{
						{Exercise: haengendesBeinheben, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt, Beine gestreckt bis ganz nach oben"},
						{Exercise: fahrradfahren, Note: "langsam, 2 Sek. für das Heranziehen jedes Knies"},
						{Exercise: nackentrainerInBauchlage, Note: "mit 4-6 Sek. Pause"},
						{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
						{Exercise: klappmesser},
						{Exercise: beinTwist},
					},
				},
			},
		},
	},
	{
		Description: "Powerblock",
		TrainingDays: []trainingDay{
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetzMitErhoehtenFuessen},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPress},
						{Exercise: baerenGang},
						{Exercise: engerLiegestuetz},
						{Exercise: Seestern},
					},
					{
						{Exercise: liegestuetzMitErhoehtenFuessen, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: sturzflug},
						{Exercise: engerLiegestuetzMitErhoehtenFuessen},
						{Exercise: trizepsdip},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: federnderLiegestuetz},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker},
						{Exercise: liegestuetzMitAbstossen},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenFuessen},
						{Exercise: federnderLiegestuetz},
						{Exercise: handstandLiegestuetz, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker, Note: "Hände etwa hüfthoch"},
						{Exercise: liegestuetzMitAbstossen},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: gesprungeneKniebeuge},
						{Exercise: ausfallschrittNachVorneImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen"},
						{Exercise: engeKniebeuge,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						{Exercise: gesprungeneKniebeuge,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: kistenSprung},
					},
					{
						{Exercise: pistole},
						{Exercise: kistenSprung},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
						{Exercise: seitensprung},
					},
					{
						{Exercise: pistole, Note: "Mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: kistenSprung},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
						{Exercise: seitensprung},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen,
							Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: curlMitHandtuch},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
						{Exercise: tuerziehenImUntergriff,
							Note: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt"},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen,
							Note: "mit 4-6 Sek. Haltezeit bei der Kontraktionam höchsten Punkt"},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						{Exercise: tuerziehenImUntergriff},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen},
						{Exercise: einarmigesTuerziehen},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
						{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: umgekehrtesBankdruecken, Note: "Füße erhöht"},
						{Exercise: einarmigesTuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der kontraktion am Höchsten Punkt"},
						{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						{Exercise: klimmzug, Note: "Bis zum Brustbein hochziehen"},
						{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: vUp},
						{Exercise: russischerTwist},
						{Exercise: kaeferIpsilateral,
							Note: "Beine gestreckt"},
						{Exercise: crunchItUp},
						{Exercise: haengendesBeinheben},
						{Exercise: beinheber},
					},
					{
						{Exercise: haengendesBeinheben},
						{Exercise: beinTwist},
						{Exercise: hueftTwist},
						{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
						{Exercise: vUp},
						{Exercise: russischerTwist},
					},
					{
						{Exercise: haengendesBeinheben},
						{Exercise: fahrradfahren},
						{Exercise: nackentrainerInBauchlage},
						{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
						{Exercise: vUp},
						{Exercise: beinTwist},
					},
					{
						{Exercise: haengendesBeinheben, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt, Beine gestreckt bis ganz nach oben"},
						{Exercise: fahrradfahren, Note: "langsam, 2 Sek. für das Heranziehen jedes Knies"},
						{Exercise: nackentrainerInBauchlage, Note: "mit 4-6 Sek. Pause"},
						{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
						{Exercise: klappmesser},
						{Exercise: beinTwist},
					},
				},
			},
		},
	},
	{
		Description: "Wechselblock",
		TrainingDays: []trainingDay{
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: militaryPressMitErhoehtenHaenden},
						{Exercise: liegestuetzMitErhoehtenHaenden},
						{Exercise: engerLiegestuetzMitErhoehtenHaenden},
						{Exercise: Seestern},
					},
					{
						{Exercise: militaryPress},
						{Exercise: klassischerLiegestuetz},
						{Exercise: engerLiegestuetz},
						{Exercise: Seestern},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: sturzflug},
						{Exercise: trizepsdip},
						{Exercise: Seestern},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: handstandLiegestuetz},
						{Exercise: sturzflug},
						{Exercise: trizepsdip, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: gesprungeneKniebeuge},
						{Exercise: ausfallschrittNachVorneImWechsel,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen und mit 1-3 Sek. Haltezeit in der Mitte"},
						{Exercise: pogoSprung},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung},
						{Exercise: gesprungeneKniebeuge,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt,
							Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel,
							Note: "auf einem Kissen und mit 1-3 Sek. Haltezeit in der Mitte"},
						{Exercise: kistenSprung},
					},
					{
						{Exercise: pistole},
						{Exercise: seitlichesKnieoeffnenImStand},
						{Exercise: kistenSprung},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
						{Exercise: seitensprung},
					},
					{
						{Exercise: pistole, Note: "im Wechsel mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: kistenSprung},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
						{Exercise: seitensprung},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: tuerziehen},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehenImUntergriff},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzugMitUnterstuetzung,
							Note: "ohne Unterstützung in der Negativphase"},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am Höchsten Punkt"},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
					},
					{
						{Exercise: klimmzug, Note: "mit 1-3 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: einarmigesTuerziehen},
						{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: russischerTwist},
						{Exercise: hueftTwist},
						{Exercise: knieheberImStehen},
					},
					{
						{Exercise: russischerTwist},
						{Exercise: bodyRock},
						{Exercise: tiefeKniebeuge, Note: "Hände hinter dem Kopf"},
					},
					{
						{Exercise: schraegerVUp, Note: "4 Sätze pro Seite"},
						{Exercise: bodyRock},
						{Exercise: tiefeKniebeuge, Note: "Arme in Streamline Position"},
					},
					{
						{Exercise: vUp},
						{Exercise: schraegerVUp, Note: "nach jedem Satz Seitenwechsel, insgesamt 4 Sätze"},
						{Exercise: russischerTwist},
						{Exercise: bergsteiger},
					},
				},
			},
			{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschrittNachHintenImWechsel,
							Note: "10 Wiederholungen pro Seite"},
						{Exercise: tuerziehen,
							Note: "8 Wiederholungen"},
						{Exercise: klassischerLiegestuetz,
							Note: "6 Wiederholungen"},
					},
					{
						{Exercise: umgekehrtesBankdruecken,
							Note: "6 Wiederholungen"},
						{Exercise: seitlicherAusfallschritt,
							Note: "12 Wiederholungen"},
						{Exercise: klassischerLiegestuetz,
							Note: "8 Wiederholungen"},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: "12 Wiederholungen"},
						{Exercise: sturzflug,
							Note: "6 Wiederholungen"},
						{Exercise: umgekehrtesBankdruecken,
							Note: "8 Wiederholungen"},
					},
					{
						{Exercise: pistole,
							Note: "12 Wiederholungen"},
						{Exercise: handstandLiegestuetz,
							Note: "6 Wiederholungen"},
						{Exercise: klimmzug,
							Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
	{
		Description: "Wechselblock",
		TrainingDays: []trainingDay{
			{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetzMitErhoehtenHaenden, Note: "brusthoch abgestützt"},
						{Exercise: bodyRock},
						{Exercise: vierPhasenLiegestuetz},
					},
					{
						{Exercise: liegestuetzMitErhoehtenHaenden},
						{Exercise: Seestern},
						{Exercise: tiefeKniebeuge},
					},
					{
						{Exercise: klassischerLiegestuetz},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: tiefeKniebeuge, Note: "Arme in T-Halte"},
					},
					{
						{Exercise: klassischerLiegestuetz},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPressMitErhoehtenFuessen},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschrittNachHintenImWechsel},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
						{Exercise: engeKniebeuge, Note: "Arme in Vorhalte mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: goodMorning, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
					},
					{
						{Exercise: ausfallschrittNachHintenImWechsel},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: gesprungeneKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel},
					},
					{
						{Exercise: einbeinigeKniebeuge},
						{Exercise: kreuzschritt},
						{Exercise: seitlicherAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: hueftStrecker},
					},
					{
						{Exercise: einbeinigeKniebeuge},
						{Exercise: seitlichesKnieoeffnenImStand},
						{Exercise: ironMike},
						{Exercise: einbeinigerHueftstrecker},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: klimmzugMitUnterstuetzung},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: umgekehrtesBankdruecken, Note: "mit gebeugten Knien"},
						{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Beine gestreckt"},
						{Exercise: tuerziehenImUntergriff},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
						{Exercise: tuerziehenImUntergriff},
					},
					{
						{Exercise: klimmzug},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: einarmigesTuerziehen},
						{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						{Exercise: tuerziehenImUntergriff},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug, Note: "mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: umgekehrtesBankdruecken, Note: "Füße erhöht"},
						{Exercise: einarmigesTuerziehen, Note: "mit 1-3 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
						{Exercise: umgekehrtesBankdrueckenImUntergriff, Note: "Füße erhöht"},
						{Exercise: klimmzug, Note: "bis zum Brustbein hochziehen"},
						{Exercise: tuerziehenImUntergriff, Note: "mit 4-6 Sek. Haltezeit bei der Kontraktion am höchsten Punkt"},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: beinheber},
						{Exercise: kaeferKontralateral, Note: "Beine gebeugt"},
						{Exercise: russischerTwist},
						{Exercise: seitlichesHueftheben},
					},
					{
						{Exercise: beinheber, Note: "mit gekreuzten Armen auf der Brust"},
						{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
						{Exercise: fahrradfahren},
						{Exercise: kaeferKontralateral, Note: "Beine gestreckt"},
					},
					{
						{Exercise: gestrecktesHaengendesBeinheben},
						{Exercise: nackentrainerInBauchlage},
						{Exercise: vUp},
						{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
					},
					{
						{Exercise: haengendesBeinheben, Note: "Beine gestreckt, Füße bis zu den Händen"},
						{Exercise: nackentrainerInBauchlage, Note: "mit 1-3 Sek. Pause"},
						{Exercise: gekreuzterBergsteiger},
						{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
					},
				},
			},
			{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschrittNachHintenImWechsel, Note: "10 Wiederholungen pro Seite"},
						{Exercise: tuerziehen, Note: "8 Wiederholungen"},
						{Exercise: klassischerLiegestuetz, Note: "6 Wiederholungen"},
					},
					{
						{Exercise: umgekehrtesBankdruecken, Note: "6 Wiederholungen"},
						{Exercise: seitlicherAusfallschritt, Note: "12 Wiederholungen"},
						{Exercise: klassischerLiegestuetz, Note: "8 Wiederholungen"},
					},
					{
						{Exercise: einbeinigeKniebeuge, Note: "12 Wiederholungen pro Seite"},
						{Exercise: sturzflug, Note: "6 Wiederholungen"},
						{Exercise: umgekehrtesBankdruecken, Note: "8 Wiederholungen"},
					},
					{
						{Exercise: pistole, Note: "6 Wiederholungen pro Seite"},
						{Exercise: handstandLiegestuetz, Note: "6 Wiederholungen"},
						{Exercise: klimmzug, Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
	{
		Description: "Wechselblock",
		TrainingDays: []trainingDay{
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: klassischerLiegestuetz},
						{Exercise: militaryPressMitErhoehtenHaenden},
						{Exercise: engerLiegestuetzMitErhoehtenHaenden},
					},
					{
						{Exercise: liegestuetzMitErhoehtenFuessen},
						{Exercise: militaryPressMitErhoehtenHaenden},
						{Exercise: engerLiegestuetzMitErhoehtenHaenden},
						{Exercise: Seestern},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: sturzflug},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: trizepsdip},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: handstandLiegestuetz},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker, Note: "Hände etwa hüfthoch"},
					},
				},
			},
			{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: aufstehenAusDemEinbeinigenKniestand, Note: "Arme in Vorhalte"},
						{Exercise: ausfallSchritt},
						{Exercise: goodMorning},
					},
					{
						{Exercise: ironMike},
						{Exercise: seitensprung},
						{Exercise: tiefeKniebeuge, Note: "Arme in T-Halte"},
					},
					{
						{Exercise: ironMike},
						{Exercise: seitensprung},
						{Exercise: tiefeKniebeuge, Note: "Hände an den Hüften"},
					},
					{
						{Exercise: ironMike},
						{Exercise: seitensprung},
						{Exercise: pointer},
						{Exercise: engeKniebeuge, Note: "Hände hinter dem Kopf"},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
						{Exercise: tuerziehenImUntergriff},
					},
					{
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
						{Exercise: tuerziehenImUntergriff},
					},
					{
						{Exercise: klimmzugMitUnterstuetzung},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: klimmzug},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdrueckenImUntergriff},
						{Exercise: einarmigesTuerziehen},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: vUp},
						{Exercise: russischerTwist},
						{Exercise: kaeferKontralateral, Note: "Beine gestreckt"},
						{Exercise: seitlichesHueftheben},
						{Exercise: fahrradfahren},
						{Exercise: beinheber},
					},
					{
						{Exercise: haengendesBeinheben},
						{Exercise: beinTwist},
						{Exercise: einbeinigerHueftstrecker},
						{Exercise: kaeferUnilateral, Note: "Beine gestreckt"},
						{Exercise: vUp},
						{Exercise: russischerTwist},
					},
					{
						{Exercise: haengendesBeinheben},
						{Exercise: fahrradfahren},
						{Exercise: nackentrainerInBauchlage, Note: "mit 4-6 Sek. Pause"},
						{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
						{Exercise: vUp},
						{Exercise: beinTwist},
					},
					{
						{Exercise: haengendesBeinheben, Note: "Beine gestreckt, Füße bis zu den Händen mit 4-6 Sek. Haltezeit am höchsten Punkt"},
						{Exercise: fahrradfahren, Note: "12 langsame Wdh. pro Seite)"},
						{Exercise: nackentrainerInBauchlage, Note: "mit 4-6 Sek. Pause"},
						{Exercise: kaeferKontralateral, Note: "Beine gebeugt"},
						{Exercise: klappmesser},
						{Exercise: beinTwist, Note: "6 langsame Wiederholungen pro Seite"},
					},
				},
			},
			{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschrittNachHintenImWechsel, Note: "10 Wiederholungen pro Seite"},
						{Exercise: tuerziehen, Note: "8 Wiederholungen"},
						{Exercise: klassischerLiegestuetz, Note: "6 Wiederholungen"},
					},
					{
						{Exercise: umgekehrtesBankdruecken, Note: "6 Wiederholungen pro Seite"},
						{Exercise: seitlicherAusfallschritt, Note: "12 Wiederholungen"},
						{Exercise: klassischerLiegestuetz, Note: "8 Wiederholungen"},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung, Note: "12 Wiederholungen pro Seite"},
						{Exercise: sturzflug, Note: "6 Wiederholungen"},
						{Exercise: umgekehrtesBankdruecken, Note: "8 Wiederholungen"},
					},
					{
						{Exercise: pistole, Note: "6 Wiederholungen pro Seite"},
						{Exercise: handstandLiegestuetz, Note: "6 Wiederholungen"},
						{Exercise: klimmzug, Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
	{
		Description: "Wechselblock",
		TrainingDays: []trainingDay{
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetzMitErhoehtenFuessen},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPress},
						{Exercise: baerenGang},
						{Exercise: engerLiegestuetz},
						{Exercise: Seestern},
					},
					{
						{Exercise: liegestuetzMitErhoehtenFuessen, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: breiterSturzflug},
						{Exercise: engerLiegestuetzMitErhoehtenFuessen},
						{Exercise: Seestern},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: federnderLiegestuetz},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker},
						{Exercise: liegestuetzMitAbstossen},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: federnderLiegestuetz},
						{Exercise: handstandLiegestuetz},
						{Exercise: sturzflug, Note: "mit Haltezeit wenn die Brust am tiefsten Punkt der Kontraktion zwischen den Händen ist"},
						{Exercise: erhoehterTrizepsstrecker, Note: "Trizepsstrecker, hände kniehoch"},
						{Exercise: liegestuetzMitAbstossen},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: kniebeugeImAusfallschritt},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: engeKniebeuge, Note: "Arme in T-Halte, mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel, Note: "auf einem Kissen"},
					},
					{
						{Exercise: kniebeugeImAusfallschritt, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: gesprungeneKniebeuge, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: rumaenischesKreuzhebenAufEinemBeinImWechsel, Note: "auf einem Kissen"},
					},
					{
						{Exercise: einbeinigeKniebeuge},
						{Exercise: kniebeugeImAusfallschritt, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: seitlicherAusfallschritt, Note: "mit 4-6 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: einbeinigerHueftstrecker},
					},
					{
						{Exercise: einbeinigeKniebeuge, Note: "knieend"},
						{Exercise: kniebeugeImAusfallschritt, Note: "dabei am tiefsten Punkt einen Rucksack über den Kopf stemen"},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
					},
				},
			},
			{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: tuerziehen, Note: "Füße sind hinter den Händen platziert, gehen Sie dafür einem Schritt zurürck"},
						{Exercise: kreuzheben},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: tiefeKniebeuge, Note: "Arme in Streamline Position"},
					},
					{
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: tiefeKniebeuge, Note: "Hände hinter dem Kopf"},
					},
					{
						{Exercise: klimmzugMitUnterstuetzung},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: vorgebeugtesSeitlichesSchulterheben},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: crunchItUp},
						{Exercise: kaeferIpsilateral},
						{Exercise: beinheber},
						{Exercise: kaeferKontralateral},
					},
					{
						{Exercise: fahrradfahren},
						{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
						{Exercise: bodyRock, Note: "mit vorwärtsgreifen"},
						{Exercise: gestreckterBeinTwist},
					},
					{
						{Exercise: schraegerVUp, Note: "gebeugten Beinen"},
						{Exercise: kreuzheben},
						{Exercise: beinheber, Note: "mit gekreuzten Armen auf der Brust"},
						{Exercise: kaeferIpsilateral, Note: "Beine gebeugt"},
					},
					{
						{Exercise: klappmesser},
						{Exercise: nackentrainerInBauchlage},
						{Exercise: schraegerVUp, Note: "mit gekreuzten Armen auf der Brust"},
						{Exercise: kaeferIpsilateral, Note: "Beine gestreckt"},
					},
				},
			},
			{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschrittNachHintenImWechsel, Note: "10 Wiederholungen pro Seite"},
						{Exercise: tuerziehen, Note: "8 Wiederholungen"},
						{Exercise: klassischerLiegestuetz, Note: "6 Wiederholungen"},
					},
					{
						{Exercise: umgekehrtesBankdruecken, Note: "6 Wiederholungen pro Seite"},
						{Exercise: seitlicherAusfallschritt, Note: "12 Wiederholungen"},
						{Exercise: klassischerLiegestuetz, Note: "8 Wiederholungen"},
					},
					{
						{Exercise: einbeinigeKniebeugeMitUnterstuetzung, Note: "12 Wiederholungen pro Seite"},
						{Exercise: sturzflug, Note: "6 Wiederholungen"},
						{Exercise: umgekehrtesBankdruecken, Note: "8 Wiederholungen"},
					},
					{
						{Exercise: pistole, Note: "6 Wiederholungen pro Seite"},
						{Exercise: handstandLiegestuetz, Note: "6 Wiederholungen"},
						{Exercise: klimmzug, Note: "8 Wiederholungen"},
					},
				},
			},
		},
	},
}
