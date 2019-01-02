# Bodyweight Alexa Handler

- individuelles bodyweight training


# Commands
proxy local from 0.0.0.0:1234 to localhost 1234

    socat -d -d -x TCP4-LISTEN:1234,fork,bind=192.168.178.30 TCP4:localhost:1234

running:
-   as server

    _LAMBDA_SERVER_PORT=1234 go run src/{bodyweight.go,model.go,developtools.go}

- as client

    TESTCLIENT=true go run src/{bodyweight.go,model.go,developtools.go}