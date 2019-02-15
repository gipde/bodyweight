package training

import (
	"fmt"
)

type exercise struct {
	Name       string       `json:"name"`
	Type       trainingType `json:"type"`
	Page       int          `json:"page"`
	Difficulty int          `json:"difficulty"`
}

type tExercise struct {
	Exercise exercise `json:"exercise"`
	Note     string   `json:"note"`
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
	"Übung für die Rumpfmuskulatur",
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
	"Supersatz",
	"Zirkelintervalle",
	"Hochintensitätssätze",
}

const (
	langsame = "langsame"
	proSeite = "pro Seite"
	wdh6     = "6 Wiederholungen"
	wdh8     = "8 Wiederholungen"
	wdh10    = "10 Wiederholungen"
	wdh12    = "12 Wiederholungen"

	saetze4        = "4 Sätze pro Seite"
	armeStreamline = "Arme in Streamline Position"
	armeTHalte     = "Arme in T-Halte"
	armeVorhalte   = "Arme in Vorhalte"
	armeGekreuzt   = "mit gekreuzten Armen auf der Brust"

	halte3  = "mit 3 Sekunden Haltezeit"
	halte13 = "mit 1-3 Sekunden Haltezeit"
	halte46 = "mit 4-6 Sekunden Haltezeit"
	pause   = "pause"
	kissen  = "auf einem Kissen"

	oTiefsterPunkt  = "am tiefsten Punkt"
	oHoechsterPunkt = "am höchsten Punkt"

	bGebebeugt         = "Beine gebeugt"
	bGestreckt         = "Beine gestreckt"
	bBisGanzOben       = "bis ganz nach oben"
	bParallelBoden     = "und parallel zum Boden"
	bBisZuHaenden      = "Füße bis zu den Händen"
	bNachHinten        = "Beine nach hinten anwinklen und Fußspitzen auf einen Stuhl absetzen um das Hochdrücken zu erleichtern"
	bErhoeht           = "Füße erhöht"
	hHuefthoch         = "Hände hüfthoch"
	hKopf              = "Hände hinter dem Kopf"
	hKnie              = "Hände kniehoch"
	bisBurstbein       = "Bis zum Brustbein hochziehen"
	fuesseHinterHaende = "Füße sind hinter den Händen platziert, gehen Sie dafür einem Schritt zurürck"
	unterGriff         = "Im Untergriff"

	brusthochAbgest = "brusthoch abgestützt"
	rucksack        = "dabei am tiefsten Punkt einen Rucksack über den Kopf stemen"

	knieend         = "knieend"
	knieHeranziehen = "langsam, 2 Sek. für das Heranziehen jedes Knies"

	brustHalte          = "mit Haltezeit wenn die Brust am tiefsten Punkt der Kontraktion zwischen den Händen ist"
	knieGebeugt         = "mit gebeugten Knien"
	vorwaertsGreifen    = "mit vorwärtsgreifen"
	klimmzugOhneUnterst = "ohne Unterstützung in der Negativphase"
)

var (
	n1 = fmt.Sprintf("%s %s", knieGebeugt, vorwaertsGreifen)
)

