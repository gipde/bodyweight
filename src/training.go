package main

var Trainings = []Training{
	Training{
		Title: "Basisprogramm für Einstiger",
		TrainingWeeks: []TrainingWeek{
			TrainingWeek{
				Description: "Muskuläre Ausdauer",
				TrainingDays: []TrainingDay{
					TrainingDay{
						Kind:   "Drücken/Ziehen",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Ligestütz mit erhöhten Händen",
								Page: 113,
							},
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name: "Seestern",
								Page: 207,
							},
							Exercise{
								Name: "Umgekehrtes Bankdrücken",
								Page: 147,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß/Core",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Ausfallschritt nach hinten im Wechsel",
								Page: 178,
							},
							Exercise{
								Name: "Rumänisches Kruezheben auf einem Bein im Wechsel",
								Page: 176,
							},
							Exercise{
								Name:    "Enge Kniebeuge",
								Page:    183,
								Comment: "Arme in Vorhalte",
							},
							Exercise{
								Name: "Schwimmer",
								Page: 174,
							},
						},
					},
					TrainingDay{
						Kind:   "Drücken/Ziehen",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Ligestütz mit erhöhten Händen",
								Page: 113,
							},
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name: "Seestern",
								Page: 207,
							},
							Exercise{
								Name: "Umgekehrtes Bankdrücken",
								Page: 147,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß/Core",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Seitlicher Ausfallschritt",
								Page: 179,
							},
							Exercise{
								Name: "Rumänisches Kruezheben auf einem Bein im Wechsel",
								Page: 176,
							},
							Exercise{
								Name:    "Enge Kniebeuge",
								Page:    183,
								Comment: "Hände hinter dem Kopf. 1-3 Sek. Haltezeit am teifsten Punkt",
							},
							Exercise{
								Name: "Schräger Crunch",
								Page: 174,
							},
						},
					},
				},
			},
			TrainingWeek{
				Description: "Muskuläre Ausdauer",
				TrainingDays: []TrainingDay{
					TrainingDay{
						Kind:   "Drücken/Ziehen",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Ligestütz mit erhöhten Händen",
								Page: 113,
							},
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name: "Seestern",
								Page: 207,
							},
							Exercise{
								Name: "Umgekehrtes Bankdrücken",
								Page: 147,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß/Core",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Ausfallschritt nach hinten im Wechsel",
								Page: 178,
							},
							Exercise{
								Name: "Rumänisches Kruezheben auf einem Bein im Wechsel",
								Page: 176,
							},
							Exercise{
								Name:    "Enge Kniebeuge",
								Page:    183,
								Comment: "Arme in Vorhalte",
							},
							Exercise{
								Name: "Schwimmer",
								Page: 174,
							},
						},
					},
					TrainingDay{
						Kind:   "Drücken/Ziehen",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Ligestütz mit erhöhten Händen",
								Page: 113,
							},
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name: "Seestern",
								Page: 207,
							},
							Exercise{
								Name: "Umgekehrtes Bankdrücken",
								Page: 147,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß/Core",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Seitlicher Ausfallschritt",
								Page: 179,
							},
							Exercise{
								Name: "Rumänisches Kruezheben auf einem Bein im Wechsel",
								Page: 176,
							},
							Exercise{
								Name:    "Enge Kniebeuge",
								Page:    183,
								Comment: "Hände hinter dem Kopf. 1-3 Sek. Haltezeit am teifsten Punkt",
							},
							Exercise{
								Name: "Schräger Crunch",
								Page: 174,
							},
						},
					},
				},
			},
			TrainingWeek{
				Description: "Kraft",
				TrainingDays: []TrainingDay{
					TrainingDay{
						Kind:   "Drücken",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Klassischer Liegestütz",
								Page: 112,
							},
							Exercise{
								Name: "Military Press mit erhöhten Händen",
								Page: 135,
							},
							Exercise{
								Name: "Enger Liegestütz mit erhöhten Händen",
								Page: 130,
							},
							Exercise{
								Name: "Seestern",
								Page: 207,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Kniebeuge im Ausfallschritt",
								Page: 178,
							},
							Exercise{
								Name: "Seitlicher Ausfallschritt",
								Page: 179,
							},
							Exercise{
								Name:    "Enge Kniebeuge",
								Page:    183,
								Comment: "Arme in T-Halte mit 1-3 Sek. Haltezeit am tiefsten Punkt",
							},
							Exercise{
								Name: "Rumänisches Kreuzheben auf einem Bein auf einem Kissen",
								Page: 176,
							},
						},
					},
					TrainingDay{
						Kind:   "Ziehen",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name: "Umgekehrtes Bankdrücken",
								Page: 147,
							},
							Exercise{
								Name: "Türziehen im Untergriff",
								Page: 146,
							},
							Exercise{
								Name: "Curl mit Handtuch",
								Page: 159,
							},
						},
					},
					TrainingDay{
						Kind:   "Core",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Beinheber",
								Page: 217,
							},
							Exercise{
								Name:    "Käfer ipsilateral",
								Page:    204,
								Comment: "Beine gebeugt",
							},
							Exercise{
								Name: "Russischer Twist",
								Page: 213,
							},
							Exercise{
								Name: "Crunch",
								Page: 215,
							},
						},
					},
				},
			},
			TrainingWeek{
				Description: "Kraft",
				TrainingDays: []TrainingDay{
					TrainingDay{
						Kind:   "Drücken",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Klassischer Liegestütz",
								Page: 112,
							},
							Exercise{
								Name: "Military Press mit erhöhten Händen",
								Page: 135,
							},
							Exercise{
								Name: "Enger Liegestütz mit erhöhten Händen",
								Page: 130,
							},
							Exercise{
								Name: "Seestern",
								Page: 207,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Kniebeuge im Ausfallschritt",
								Page: 178,
							},
							Exercise{
								Name: "Seitlicher Ausfallschritt",
								Page: 179,
							},
							Exercise{
								Name:    "Enge Kniebeuge",
								Page:    183,
								Comment: "Arme in T-Halte mit 1-3 Sek. Haltezeit am tiefsten Punkt",
							},
							Exercise{
								Name: "Rumänisches Kreuzheben auf einem Bein auf einem Kissen",
								Page: 176,
							},
						},
					},
					TrainingDay{
						Kind:   "Ziehen",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name: "Umgekehrtes Bankdrücken",
								Page: 147,
							},
							Exercise{
								Name: "Türziehen im Untergriff",
								Page: 146,
							},
							Exercise{
								Name: "Curl mit Handtuch",
								Page: 159,
							},
						},
					},
					TrainingDay{
						Kind:   "Core",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Beinheber",
								Page: 217,
							},
							Exercise{
								Name:    "Käfer ipsilateral",
								Page:    204,
								Comment: "Beine gebeugt",
							},
							Exercise{
								Name: "Russischer Twist",
								Page: 213,
							},
							Exercise{
								Name: "Crunch",
								Page: 215,
							},
						},
					},
				},
			},
			TrainingWeek{
				Description: "Powerblock",
				TrainingDays: []TrainingDay{
					TrainingDay{
						Kind:   "Drücken",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Ligestütz mit erhöhten Füßen",
								Page: 113,
							},
							Exercise{
								Name: "Liegestütz mit Abstoßen",
								Page: 117,
							},
							Exercise{
								Name: "Military Press",
								Page: 134,
							},
							Exercise{
								Name: "Bärengang",
								Page: 110,
							},
							Exercise{
								Name: "Enger Liegestütz",
								Page: 130,
							},
							Exercise{
								Name: "Seestern",
								Page: 207,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name:    "Ausfallschritt nach hinten im Wechsel",
								Page:    178,
								Comment: " mit 4-6 Sek. Haltezeit am tiefsten Punkt",
							},
							Exercise{
								Name: "Gesprungene Kniebeuge",
								Page: 191,
							},
							Exercise{
								Name:    "Ausfallschritt nach vorn im Wechsel",
								Page:    177,
								Comment: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
							},
							Exercise{
								Name: "Seitlicher Ausfallschritt",
								Page: 179,
							},
							Exercise{
								Name: "Rumänisches Kreuzheben auf einem Bein auf einem Kissen",
								Page: 176,
							},
							Exercise{
								Name:    "Enge Kniebeuge",
								Page:    183,
								Comment: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt",
							},
						},
					},
					TrainingDay{
						Kind:   "Ziehen",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name:    "Klimmzug im Obergriff mit Unterstützung, ohne Unterstützung in der Negativphase",
								Page:    150,
								Comment: "stellen Sie sich auf einen Stuhl oder springen Sie hoch und konzentrieren Sie sich auf das Absenken",
							},
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name:    "Türziehen",
								Page:    145,
								Comment: "mit 4-6 Sek. Haltezeit am höchsten Punkt",
							},
							Exercise{
								Name: "Curl mit Handtuch",
								Page: 159,
							},
							Exercise{
								Name:    "Umgekehrtes Bankdrücken im Untergriff",
								Page:    148,
								Comment: "Beine gestreckt",
							},
							Exercise{
								Name: "Türziehen im Untergriff",
								Page: 146,
							},
						},
					},
					TrainingDay{
						Kind:   "Core",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name: "V-Up",
								Page: 219,
							},
							Exercise{
								Name: "Russischer Twist",
								Page: 213,
							},
							Exercise{
								Name:    "Käfer ipsilateral",
								Page:    204,
								Comment: "Beine gestreckt",
							},
							Exercise{
								Name: "Crunch it Up",
								Page: 214,
							},
							Exercise{
								Name: "Hängendes Beinheben",
								Page: 223,
							},
							Exercise{
								Name: "Beinheber",
								Page: 217,
							},
						},
					},
				},
			},
			TrainingWeek{
				Description: "Powerblock",
				TrainingDays: []TrainingDay{
					TrainingDay{
						Kind:   "Drücken",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Ligestütz mit erhöhten Füßen",
								Page: 113,
							},
							Exercise{
								Name: "Liegestütz mit Abstoßen",
								Page: 117,
							},
							Exercise{
								Name: "Military Press",
								Page: 134,
							},
							Exercise{
								Name: "Bärengang",
								Page: 110,
							},
							Exercise{
								Name: "Enger Liegestütz",
								Page: 130,
							},
							Exercise{
								Name: "Seestern",
								Page: 207,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name:    "Ausfallschritt nach hinten im Wechsel",
								Page:    178,
								Comment: " mit 4-6 Sek. Haltezeit am tiefsten Punkt",
							},
							Exercise{
								Name: "Gesprungene Kniebeuge",
								Page: 191,
							},
							Exercise{
								Name:    "Ausfallschritt nach vorn im Wechsel",
								Page:    177,
								Comment: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
							},
							Exercise{
								Name: "Seitlicher Ausfallschritt",
								Page: 179,
							},
							Exercise{
								Name: "Rumänisches Kreuzheben auf einem Bein auf einem Kissen",
								Page: 176,
							},
							Exercise{
								Name:    "Enge Kniebeuge",
								Page:    199,
								Comment: "Arme in Streamline-Position mit 3 Sek. Haltezeit am tiefsten Punkt",
							},
						},
					},
					TrainingDay{
						Kind:   "Ziehen",
						Method: "Interfallsätze",
						Exercises: []Exercise{
							Exercise{
								Name:    "Klimmzug im Obergriff mit Unterstützung, ohne Unterstützung in der Negativphase",
								Page:    150,
								Comment: "stellen Sie sich auf einen Stuhl oder springen Sie hoch und konzentrieren Sie sich auf das Absenken",
							},
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name:    "Türziehen",
								Page:    145,
								Comment: "mit 4-6 Sek. Haltezeit am höchsten Punkt",
							},
							Exercise{
								Name: "Curl mit Handtuch",
								Page: 159,
							},
							Exercise{
								Name:    "Umgekehrtes Bankdrücken im Untergriff",
								Page:    148,
								Comment: "Beine gestreckt",
							},
							Exercise{
								Name: "Türziehen im Untergriff",
								Page: 146,
							},
						},
					},
					TrainingDay{
						Kind:   "Core",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name: "V-Up",
								Page: 219,
							},
							Exercise{
								Name: "Russischer Twist",
								Page: 213,
							},
							Exercise{
								Name:    "Käfer ipsilateral",
								Page:    204,
								Comment: "Beine gestreckt",
							},
							Exercise{
								Name: "Crunch it Up",
								Page: 214,
							},
							Exercise{
								Name: "Hängendes Beinheben",
								Page: 223,
							},
							Exercise{
								Name: "Beinheber",
								Page: 217,
							},
						},
					},
				},
			},
			TrainingWeek{
				Description: "Wechselblock",
				TrainingDays: []TrainingDay{
					TrainingDay{
						Kind:   "Drücken",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Military Press mit erhöhten Händen",
								Page: 135,
							},
							Exercise{
								Name: "Liegestütz mit erhöhten Händen",
								Page: 113,
							},
							Exercise{
								Name: "Enger Liegestütz mit erhöhten Händen",
								Page: 130,
							},
							Exercise{
								Name: "Seestern",
								Page: 207,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name:    "Ausfallschritt nach hinten im Wechsel",
								Page:    178,
								Comment: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
							},
							Exercise{
								Name: "Gesprungene Kniebeuge",
								Page: 191,
							},
							Exercise{
								Name:    "Ausfallschritt nach vorn im Wechsel",
								Page:    177,
								Comment: "mit 4-6 Sek. Haltezeit am tiefsten Punkt",
							},
							Exercise{
								Name: "Seitlicher Ausfallschritt",
								Page: 179,
							},
							Exercise{
								Name:    "Rumänisches Kreuzheben auf einem Bein im Wechsel auf einem Kissen",
								Page:    176,
								Comment: "mit 1-3 Sek. Haltezeit in der Mitte",
							},
							Exercise{
								Name: "Pogo-Sprung",
								Page: 199,
							},
						},
					},
					TrainingDay{
						Kind:   "Ziehen",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name: "Umgekehrtes Bankdrücken",
								Page: 147,
							},
							Exercise{
								Name: "Türziehen im Untergriff",
								Page: 146,
							},
							Exercise{
								Name: "Curl mit Handtuch",
								Page: 159,
							},
						},
					},
					TrainingDay{
						Kind:   "Core",
						Method: "Hochintensitätssätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Russischer Twist",
								Page: 213,
							},
							Exercise{
								Name: "Hüfttwist",
								Page: 206,
							},
							Exercise{
								Name: "Knieheben im Stehen",
								Page: 228,
							},
						},
					},
					TrainingDay{
						Kind:   "Alles",
						Method: "Zirkelintervalle",
						Exercises: []Exercise{
							Exercise{
								Name:    "Ausfallschritt nach hinten im Wechsel",
								Page:    178,
								Comment: "10 Wiederholungen pro Seite",
							},
							Exercise{
								Name:    "Türziehen",
								Page:    145,
								Comment: "8 Wiederholungen",
							},
							Exercise{
								Name:    "klassischer Liegestütz",
								Page:    112,
								Comment: "6 Wiederholungen",
							},
						},
					},
				},
			},
			TrainingWeek{
				Description: "Wechselblock",
				TrainingDays: []TrainingDay{
					TrainingDay{
						Kind:   "Drücken",
						Method: "Hochintensitätssätze",
						Exercises: []Exercise{
							Exercise{
								Name:    "Ligestütz mit erhöhten Händen",
								Page:    113,
								Comment: "brusthoch abgestützt",
							},
							Exercise{
								Name: "Bodyrock",
								Page: 208,
							},
							Exercise{
								Name: "4-Phasen-Liegestütz",
								Page: 229,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Ausfallschritt nach hinten im Wechsel",
								Page: 178,
							},
							Exercise{
								Name: "Rumänisches Kruezheben auf einem Bein im Wechsel",
								Page: 176,
							},
							Exercise{
								Name:    "Enge Kniebeuge",
								Page:    183,
								Comment: "Arme in Vorhalte mit 1-3 Sek. Haltezeit am tiefsten Punkt",
							},
							Exercise{
								Name:    "Good Morning",
								Page:    164,
								Comment: "mit 1-3 Sek. Haltezeit am tiefsten Punkt",
							},
						},
					},
					TrainingDay{
						Kind:   "Ziehen",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name:    "Klimmzug im Obergriff mit Unterstützung, ohne Unterstützung in der Negativphase",
								Page:    150,
								Comment: "stellen Sie sich auf einen Stuhl oder springen Sie hoch und konzentrieren Sie sich auf das Absenken",
							},
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name:    "Türziehen",
								Page:    145,
								Comment: "mit 4-6 Sek. Haltenzeit am höchsten Punkt",
							},
							Exercise{
								Name: "Umgekehrtes Bankdrücken mit gebeugten Kniene",
								Page: 147,
							},
							Exercise{
								Name:    "Umgekehrtes Bankdrücken im Untergriff",
								Page:    148,
								Comment: "Beine gestreckt",
							},
							Exercise{
								Name: "Türziehen im Untergriff",
								Page: 146,
							},
						},
					},
					TrainingDay{
						Kind:   "Core",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Beinheber",
								Page: 217,
							},
							Exercise{
								Name:    "Käfer kontralateral",
								Page:    204,
								Comment: "Beine gebeugt",
							},
							Exercise{
								Name: "Russischer Twist",
								Page: 213,
							},
							Exercise{
								Name: "Seitliches Hüftheben",
								Page: 212,
							},
						},
					},
					TrainingDay{
						Kind:   "Alles",
						Method: "Zirkelintervalle",
						Exercises: []Exercise{
							Exercise{
								Name:    "Ausfallschritt nach hinten im Wechsel",
								Page:    178,
								Comment: "10 Wiederholgen",
							},
							Exercise{
								Name:    "Türziehen",
								Page:    145,
								Comment: "8 Wiederholungen",
							},
							Exercise{
								Name:    "klassischer Liegestütz",
								Page:    112,
								Comment: "6 Wiederholungen",
							},
						},
					},
				},
			},
			TrainingWeek{
				Description: "Wechselblock",
				TrainingDays: []TrainingDay{
					TrainingDay{
						Kind:   "Drücken",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Liegestütz",
								Page: 111,
							},
							Exercise{
								Name: "Military Press mit erhöhten Händen",
								Page: 135,
							},
							Exercise{
								Name: "Enger Liegestütz mit erhöhten Händen",
								Page: 130,
							},
						},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß",
						Method: "Hochintensitätssätze",
						Exercises: []Exercise{
							Exercise{
								Name:    "Aufstehen aus dem einbeinigen Kniestand",
								Page:    169,
								Comment: "Arme in Vorhalte",
							},
							Exercise{
								Name: "Ausfallschritt",
								Page: 177,
							},
							Exercise{
								Name: "Good Morning",
								Page: 164,
							},
						},
					},
					TrainingDay{
						Kind:   "Ziehen",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Umgekehrtes Bankdrücken",
								Page: 147,
							},
							Exercise{
								Name: "Türziehen",
								Page: 145,
							},
							Exercise{
								Name: "Umgekehrtes Bankdrücken im Untergriff",
								Page: 148,
							},
							Exercise{
								Name: "Türziehen im Untergriff",
								Page: 146,
							},
						},
					},
					TrainingDay{
						Kind:   "Core",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name: "V-Up",
								Page: 219,
							},
							Exercise{
								Name: "Russischer Twist",
								Page: 213,
							},
							Exercise{
								Name:    "Käfer kontralateral",
								Page:    204,
								Comment: "Beine gestreckt",
							},
							Exercise{
								Name: "Seitliches Hüftheben",
								Page: 212,
							},
							Exercise{
								Name: "Fahrradfahren",
								Page: 218,
							},
							Exercise{
								Name: "Beinheber",
								Page: 217,
							},
						},
					},
					TrainingDay{
						Kind:   "Alles",
						Method: "Zirkelintervalle",
						Exercises: []Exercise{
							Exercise{
								Name:    "Ausfallschritt nach hinten im Wechsel",
								Page:    178,
								Comment: "10 Wiederholugen pro Seite",
							},
							Exercise{
								Name:    "Türziehen",
								Page:    176,
								Comment: "8 Wiederholungen",
							},
							Exercise{
								Name:    "klassischer Liegestütz",
								Page:    112,
								Comment: "6 Wiederholungen",
							},
						},
					},
				},
			},
			TrainingWeek{
				Description: "Wechselblock",
				TrainingDays: []TrainingDay{
					TrainingDay{
						Kind:   "Drücken",
						Method: "Supersätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Liegestütz mit erhöhten Füßen",
								Page: 113,
							},
							Exercise{
								Name: "Liegestütz mit Abstoßen",
								Page: 117,
							},
							Exercise{
								Name: "Military Press",
								Page: 134,
							},
							Exercise{
								Name: "Bärengang",
								Page: 110,
							},
							Exercise{
								Name: "Enger Liegestütz",
								Page: 130,
							},
							Exercise{
								Name: "Seestern",
								Page: 207,
							}},
					},
					TrainingDay{
						Kind:   "Beine und Gesäß",
						Method: "Intervallsätze",
						Exercises: []Exercise{
							Exercise{
								Name: "Kniebeuge im Ausfallschritt",
								Page: 178,
							},
							Exercise{
								Name: "Seitlicher Ausfallschritt",
								Page: 179,
							},
							Exercise{
								Name:    "Enge Kniebeuge",
								Page:    183,
								Comment: "Arme in T-Halte mit 4-6 Sek. Haltezeit am tiefsten Punkt",
							},
							Exercise{
								Name: "Rumänisches Kreuzheben auf einem Bein auf einem Kissen",
								Page: 176,
							},
						},
					},
					TrainingDay{
						Kind:   "Ziehen",
						Method: "Hochintensitätssätze",
						Exercises: []Exercise{
							Exercise{
								Name:    "Türziehen",
								Page:    145,
								Comment: "Füße sind hinter den Händen plaziert, gehen Sie dafür einen Schritt zurück",
							},
							Exercise{
								Name: "Kreuzheben",
								Page: 168,
							},
							Exercise{
								Name: "Curl mit Handtuch",
								Page: 159,
							},
						},
					},
					TrainingDay{
						Kind:   "Core",
						Method: "Stufenintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Crunch it Up",
								Page: 214,
							},
							Exercise{
								Name:    "Käfer ipsilateral",
								Page:    204,
								Comment: "Beine gestreckt",
							},
							Exercise{
								Name: "Beinheber",
								Page: 217,
							},
							Exercise{
								Name:    "Käfer kontralateral",
								Page:    204,
								Comment: "Beine gebeugt",
							},
						},
					},
					TrainingDay{
						Kind:   "Alles",
						Method: "Zirkelintervalle",
						Exercises: []Exercise{
							Exercise{
								Name: "Ausfallschritt nach hinten im Wechsel",
								Page: 178,
								Comment:"10 Wiederholungen pro Seite",
							},
							Exercise{
								Name: "Türziehen",
								Page: 145,
								Comment:"8 Wiederholungen",
							},
							Exercise{
								Name: "klassischer Liegestütz",
								Page: 112,
								Comment:"6 Wiederholungen",
							},
						},
					},
				},
			},
		},
	},
}
