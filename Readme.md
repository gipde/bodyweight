# Bodyweight Alexa Handler
[![Build Status](https://travis-ci.org/gipde/bodyweight.svg?branch=master)](https://travis-ci.org/gipde/bodyweight)
[![Coverage Status](https://coveralls.io/repos/github/gipde/bodyweight/badge.svg?branch=master)](https://coveralls.io/github/gipde/bodyweight?branch=master)

## individuelles bodyweight training

Der Skill unterstützt dich bei der Ausführung des Bodyweight Training. Alle Übungen basieren auf dem Buch "Fitte Ohne Geräte" von Mark Lauren

## TODO
- Hintergrundmusik beim Training \
  hierbei muss immer die Lautstärke an den enstprechenden Stellen reduzieret werden sonst sonst hört man die Ansagen nicht.
- Panics abfangen, und als error senden (z.b. bei Type Assertions)

## Commands
proxy local from 0.0.0.0:1234 to localhost 1234

    socat -d -d -x TCP4-LISTEN:1234,fork,bind=192.168.178.30 TCP4:localhost:1234

running the function:
- as server

        _LAMBDA_SERVER_PORT=1234 go run .


- as client
    
    this will connect to 0.0.0.0:1234 (where socat ist listen). the request will be forwarded to 127.0.0.1:1234

        TESTCLIENT=true go run .

## Sprach Modell

    Request: LaunchRequest
        starte das bodyweight training
            Willkommen ->       name   ja   -> persönliche Begrüssung
                                |
                                ------ nein   -> Name nachfragen -> Willkommen

    Intent: DefineUser
        mein Name ist {name}
        ich bin {name}
            ---- ändere den Namen

    Intent: ExplainTraining
        erkläre das heutige training
        erkläre das training
            ---- erklärt das training

    Intent: ExplainExercise
        erkläre die nächste übung (immer nur eine einzige)
        starte die nächste übung
            ----- erkläre die Übung

    Intent: StartTraining
        du kannst loslegen
        bereit
            ----- startet die nächste übung mit zeitunterstütung und übungsansage

    Intent: ExplainTrainingMethod
        erkläre Das stufenintervall
        erkläre die Intervallsätze
        erkläre den Supersatz
        erkläre die Zirkelintervalle
        erkläre den Hochintensitätssatz
            ----- erklärt die trainingsmethoden

    Inten: AMAZON.HelpIntent
        kannst du mir helfen
            ----- erkläre alle Funktionen

    Intent: AMAZON.StopIntent
        ich mag nicht mehr
        beende das training
        stoppe das Training
        Ende



                    