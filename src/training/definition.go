package training

import (
	"strings"
)

//State contains every State of the training
type State struct {
	Level trainingLevel `json:"level"`
	Week  int           `json:"week"`
	Day   int           `json:"day"`
	Unit  int           `json:"unit"`
}

type exercise struct {
	Name string       `json:"name"`
	Type trainingType `json:"type"`
	Page int          `json:"page"`
}

type noteList []string

type tExercise struct {
	Exercise exercise `json:"exercise"`
	Note     noteList `json:"note"`
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
	muscle      = "Muskuläre Ausdauer"
	power       = "Kraft"
	powerblock  = "Powerblock"
	changeblock = "Wechselblock"
)

const (
	wdh6    = "6 Wiederholungen"
	wdh8    = "8 Wiederholungen"
	wdh10   = "10 Wiederholungen"
	wdh12   = "12 Wiederholungen"
	saetze4 = "4 Sätze"

	sek3        = "mit 3 Sekunden"
	sek13       = "mit 1-3 Sekunden"
	sek46       = "mit 4-6 Sekunden"
	haltezeit   = "Haltezeit"
	pause       = "Pause"
	kontraktion = "bei der Kontraktion"

	langsame = "langsame"
	proSeite = "pro Seite"

	armeStreamline = "Arme in Streamline Position"
	armeTHalte     = "Arme in T-Halte"
	armeVorhalte   = "Arme in Vorhalte"
	armeGekreuzt   = "mit gekreuzten Armen auf der Brust"

	oTiefsterPunkt  = "am tiefsten Punkt"
	oHoechsterPunkt = "am höchsten Punkt"

	bErhoeht      = "mit erhöhten Füßen"
	bGebeugt      = "mit gebeugten Beinen"
	bGestreckt    = "mit gestreckten Beinen"
	bBisZuHaenden = "Füße bis zu den Händen"
	bNachHinten   = "Beine nach hinten anwinklen und Fußspitzen auf einen Stuhl " +
		"absetzen um das Hochdrücken zu erleichtern"
	fuesseHinterHaende = "Füße sind hinter den Händen platziert, gehen Sie dafür einem Schritt zurürck"

	kissen         = "auf einem Kissen"
	bBisGanzOben   = "bis ganz nach oben"
	bParallelBoden = "und parallel zum Boden"

	hErhoeht   = "Hände erhöht"
	hHuefthoch = "Hände etwa hüfthoch"
	hHueften   = "Hände an den Hüften"
	hKopf      = "Hände hinter dem Kopf"
	hKnie      = "Hände kniehoch"

	bisBurstbein = "bis zum Brustbein hochziehen"
	untergriff   = "im Untergriff"
	breiter      = "breit"

	brusthochAbgest = "brusthoch abgestützt"
	rucksack        = "dabei am tiefsten Punkt einen Rucksack über den Kopf stemen"

	knieHeranziehen = "langsam, 2 Sek. für das Heranziehen jedes Knies"

	brustHalte          = "mit Haltezeit wenn die Brust am tiefsten Punkt der Kontraktion zwischen den Händen ist"
	knieGebeugt         = "mit gebeugten Knien"
	vorwaertsGreifen    = "mit vorwärtsgreifen"
	klimmzugOhneUnterst = "mit Unterstützung in der Positivphase"

	mitte          = "in der Mitte"
	wechsel        = "im Wechsel"
	bauchlage      = "in Bauchlage"
	unterstuetzung = "mit Unterstützung"
	seitenwechsel  = "nach jedem Satz seitenwechsel, insgesamt 4 Sätze"

	einarmig  = "einarmig"
	einbeinig = "einbeining"

	vorne  = "nach vorne"
	hinten = "nach hinten"

	knieend   = "knieend"
	gekreuzt  = "gekreuzt"
	schraeg   = "schräg"
	tief      = "tiefe Kniebeuge"
	gestreckt = "gestreckt"

	ipsi   = "ipsilateral"
	kontra = "kontralateral"
	uni    = "unilateral"

	kniebeuge = "Kniebeuge im Ausfallschritt"
)

