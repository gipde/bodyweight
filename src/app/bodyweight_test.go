package app

import (
	"encoding/json"
	"log"

	"testing"
)

const request = `{
    "version": "1.0",
    "session": {
        "new": false,
        "sessionId": "amzn1.echo-api.session.47ef839e-3501-4e1c-8480-9205901b7737",
        "application": {
            "applicationId": "amzn1.ask.skill.cb79a3b0-5f4b-49f4-9bcc-7ffb173f3b50"
        },
        "attributes": null,
        "user": {
            "userId": "amzn1.ask.account.AHFVOOONTT63BT24LRX7UDXKIXWQPJ2FKODA7YETDTA6FBQXCK6NGJZMT4UTFLY5VQR47TN4DCFAUTGW52LFT6676X3PJWBA7BHCWZ7SBPTIGUK4MROW5PWAWOW6W73RF4557ZM62SLDYS7CDGD3EL5AQZ53KDPRFAD6T5FH4VNO6MVBC445OEI465VDMLR6N7PCJT6H4AVWAJI",
            "accessToken": "",
            "permission": {
                "consentToken": ""
            }
        }
    },
    "context": {
        "system": {
            "apiAccessToken": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IjEifQ.eyJhdWQiOiJodHRwczovL2FwaS5hbWF6b25hbGV4YS5jb20iLCJpc3MiOiJBbGV4YVNraWxsS2l0Iiwic3ViIjoiYW16bjEuYXNrLnNraWxsLmNiNzlhM2IwLTVmNGItNDlmNC05YmNjLTdmZmIxNzNmM2I1MCIsImV4cCI6MTU0NjU2NDIxNSwiaWF0IjoxNTQ2NTYwNjE1LCJuYmYiOjE1NDY1NjA2MTUsInByaXZhdGVDbGFpbXMiOnsiY29uc2VudFRva2VuIjpudWxsLCJkZXZpY2VJZCI6ImFtem4xLmFzay5kZXZpY2UuQUdZTzZNV1hSWURTUjRHRktPWlFNN1JKVzJCTFZZN1NNTlhTUzNTNUJBVkRETFlTWkZLSlZaSEZXREIyNVZLTk9MR09RTFlGNEk3U1AyMkdEUlRZTjVCVE80QVZUN0JEQ0VGMkEzSDZNMktDWlVZUUlCMkhETEVTS05TRFZTQU1ORkNUMkgyQUZISU5BVDdOVU1NUUlPSUhURkdBIiwidXNlcklkIjoiYW16bjEuYXNrLmFjY291bnQuQUhGVk9PT05UVDYzQlQyNExSWDdVRFhLSVhXUVBKMkZLT0RBN1lFVERUQTZGQlFYQ0s2TkdKWk1UNFVURkxZNVZRUjQ3VE40RENGQVVUR1c1MkxGVDY2NzZYM1BKV0JBN0JIQ1daN1NCUFRJR1VLNE1ST1c1UFdBV09XNlc3M1JGNDU1N1pNNjJTTERZUzdDREdEM0VMNUFRWjUzS0RQUkZBRDZUNUZINFZOTzZNVkJDNDQ1T0VJNDY1VkRNTFI2TjdQQ0pUNkg0QVZXQUpJIn19.iX-WBTyjyBHkBWtM69UFgftDGW3A4tNtS794CbowmgKkMCtZ_96K37U6xGolAq8sARqj9LaUI49PXEXEloysV2lgY_WvT8siLrhHqH_XAPqoxoOV9oqmjUw7l91xFFPENpVOYA-rCJtdMX8d5lMPhQ0TEZyUTvDfr8HMIchLutOLOzJSm2_381hWHj1qnR9-cV8eJKgXw_xMo6i9b3fP0_rOBwty5rUFXZCECFROwVkmdoh2hq8n4KoiyMGWMHW6c8WwFJrSJqFFaaSULqA-TsiqLE-6EY4S1GnmgtDfcYZjWB0vwb3V0THimsVrLlW8R2JvQf3acco8L0At4J3B8Q",
            "apiEndpoint": "https://api.eu.amazonalexa.com",
            "application": {
                "applicationId": "amzn1.ask.skill.cb79a3b0-5f4b-49f4-9bcc-7ffb173f3b50"
            },
            "device": {
                "deviceId": "amzn1.ask.device.AGYO6MWXRYDSR4GFKOZQM7RJW2BLVY7SMNXSS3S5BAVDDLYSZFKJVZHFWDB25VKNOLGOQLYF4I7SP22GDRTYN5BTO4AVT7BDCEF2A3H6M2KCZUYQIB2HDLESKNSDVSAMNFCT2H2AFHINAT7NUMMQIOIHTFGA",
                "supportedInterfaces": {
                    "AudioPlayer": {}
                }
            },
            "user": {
                "userId": "amzn1.ask.account.AHFVOOONTT63BT24LRX7UDXKIXWQPJ2FKODA7YETDTA6FBQXCK6NGJZMT4UTFLY5VQR47TN4DCFAUTGW52LFT6676X3PJWBA7BHCWZ7SBPTIGUK4MROW5PWAWOW6W73RF4557ZM62SLDYS7CDGD3EL5AQZ53KDPRFAD6T5FH4VNO6MVBC445OEI465VDMLR6N7PCJT6H4AVWAJI",
                "accessToken": "",
                "permission": {
                    "consentToken": ""
                }
            }
        },
        "AudioPlayer": {
            "playerActivity": "IDLE"
        }
    },
    "request": {
        "type": "IntentRequest",
        "requestId": "amzn1.echo-api.request.c5584cad-8146-4347-8e2f-ee05d89b981c",
        "timestamp": "2019-01-04T00:10:15Z",
        "dialogState": "IN_PROGRESS",
        "locale": "de-DE",
        "intent": {
            "name": "StartTraining",
            "confirmationStatus": "NONE",
            "slots": {
                "ex_number": {
                    "name": "ex_number",
                    "value": "ersten",
                    "confirmationStatus": "CONFIRMED",
                    "resolutions": {
                        "resolutionsPerAuthority": [
                            {
                                "authority": "amzn1.er-authority.echo-sdk.amzn1.ask.skill.cb79a3b0-5f4b-49f4-9bcc-7ffb173f3b50.Ordinal",
                                "status": {
                                    "code": "ER_SUCCESS_MATCH"
                                },
                                "values": [
                                    {
                                        "value": {
                                            "id": "id1",
                                            "name": "erste"
                                        }
                                    }
                                ]
                            }
                        ]
                    }
                },
                "user": {
                    "name": "user",
                    "value": "werner",
                    "confirmationStatus": "NONE"
                }
            }
        },
        "reason": "",
        "error": {
            "type": "",
            "message": ""
        }
    }
}`

func BuildRequest(requestType string, intent string) Request {
	r := Request{
		RequestBody: RequestBody{
			Type: requestType,
		},
	}
	if requestType == INTENT_REQUEST {
		r.RequestBody.Intent = Intent{
			Name: intent,
		}
	}
	return r
}

func TestHandler(t *testing.T) {
	r, _ := HandleRequest(nil, BuildRequest(LAUNCH_REQUEST, ""))
	log.Printf("Result:  %+v\n", r)
	json, _ := json.MarshalIndent(r, "", "  ")
	log.Println("Json: ", string(json))
}

func TestResponseJson(t *testing.T) {
	r, _ := HandleRequest(nil, BuildRequest(INTENT_REQUEST, START_TRAINING))
	log.Printf("Result:  %+v\n", r)
	json, _ := json.MarshalIndent(r, "", "  ")
	log.Println("Json: ", string(json))
}

func TestUnmarshal(t *testing.T) {
	r := Request{}
	json.Unmarshal([]byte(request), &r)
	log.Printf("%+v\n\n", r)
	log.Printf("%+v\n\n", r.RequestBody.Intent.Slots["ex_number"].Resolutions.ResolutionsPerAuthority[0].Values[0]["value"].ID)
}
