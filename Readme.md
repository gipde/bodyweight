# Bodyweight Alexa Handler
[![Build Status](https://travis-ci.org/gipde/bodyweight.svg?branch=master)](https://travis-ci.org/gipde/bodyweight)
[![Coverage Status](https://coveralls.io/repos/github/gipde/bodyweight/badge.svg?branch=master)](https://coveralls.io/github/gipde/bodyweight?branch=master)

- individuelles bodyweight training


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

Willkommen ->       name   ja   -> persönliche begrüssung
                    |
                            nein   -> Name nachfragen -> Willkommen
                    