var (
	baerenGang                          = exercise{"Bärengang", druecken, 110}
	liegestuetz                         = exercise{"Liegestütz", druecken, 112}
	liegestuetzMitAbstossen             = exercise{"Liegestütz mit Abstoßen", druecken, 117}
	federnderLiegestuetz                = exercise{"Federnder Liegestütz", druecken, 117}
	sturzflug                           = exercise{"Sturzflug", druecken, 120}
	einarmigerLiegestuetz               = exercise{"Einarmiger Liegestütz", druecken, 124}
	engerLiegestuetz                    = exercise{"Enger Liegestütz", druecken, 130}
	erhoehterTrizepsstrecker            = exercise{"Erhöhter Trizepsstrecker", druecken, 131}
	trizepsdip                          = exercise{"Trizepsdip", druecken, 132}
	militaryPress                       = exercise{"Military Press", druecken, 134}
	handstandLiegestuetz                = exercise{"Handstandliegestütz", druecken, 138}
	tuerziehen                          = exercise{"Türziehen", ziehen, 145}
	umgekehrtesBankdruecken             = exercise{"Umgekehrtes Bankdrücken", ziehen, 147}
	klimmzug                            = exercise{"Klimmzug", ziehen, 150}
	vorgebeugtesSeitlichesSchulterheben = exercise{"Vorgebeugtes seitliches Schulterheben", ziehen, 156}
	curlMitHandtuch                     = exercise{"Curl mit Handtuch", ziehen, 159}
	goodMorning                         = exercise{"Good Morning", beineUndGesaess, 164}
	storch                              = exercise{"Storchhaltung", beineUndGesaess, 165}
	seitlichesKnieoeffnenImStand        = exercise{"Seitliches Knieöffnen im Stand", beineUndGesaess, 166}
	kreuzschritt                        = exercise{"Kreuzschritt", beineUndGesaess, 167}
	kreuzheben                          = exercise{"Kreuzheben", beineUndGesaess, 168}
	aufstehenAusDemEinbeinigenKniestand = exercise{"Aufstehen aus dem einbeinigen Kniestand", beineUndGesaess, 169}
	schwimmer                           = exercise{"Schwimmer", beineUndGesaess, 174}
	pointer                             = exercise{"Pointer", beineUndGesaess, 174}
	knieenderBeinwechsel                = exercise{"Kniender Beinwechsel", beineUndGesaess, 175}
	rumaenischesKreuzheben              = exercise{"Rumänisches Kreuzheben auf einem Bein", beineUndGesaess, 176}
	ausfallschritt                      = exercise{"Ausfallschritt", beineUndGesaess, 177}
	seitlicherAusfallschritt            = exercise{"Seitlicher Ausfallschritt", beineUndGesaess, 179}
	hueftStrecker                       = exercise{"Hüftstrecker", beineUndGesaess, 181}
	engeKniebeuge                       = exercise{"Enge Kniebeuge", beineUndGesaess, 183}
	einbeinigeKniebeuge                 = exercise{"Einbeinige Kniebeuge", beineUndGesaess, 186}
	pistole                             = exercise{"Pistole", beineUndGesaess, 188}
	gesprungeneKniebeuge                = exercise{"Gesprungene Kniebeuge", beineUndGesaess, 191}
	kistenSprung                        = exercise{"Kistensprung", beineUndGesaess, 192}
	seitsprung                          = exercise{"Seitsprung", beineUndGesaess, 194}
	ironMike                            = exercise{"Iron Mike", beineUndGesaess, 195}
	pogoSprung                          = exercise{"Pogo Sprung", beineUndGesaess, 199}
	bergsteiger                         = exercise{"Bergsteiger", core, 202}
	kaefer                              = exercise{"Käfer", core, 204}
	hueftTwist                          = exercise{"Hüfttwist", core, 206}
	seestern                            = exercise{"Seestern", core, 207}
	bodyRock                            = exercise{"Bodyrock", core, 208}
	seitlichesHueftheben                = exercise{"Seitliches Hüftheben", core, 212}
	russischerTwist                     = exercise{"Russischer Twist", core, 213}
	crunchItUp                          = exercise{"Crunch It Up", core, 214}
	crunch                              = exercise{"Crunch", core, 215}
	beinheber                           = exercise{"Beinheber", core, 217}
	fahrradfahren                       = exercise{"Fahrradfahren", core, 218}
	vUp                                 = exercise{"V-Up", core, 219}
	schraegerVUp                        = exercise{"Schräger V-Up", core, 220}
	beinTwist                           = exercise{"Beintwist", core, 221}
	klappmesser                         = exercise{"Klappmesser", core, 222}
	haengendesBeinheben                 = exercise{"Hängendes Beinheben", core, 223}
	nackentrainer                       = exercise{"Nackentrainer", core, 227}
	knieheberImStehen                   = exercise{"Knieheber im Stehen", ganzKoerper, 228}
	vierPhasenLiegestuetz               = exercise{"4 Phasen Liegestütz", ganzKoerper, 229}
)