var (
	baerenGang                                  = exercise{"Bärengang", druecken, 110, 1}
	klassischerLiegestuetz                      = exercise{"Klassischer Liegestütz", druecken, 112, 2}
	liegestuetzMitErhoehtenHaenden              = exercise{"Liegestütz mit erhöhten Händen", druecken, 112, 1}
	liegestuetzMitErhoehtenFuessen              = exercise{"Liegestütz mit erhöhten Füßen", druecken, 113, 3}
	liegestuetzMitAbstossen                     = exercise{"Liegestütz mit Abstoßen", druecken, 117, 3}
	federnderLiegestuetz                        = exercise{"Federnder Liegestütz", druecken, 117, 3}
	sturzflug                                   = exercise{"Sturzflug", druecken, 120, 3}
	breiterSturzflug                            = exercise{"Breiter Sturzflug", druecken, 121, 3}
	einarmigerLiegestuetz                       = exercise{"Einarmiger Liegestütz", druecken, 124, 4}
	einarmigerLiegestuetzMitErhoehtenHaenden    = exercise{"Einarmiger Liegestütz mit erhöhten Händen im Wechsel", druecken, 125, 4}
	einarmigerLiegestuetzMitErhoehtenFuessen    = exercise{"Einarmiger Liegestütz mit erhöhten Füßen", druecken, 126, 4}
	trizepsdip                                  = exercise{"Trizepsdip", druecken, 130, 4}
	engerLiegestuetz                            = exercise{"Enger Liegestütz", druecken, 130, 3}
	engerLiegestuetzMitErhoehtenHaenden         = exercise{"Enger Liegestütz mit erhöhten Händen", druecken, 130, 2}
	engerLiegestuetzMitErhoehtenFuessen         = exercise{"Enger Liegestütz mit erhöhten Füßen", druecken, 130, 2}
	erhoehterTrizepsstrecker                    = exercise{"Erhöhter Trizepsstrecker", druecken, 131, 4}
	trizepsdipMitUnterstuetzung                 = exercise{"Trizepsdip mit Unterstützung", druecken, 132, 2}
	militaryPress                               = exercise{"Military Press", druecken, 134, 3}
	militaryPressMitErhoehtenFuessen            = exercise{"Military Press mit erhöhten Füßen", druecken, 135, 3}
	militaryPressMitErhoehtenHaenden            = exercise{"Military Press mit erhöhten Händen", druecken, 135, 2}
	handstandLiegestuetz                        = exercise{"Handstandliegestütz", druecken, 138, 1}
	tuerziehenImUntergriff                      = exercise{"Türziehen im Untergriff", ziehen, 145, 3}
	tuerziehen                                  = exercise{"Türziehen", ziehen, 145, 1}
	einarmigesTuerziehen                        = exercise{"Einarmiges Türziehen", ziehen, 146, 2}
	umgekehrtesBankdruecken                     = exercise{"Umgekehrtes Bankdrücken", ziehen, 147, 2}
	umgekehrtesBankdrueckenImUntergriff         = exercise{"Umgekehrtes Bankdrücken im Untergriff", ziehen, 148, 3}
	klimmzug                                    = exercise{"Klimmzug", ziehen, 150, 3}
	klimmzugMitUnterstuetzung                   = exercise{"Klimmzug mit Unterstützung", ziehen, 152, 2}
	vorgebeugtesSeitlichesSchulterheben         = exercise{"Vorgebeugtes seitliches Schulterheben", ziehen, 156, 1}
	curlMitHandtuch                             = exercise{"Curl mit Handtuch", ziehen, 159, 2}
	goodMorning                                 = exercise{"Good Morning", beineUndGesaess, 164, 2}
	storch                                      = exercise{"Storch", beineUndGesaess, 165, 1}
	seitlichesKnieoeffnenImStand                = exercise{"Seitliches Knieöffnen im Stand", beineUndGesaess, 166, 1}
	kreuzschritt                                = exercise{"Kreuzschritt", beineUndGesaess, 167, 1}
	kreuzheben                                  = exercise{"Kreuzheben", beineUndGesaess, 168, 1}
	aufstehenAusDemEinbeinigenKniestand         = exercise{"Aufstehen aus dem einbeinigen Kniestand", beineUndGesaess, 169, 1}
	schwimmer                                   = exercise{"Schwimmer", beineUndGesaess, 174, 2}
	pointer                                     = exercise{"Pointer", beineUndGesaess, 174, 2}
	knieenderBeinwechsel                        = exercise{"Kniender Beinwechsel", beineUndGesaess, 175, 2}
	rumaenischesKreuzhebenAufEinemBeinImWechsel = exercise{"Rumänisches Kreuzheben auf einem Bein im Wechsel", beineUndGesaess, 176, 2}
	ausfallSchritt                              = exercise{"Ausfallschritt", beineUndGesaess, 177, 2}
	ausfallschrittNachVorneImWechsel            = exercise{"Ausfallschritt nach vorn im Wechsel", beineUndGesaess, 177, 1}
	ausfallschrittNachHintenImWechsel           = exercise{"Ausfallschritt nach hinten im Wechsel", beineUndGesaess, 178, 1}
	kniebeugeImAusfallschritt                   = exercise{"Kniebeuge im Ausfallschritt", beineUndGesaess, 178, 2}
	seitlicherAusfallschritt                    = exercise{"Seitlicher Ausfallschritt", beineUndGesaess, 179, 2}
	hueftStrecker                               = exercise{"Hüftstrecker", beineUndGesaess, 181, 2}
	engeKniebeuge                               = exercise{"Enge Kniebeuge", beineUndGesaess, 183, 1}
	einbeinigeKniebeugeMitUnterstuetzung        = exercise{"Einbeinige Kniebeuge mit Unterstützung im Wechsel", beineUndGesaess, 187, 4}
	pistole                                     = exercise{"Pistole", beineUndGesaess, 188, 4}
	einbeinigeKniebeuge                         = exercise{"Einbeinige Kniebeuge im Wechsel", beineUndGesaess, 188, 4}
	gesprungeneKniebeuge                        = exercise{"Gesprungene Kniebeuge", beineUndGesaess, 191, 1}
	kistenSprung                                = exercise{"Kistensprung", beineUndGesaess, 192, 1}
	einbeinigerHueftstrecker                    = exercise{"Einbeiniger Hüftstrecker", beineUndGesaess, 194, 2}
	seitensprung                                = exercise{"Seitsprung", beineUndGesaess, 194, 2}
	ironMike                                    = exercise{"Iron Mike", beineUndGesaess, 195, 3}
	pogoSprung                                  = exercise{"Pogo Sprung", beineUndGesaess, 199, 2}
	bergsteiger                                 = exercise{"Bergsteiger", core, 202, 2}
	gekreuzterBergsteiger                       = exercise{"gekreuzter Bergsteiger", core, 203, 2}
	kaeferIpsilateral                           = exercise{"Käfer ipsilateral", core, 204, 2}
	kaefer                                      = exercise{"Käfer", core, 204, 2}
	seitlichesHueftheben                        = exercise{"Seitliches Hüftheben", core, 204, 3}
	kaeferKontralateral                         = exercise{"Käfer kontralateral", core, 204, 2}
	kaeferUnilateral                            = exercise{"Käfer unilateral", core, 204, 2}
	hueftTwist                                  = exercise{"Hüfttwist", core, 206, 1}
	seestern                                    = exercise{"Seestern", core, 207, 1}
	bodyRock                                    = exercise{"Bodyrock", core, 208, 1}
	russischerTwist                             = exercise{"Russischer Twist", core, 213, 1}
	crunchItUp                                  = exercise{"Crunch It Up", core, 214, 1}
	schraegerCrunch                             = exercise{"Schräger Crunch", core, 215, 1}
	crunch                                      = exercise{"Crunch", core, 215, 1}
	beinheber                                   = exercise{"Beinheber", core, 217, 1}
	fahrradfahren                               = exercise{"Fahrradfahren", core, 218, 2}
	vUp                                         = exercise{"V-Up", core, 219, 2}
	schraegerVUp                                = exercise{"Schräger V-Up", core, 220, 3}
	tiefeKniebeuge                              = exercise{"Tiefe Kniebeuge", core, 221, 3}
	gestreckterBeinTwist                        = exercise{"Gestreckter Beintwist", core, 221, 3}
	beinTwist                                   = exercise{"Beintwist", core, 221, 2}
	klappmesser                                 = exercise{"Klappmesser", core, 222, 3}
	gestrecktesHaengendesBeinheben              = exercise{"Gestrecktes hängendes Beinheben", core, 223, 1}
	haengendesBeinheben                         = exercise{"Hängendes Beinheben", core, 223, 3}
	nackentrainerInBauchlage                    = exercise{"Nackentrainer in Bauchlage", core, 227, 2}
	knieheberImStehen                           = exercise{"Knieheber im Stehen", ganzKoerper, 228, 1}
	vierPhasenLiegestuetz                       = exercise{"4 Phasen Liegestütz", ganzKoerper, 229, 2}
)