var trainings [10]trainingWeek
var trainingDef = [...]trainingWeek{
	{
		Description: muscle,
		TrainingDays: []trainingDay{
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetz,
							Note: noteList{hErhoeht}},
						{Exercise: tuerziehen},
						{Exercise: seestern},
						{Exercise: umgekehrtesBankdruecken},
					},
					{
						{Exercise: liegestuetz},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: militaryPress},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: einarmigerLiegestuetz,
							Note: noteList{hErhoeht, wechsel}},
						{Exercise: klimmzug,
							Note: noteList{klimmzugOhneUnterst}},
						{Exercise: militaryPress,
							Note: noteList{bErhoeht}},
						{Exercise: umgekehrtesBankdruecken},
					},
					{
						{Exercise: einarmigerLiegestuetz,
							Note: noteList{hErhoeht}},
						{Exercise: klimmzug},
						{Exercise: sturzflug},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{bErhoeht}},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{wechsel}},
						{Exercise: engeKniebeuge,
							Note: noteList{armeVorhalte}},
						{Exercise: schwimmer},
					},
					{
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{wechsel}},
						{Exercise: gesprungeneKniebeuge,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: kaefer,
							Note: noteList{uni, bGestreckt}},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{unterstuetzung, wechsel}},
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: hueftStrecker},
						{Exercise: kaefer,
							Note: noteList{uni, bGebeugt}},
					},
					{
						{Exercise: pistole},
						{Exercise: ironMike},
						{Exercise: hueftStrecker,
							Note: noteList{einbeinig, wechsel}},
						{Exercise: nackentrainer,
							Note: noteList{bauchlage}},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetz,
							Note: noteList{hErhoeht}},
						{Exercise: tuerziehen},
						{Exercise: seestern},
						{Exercise: umgekehrtesBankdruecken},
					},
					{
						{Exercise: liegestuetz},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: militaryPress},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: einarmigerLiegestuetz,
							Note: noteList{hErhoeht, wechsel}},
						{Exercise: klimmzug,
							Note: noteList{klimmzugOhneUnterst}},
						{Exercise: militaryPress},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: einarmigerLiegestuetz,
							Note: noteList{hErhoeht, wechsel}},
						{Exercise: klimmzug},
						{Exercise: militaryPress,
							Note: noteList{bErhoeht}},
						{Exercise: tuerziehen,
							Note: noteList{einarmig}},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: seitlicherAusfallschritt},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{wechsel}},
						{Exercise: engeKniebeuge,
							Note: noteList{hKopf, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: crunch,
							Note: noteList{schraeg}},
					},
					{
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{wechsel}},
						{Exercise: gesprungeneKniebeuge,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: russischerTwist},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{unterstuetzung, wechsel}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{wechsel}},
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
		Description: power,
		TrainingDays: []trainingDay{
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetz},
						{Exercise: militaryPress,
							Note: noteList{hErhoeht}},
						{Exercise: engerLiegestuetz,
							Note: noteList{hErhoeht}},
						{Exercise: seestern},
					},
					{
						{Exercise: liegestuetz,
							Note: noteList{bErhoeht}},
						{Exercise: militaryPress},
						{Exercise: engerLiegestuetz},
						{Exercise: trizepsdip,
							Note: noteList{unterstuetzung, bNachHinten}},
					},
					{
						{Exercise: einarmigerLiegestuetz,
							Note: noteList{hErhoeht, wechsel}},
						{Exercise: sturzflug},
						{Exercise: militaryPress,
							Note: noteList{bErhoeht}},
						{Exercise: trizepsdip,
							Note: noteList{unterstuetzung}},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: handstandLiegestuetz},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker,
							Note: noteList{hHuefthoch}},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschritt,
							Note: noteList{kniebeuge}},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: engeKniebeuge,
							Note: noteList{armeVorhalte, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{kissen}},
					},
					{
						{Exercise: ausfallschritt,
							Note: noteList{kniebeuge, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: gesprungeneKniebeuge,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{kissen}},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{unterstuetzung}},
						{Exercise: ausfallschritt,
							Note: noteList{kniebeuge, sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: hueftStrecker,
							Note: noteList{einbeinig, kissen}},
					},
					{
						{Exercise: knieenderBeinwechsel},
						{Exercise: ausfallschritt,
							Note: noteList{kniebeuge, sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{sek13, pause}},
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
						{Exercise: tuerziehen,
							Note: noteList{untergriff}},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug,
							Note: noteList{klimmzugOhneUnterst}},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen,
							Note: noteList{untergriff}},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen,
							Note: noteList{sek13, haltezeit, kontraktion, oHoechsterPunkt}},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff}},
					},
					{
						{Exercise: klimmzug,
							Note: noteList{sek13, haltezeit, oHoechsterPunkt}},
						{Exercise: tuerziehen,
							Note: noteList{einarmig, wechsel}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{bErhoeht, untergriff}},
						{Exercise: tuerziehen,
							Note: noteList{sek46, haltezeit, kontraktion, oHoechsterPunkt}},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: beinheber},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGebeugt}},
						{Exercise: russischerTwist},
						{Exercise: crunch},
					},
					{
						{Exercise: beinheber,
							Note: noteList{armeGekreuzt}},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGestreckt}},
						{Exercise: fahrradfahren},
						{Exercise: kaefer,
							Note: noteList{kontra, bGebeugt}},
					},
					{
						{Exercise: haengendesBeinheben,
							Note: noteList{bGestreckt, bParallelBoden}},
						{Exercise: nackentrainer,
							Note: noteList{bauchlage}},
						{Exercise: vUp},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGestreckt}},
					},
					{
						{Exercise: haengendesBeinheben,
							Note: noteList{bGestreckt, bBisGanzOben}},
						{Exercise: nackentrainer,
							Note: noteList{bauchlage}},
						{Exercise: klappmesser},
						{Exercise: kaefer,
							Note: noteList{kontra, bGestreckt}},
					},
				},
			},
		},
	},
	{
		Description: powerblock,
		TrainingDays: []trainingDay{
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetz,
							Note: noteList{bErhoeht}},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPress},
						{Exercise: baerenGang},
						{Exercise: engerLiegestuetz},
						{Exercise: seestern},
					},
					{
						{Exercise: liegestuetz,
							Note: noteList{bErhoeht, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPress,
							Note: noteList{bErhoeht}},
						{Exercise: sturzflug},
						{Exercise: engerLiegestuetz,
							Note: noteList{bErhoeht}},
						{Exercise: trizepsdip},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: federnderLiegestuetz},
						{Exercise: militaryPress,
							Note: noteList{bErhoeht}},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker},
						{Exercise: liegestuetzMitAbstossen},
					},
					{
						{Exercise: einarmigerLiegestuetz,
							Note: noteList{bErhoeht}},
						{Exercise: federnderLiegestuetz},
						{Exercise: handstandLiegestuetz,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker,
							Note: noteList{hHuefthoch}},
						{Exercise: liegestuetzMitAbstossen},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: gesprungeneKniebeuge},
						{Exercise: ausfallschritt,
							Note: noteList{vorne, wechsel, sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{kissen}},
						{Exercise: engeKniebeuge,
							Note: noteList{armeStreamline, sek3, haltezeit, oTiefsterPunkt}},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{unterstuetzung, wechsel}},
						{Exercise: gesprungeneKniebeuge,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{kissen, sek13, haltezeit, oHoechsterPunkt}},
						{Exercise: kistenSprung},
					},
					{
						{Exercise: pistole},
						{Exercise: kistenSprung},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
						{Exercise: seitsprung},
					},
					{
						{Exercise: pistole,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: kistenSprung},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
						{Exercise: seitsprung},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: klimmzug,
							Note: noteList{klimmzugOhneUnterst}},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen,
							Note: noteList{sek46, haltezeit, oHoechsterPunkt}},
						{Exercise: curlMitHandtuch},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff, bGestreckt}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff}},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen,
							Note: noteList{sek46, haltezeit, kontraktion, oHoechsterPunkt}},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff, bErhoeht}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff}},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen,
							Note: noteList{einarmig}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff, sek46, haltezeit, oHoechsterPunkt}},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{bErhoeht}},
						{Exercise: tuerziehen,
							Note: noteList{einarmig, sek13, haltezeit, oHoechsterPunkt}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff, bErhoeht}},
						{Exercise: klimmzug,
							Note: noteList{bisBurstbein}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff, sek46, haltezeit, oHoechsterPunkt}},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: vUp},
						{Exercise: russischerTwist},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGestreckt}},
						{Exercise: crunchItUp},
						{Exercise: haengendesBeinheben},
						{Exercise: beinheber},
					},
					{
						{Exercise: haengendesBeinheben,
							Note: noteList{knieGebeugt}},
						{Exercise: beinTwist},
						{Exercise: hueftTwist},
						{Exercise: kaefer,
							Note: noteList{uni, bGestreckt}},
						{Exercise: vUp},
						{Exercise: russischerTwist},
					},
					{
						{Exercise: haengendesBeinheben},
						{Exercise: fahrradfahren},
						{Exercise: nackentrainer,
							Note: noteList{bauchlage}},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGebeugt}},
						{Exercise: vUp},
						{Exercise: beinTwist},
					},
					{
						{Exercise: haengendesBeinheben,
							Note: noteList{sek46, haltezeit, oHoechsterPunkt, bGestreckt, bBisGanzOben}},
						{Exercise: fahrradfahren,
							Note: noteList{knieHeranziehen}},
						{Exercise: nackentrainer,
							Note: noteList{bauchlage, sek46, pause}},
						{Exercise: kaefer,
							Note: noteList{uni, bGestreckt}},
						{Exercise: klappmesser},
						{Exercise: beinTwist},
					},
				},
			},
		},
	},
	{
		Description: changeblock,
		TrainingDays: []trainingDay{
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: militaryPress,
							Note: noteList{hErhoeht}},
						{Exercise: liegestuetz,
							Note: noteList{hErhoeht}},
						{Exercise: engerLiegestuetz,
							Note: noteList{hErhoeht}},
						{Exercise: seestern},
					},
					{
						{Exercise: militaryPress},
						{Exercise: liegestuetz},
						{Exercise: engerLiegestuetz},
						{Exercise: seestern},
					},
					{
						{Exercise: einarmigerLiegestuetz,
							Note: noteList{hErhoeht}},
						{Exercise: sturzflug},
						{Exercise: trizepsdip},
						{Exercise: seestern},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: handstandLiegestuetz},
						{Exercise: sturzflug},
						{Exercise: trizepsdip,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: gesprungeneKniebeuge},
						{Exercise: ausfallschritt,
							Note: noteList{vorne, wechsel, sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{wechsel, kissen, sek13, haltezeit, mitte}},
						{Exercise: pogoSprung},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{unterstuetzung, wechsel}},
						{Exercise: gesprungeneKniebeuge,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{wechsel, kissen, sek13, haltezeit, mitte}},
						{Exercise: kistenSprung},
					},
					{
						{Exercise: pistole},
						{Exercise: seitlichesKnieoeffnenImStand},
						{Exercise: kistenSprung},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
						{Exercise: seitsprung},
					},
					{
						{Exercise: pistole,
							Note: noteList{wechsel, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: kistenSprung},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
						{Exercise: seitsprung},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: tuerziehen},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen,
							Note: noteList{untergriff}},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug,
							Note: noteList{klimmzugOhneUnterst}},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen,
							Note: noteList{sek13, haltezeit, kontraktion, oHoechsterPunkt}},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff}},
					},
					{
						{Exercise: klimmzug,
							Note: noteList{sek13, haltezeit, oHoechsterPunkt}},
						{Exercise: tuerziehen,
							Note: noteList{einarmig}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff, bErhoeht}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff, sek46, haltezeit, kontraktion, oHoechsterPunkt}},
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
						{Exercise: engeKniebeuge,
							Note: noteList{tief, hKopf}},
					},
					{
						{Exercise: schraegerVUp,
							Note: noteList{saetze4, proSeite}},
						{Exercise: bodyRock},
						{Exercise: engeKniebeuge,
							Note: noteList{tief, armeStreamline}},
					},
					{
						{Exercise: vUp},
						{Exercise: schraegerVUp,
							Note: noteList{seitenwechsel}},
						{Exercise: russischerTwist},
						{Exercise: bergsteiger},
					},
				},
			},
			{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, wdh10, proSeite}},
						{Exercise: tuerziehen,
							Note: noteList{wdh8}},
						{Exercise: liegestuetz,
							Note: noteList{wdh6}},
					},
					{
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{wdh6}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{wechsel, wdh12}},
						{Exercise: liegestuetz,
							Note: noteList{wdh8}},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{wechsel, wdh12}},
						{Exercise: sturzflug,
							Note: noteList{wdh6}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{wdh8}},
					},
					{
						{Exercise: pistole,
							Note: noteList{wechsel, wdh12}},
						{Exercise: handstandLiegestuetz,
							Note: noteList{wdh6}},
						{Exercise: klimmzug,
							Note: noteList{wdh8}},
					},
				},
			},
		},
	},
	{
		Description: changeblock,
		TrainingDays: []trainingDay{
			{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetz,
							Note: noteList{hErhoeht, brusthochAbgest}},
						{Exercise: bodyRock},
						{Exercise: vierPhasenLiegestuetz},
					},
					{
						{Exercise: liegestuetz,
							Note: noteList{hErhoeht}},
						{Exercise: seestern},
						{Exercise: engeKniebeuge,
							Note: noteList{tief}},
					},
					{
						{Exercise: liegestuetz},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: engeKniebeuge,
							Note: noteList{tief, armeTHalte}},
					},
					{
						{Exercise: liegestuetz},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPress,
							Note: noteList{bErhoeht}},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel}},
						{Exercise: rumaenischesKreuzheben},
						{Exercise: engeKniebeuge,
							Note: noteList{armeVorhalte, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: goodMorning,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
					},
					{
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: gesprungeneKniebeuge,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{wechsel}},
					},
					{
						{Exercise: einbeinigeKniebeuge},
						{Exercise: kreuzschritt},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: hueftStrecker},
					},
					{
						{Exercise: einbeinigeKniebeuge},
						{Exercise: seitlichesKnieoeffnenImStand},
						{Exercise: ironMike},
						{Exercise: hueftStrecker,
							Note: noteList{einbeinig}},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: klimmzug,
							Note: noteList{klimmzugOhneUnterst}},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen,
							Note: noteList{sek46, haltezeit, oHoechsterPunkt}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{knieGebeugt}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff, bGestreckt}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff}},
					},
					{
						{Exercise: klimmzug},
						{Exercise: tuerziehen},
						{Exercise: tuerziehen,
							Note: noteList{sek46, haltezeit, oHoechsterPunkt}},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff, bErhoeht}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff}},
					},
					{
						{Exercise: klimmzug},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen,
							Note: noteList{einarmig}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff, bErhoeht}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff}},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: klimmzug,
							Note: noteList{sek46, haltezeit, oHoechsterPunkt}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{bErhoeht}},
						{Exercise: tuerziehen,
							Note: noteList{einarmig, sek13, haltezeit, kontraktion, oHoechsterPunkt}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff, bErhoeht}},
						{Exercise: klimmzug,
							Note: noteList{bisBurstbein}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff, sek46, haltezeit, kontraktion, oHoechsterPunkt}},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: beinheber},
						{Exercise: kaefer,
							Note: noteList{kontra, bGebeugt}},
						{Exercise: russischerTwist},
						{Exercise: seitlichesHueftheben},
					},
					{
						{Exercise: beinheber,
							Note: noteList{armeGekreuzt}},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGestreckt}},
						{Exercise: fahrradfahren},
						{Exercise: kaefer,
							Note: noteList{kontra, bGestreckt}},
					},
					{
						{Exercise: haengendesBeinheben,
							Note: noteList{gestreckt}},
						{Exercise: nackentrainer,
							Note: noteList{bauchlage}},
						{Exercise: vUp},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGestreckt}},
					},
					{
						{Exercise: haengendesBeinheben,
							Note: noteList{bGestreckt, bBisZuHaenden}},
						{Exercise: nackentrainer,
							Note: noteList{bauchlage, sek13, pause}},
						{Exercise: bergsteiger,
							Note: noteList{gekreuzt}},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGebeugt}},
					},
				},
			},
			{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, wdh10}},
						{Exercise: tuerziehen,
							Note: noteList{wdh8}},
						{Exercise: liegestuetz,
							Note: noteList{wdh6}},
					},
					{
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{wdh6}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{wechsel, wdh12}},
						{Exercise: liegestuetz,
							Note: noteList{wdh8}},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{wechsel, wdh12}},
						{Exercise: sturzflug,
							Note: noteList{wdh6}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{wdh8}},
					},
					{
						{Exercise: pistole,
							Note: noteList{wechsel, wdh12, proSeite}},
						{Exercise: handstandLiegestuetz,
							Note: noteList{wdh6}},
						{Exercise: klimmzug,
							Note: noteList{wdh8}},
					},
				},
			},
		},
	},
	{
		Description: changeblock,
		TrainingDays: []trainingDay{
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetz},
						{Exercise: militaryPress,
							Note: noteList{hErhoeht}},
						{Exercise: engerLiegestuetz,
							Note: noteList{hErhoeht}},
					},
					{
						{Exercise: liegestuetz,
							Note: noteList{bErhoeht}},
						{Exercise: militaryPress,
							Note: noteList{hErhoeht}},
						{Exercise: engerLiegestuetz,
							Note: noteList{hErhoeht}},
						{Exercise: seestern},
					},
					{
						{Exercise: einarmigerLiegestuetz,
							Note: noteList{hErhoeht}},
						{Exercise: sturzflug},
						{Exercise: militaryPress,
							Note: noteList{bErhoeht}},
						{Exercise: trizepsdip},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: handstandLiegestuetz},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker,
							Note: noteList{hHuefthoch}},
					},
				},
			},
			{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: aufstehenAusDemEinbeinigenKniestand,
							Note: noteList{armeVorhalte}},
						{Exercise: ausfallschritt,
							Note: noteList{vorne}},
						{Exercise: goodMorning},
					},
					{
						{Exercise: ironMike},
						{Exercise: seitsprung},
						{Exercise: engeKniebeuge,
							Note: noteList{tief, armeTHalte}},
					},
					{
						{Exercise: ironMike},
						{Exercise: seitsprung},
						{Exercise: engeKniebeuge,
							Note: noteList{tief, hHueften}},
					},
					{
						{Exercise: ironMike},
						{Exercise: seitsprung},
						{Exercise: pointer},
						{Exercise: engeKniebeuge,
							Note: noteList{hKopf}},
					},
				},
			},
			{
				Method: stufenIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff}},
					},
					{
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff}},
						{Exercise: tuerziehen,
							Note: noteList{untergriff}},
					},
					{
						{Exercise: klimmzug,
							Note: noteList{klimmzugOhneUnterst}},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff}},
						{Exercise: tuerziehen},
					},
					{
						{Exercise: klimmzug},
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{untergriff}},
						{Exercise: tuerziehen,
							Note: noteList{einarmig}},
					},
				},
			},
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: vUp},
						{Exercise: russischerTwist},
						{Exercise: kaefer,
							Note: noteList{kontra, bGestreckt}},
						{Exercise: seitlichesHueftheben},
						{Exercise: fahrradfahren},
						{Exercise: beinheber},
					},
					{
						{Exercise: haengendesBeinheben},
						{Exercise: beinTwist},
						{Exercise: hueftStrecker,
							Note: noteList{einbeinig, wechsel}},
						{Exercise: kaefer,
							Note: noteList{uni, bGestreckt}},
						{Exercise: vUp},
						{Exercise: russischerTwist},
					},
					{
						{Exercise: haengendesBeinheben},
						{Exercise: fahrradfahren},
						{Exercise: nackentrainer,
							Note: noteList{bauchlage, sek46, pause}},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGebeugt}},
						{Exercise: vUp},
						{Exercise: beinTwist},
					},
					{
						{Exercise: haengendesBeinheben,
							Note: noteList{bGestreckt, bBisZuHaenden, sek46, haltezeit, oHoechsterPunkt}},
						{Exercise: fahrradfahren,
							Note: noteList{langsame, wdh12, proSeite}},
						{Exercise: nackentrainer,
							Note: noteList{bauchlage, sek46, pause}},
						{Exercise: kaefer,
							Note: noteList{kontra, bGebeugt}},
						{Exercise: klappmesser},
						{Exercise: beinTwist,
							Note: noteList{langsame, wdh6, proSeite}},
					},
				},
			},
			{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, wdh10}},
						{Exercise: tuerziehen,
							Note: noteList{wdh8}},
						{Exercise: liegestuetz,
							Note: noteList{wdh6}},
					},
					{
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{wdh6, proSeite}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{wechsel, wdh12}},
						{Exercise: liegestuetz,
							Note: noteList{wdh8}},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{unterstuetzung, wechsel, wdh12, proSeite}},
						{Exercise: sturzflug,
							Note: noteList{wdh6}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{wdh8}},
					},
					{
						{Exercise: pistole,
							Note: noteList{wechsel, wdh12, proSeite}},
						{Exercise: handstandLiegestuetz,
							Note: noteList{wdh6}},
						{Exercise: klimmzug,
							Note: noteList{wdh8}},
					},
				},
			},
		},
	},
	{
		Description: changeblock,
		TrainingDays: []trainingDay{
			{
				Method: superSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: liegestuetz,
							Note: noteList{bErhoeht}},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPress},
						{Exercise: baerenGang},
						{Exercise: engerLiegestuetz},
						{Exercise: seestern},
					},
					{
						{Exercise: liegestuetz,
							Note: noteList{bErhoeht, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: liegestuetzMitAbstossen},
						{Exercise: militaryPress,
							Note: noteList{bErhoeht}},
						{Exercise: sturzflug,
							Note: noteList{breiter}},
						{Exercise: engerLiegestuetz,
							Note: noteList{bErhoeht}},
						{Exercise: seestern},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: federnderLiegestuetz},
						{Exercise: militaryPress,
							Note: noteList{bErhoeht}},
						{Exercise: sturzflug},
						{Exercise: erhoehterTrizepsstrecker},
						{Exercise: liegestuetzMitAbstossen},
					},
					{
						{Exercise: einarmigerLiegestuetz},
						{Exercise: federnderLiegestuetz},
						{Exercise: handstandLiegestuetz},
						{Exercise: sturzflug,
							Note: noteList{brustHalte}},
						{Exercise: erhoehterTrizepsstrecker,
							Note: noteList{hKnie}},
						{Exercise: liegestuetzMitAbstossen},
					},
				},
			},
			{
				Method: intervallSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschritt,
							Note: noteList{kniebeuge}},
						{Exercise: seitlicherAusfallschritt},
						{Exercise: engeKniebeuge,
							Note: noteList{armeTHalte, sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{kissen}},
					},
					{
						{Exercise: ausfallschritt,
							Note: noteList{kniebeuge, sek13, haltezeit, oTiefsterPunkt}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: gesprungeneKniebeuge,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: rumaenischesKreuzheben,
							Note: noteList{kissen}},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{unterstuetzung}},
						{Exercise: ausfallschritt,
							Note: noteList{kniebeuge, sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{sek46, haltezeit, oTiefsterPunkt}},
						{Exercise: hueftStrecker,
							Note: noteList{einbeinig}},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{knieend}},
						{Exercise: ausfallschritt,
							Note: noteList{kniebeuge, rucksack}},
						{Exercise: ironMike},
						{Exercise: knieenderBeinwechsel},
					},
				},
			},
			{
				Method: hochIntensitaetsSatz,
				Exercises: [][]tExercise{
					{
						{Exercise: tuerziehen,
							Note: noteList{fuesseHinterHaende}},
						{Exercise: kreuzheben},
						{Exercise: curlMitHandtuch},
					},
					{
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: engeKniebeuge,
							Note: noteList{tief, armeStreamline}},
					},
					{
						{Exercise: umgekehrtesBankdruecken},
						{Exercise: tuerziehen},
						{Exercise: engeKniebeuge,
							Note: noteList{tief, hKopf}},
					},
					{
						{Exercise: klimmzug,
							Note: noteList{klimmzugOhneUnterst}},
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
						{Exercise: kaefer,
							Note: noteList{ipsi, bGestreckt}},
						{Exercise: beinheber},
						{Exercise: kaefer,
							Note: noteList{kontra, bGebeugt}},
					},
					{
						{Exercise: fahrradfahren},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGestreckt}},
						{Exercise: bodyRock,
							Note: noteList{vorwaertsGreifen}},
						{Exercise: beinTwist,
							Note: noteList{gestreckt}},
					},
					{
						{Exercise: schraegerVUp,
							Note: noteList{bGebeugt}},
						{Exercise: kreuzheben},
						{Exercise: beinheber,
							Note: noteList{armeGekreuzt}},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGebeugt}},
					},
					{
						{Exercise: klappmesser},
						{Exercise: nackentrainer,
							Note: noteList{bauchlage}},
						{Exercise: schraegerVUp,
							Note: noteList{wechsel}},
						{Exercise: kaefer,
							Note: noteList{ipsi, bGestreckt}},
					},
				},
			},
			{
				Method: zirkelIntervall,
				Exercises: [][]tExercise{
					{
						{Exercise: ausfallschritt,
							Note: noteList{hinten, wechsel, wdh10, proSeite}},
						{Exercise: tuerziehen,
							Note: noteList{wdh8}},
						{Exercise: liegestuetz,
							Note: noteList{wdh6}},
					},
					{
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{wdh6, proSeite}},
						{Exercise: seitlicherAusfallschritt,
							Note: noteList{wechsel, wdh12}},
						{Exercise: liegestuetz,
							Note: noteList{wdh8}},
					},
					{
						{Exercise: einbeinigeKniebeuge,
							Note: noteList{unterstuetzung, wechsel, wdh12, proSeite}},
						{Exercise: sturzflug,
							Note: noteList{wdh6}},
						{Exercise: umgekehrtesBankdruecken,
							Note: noteList{wdh8}},
					},
					{
						{Exercise: pistole,
							Note: noteList{wechsel, wdh12, proSeite}},
						{Exercise: handstandLiegestuetz,
							Note: noteList{wdh6}},
						{Exercise: klimmzug,
							Note: noteList{wdh8}},
					},
				},
			},
		},
	},
}

func init() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			trainings[i*2+j] = trainingDef[i]
		}
	}
	for i := 3; i < 7; i++ {
		trainings[i+3] = trainingDef[i]
	}
}

func (t trainingMethod) name() string {
	return trainingMethods[t]
}

func (t trainingType) name() string {
	return trainingTypes[t]
}

func (t trainingLevel) name() string {
	return trainingLevels[t]
}

func (n noteList) text() string {
	return strings.Join(n, " ")
}