var trainings = []trainingWeek{
	{
		Description: "Muskuläre Ausdauer",
		TrainingDays: []trainingDay{
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{{liegestuetzMitErhoehtenHaenden, ""},
						{tuerziehen, ""},
						{seestern, ""},
						{umgekehrtesBankdruecken, ""}},
					{{klassischerLiegestuetz, ""},
						{umgekehrtesBankdruecken, ""},
						{militaryPress, ""},
						{tuerziehen, ""}},
					{{einarmigerLiegestuetzMitErhoehtenHaenden, ""},
						{klimmzugMitUnterstuetzung, ""},
						{militaryPressMitErhoehtenFuessen, ""},
						{umgekehrtesBankdruecken, ""}},
					{{einarmigerLiegestuetzMitErhoehtenHaenden, ""},
						{klimmzugMitUnterstuetzung, ""},
						{sturzflug, ""},
						{umgekehrtesBankdruecken, "Die Füße sind erhöht. "}},
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
						{Exercise: ausfallschrittNachHintenImWechsel, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
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
						{Exercise: seestern},
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
						{Exercise: gesprungeneKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
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
						{Exercise: seestern},
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
						{Exercise: ausfallschrittNachHintenImWechsel, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
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
						{Exercise: seestern},
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
						{Exercise: gesprungeneKniebeuge, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
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
						{Exercise: seestern},
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
						{Exercise: seestern},
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
						{Exercise: seestern},
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
						{Exercise: seestern},
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
						{Exercise: seestern},
					},
					{
						{Exercise: militaryPress},
						{Exercise: klassischerLiegestuetz},
						{Exercise: engerLiegestuetz},
						{Exercise: seestern},
					},
					{
						{Exercise: einarmigerLiegestuetzMitErhoehtenHaenden},
						{Exercise: sturzflug},
						{Exercise: trizepsdip},
						{Exercise: seestern},
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
						{Exercise: seestern},
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
						{Exercise: seestern},
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
						{Exercise: seestern},
					},
					{
						{Exercise: liegestuetzMitErhoehtenFuessen, Note: "mit 1-3 Sek. Haltezeit am tiefsten Punkt"},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPressMitErhoehtenFuessen},
						{Exercise: breiterSturzflug},
						{Exercise: engerLiegestuetzMitErhoehtenFuessen},
						{Exercise: seestern},
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
