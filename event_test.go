package alexa

import (
	"encoding/json"
	"testing"
)

func TestEvent_UnmarshalJSON(t *testing.T) {
	var event Event
	rawData := []byte(`{
		"version": "1.0",
  		"session": {
    		"new": true,
    		"sessionId": "amzn1.echo-api.session.0000000-0000-0000-0000-00000000000",
    		"application": {
      			"applicationId": "amzn1.echo-sdk-ams.skill.000000-d0ed-0000-ad00-000000d00ebe"
    		},
    		"attributes": {},
    		"user": {
      			"userId": "amzn1.account.AM3B00000000000000000000000"
    		}
  		},
  		"context": {
    		"System": {
      			"application": {
        			"applicationId": "amzn1.echo-sdk-ams.skill.000000-d0ed-0000-ad00-000000d00ebe"
				},
      			"user": {
	        		"userId": "amzn1.account.AM3B00000000000000000000000"
    	  		},
      			"device": {
        			"deviceId": "string",
        			"supportedInterfaces": {
          				"AudioPlayer": {}
					}
      			}
    		},
    		"AudioPlayer": {
      			"offsetInMilliseconds": 0,
      			"playerActivity": "IDLE"
    		}
  		},
  		"request": {
    		"type": "LaunchRequest",
    		"requestId": "amzn1.echo-api.request.0000000-0000-0000-0000-00000000000",
    		"timestamp": "2015-05-13T12:34:56Z",
    		"locale": "string"
  		}
	}`)

	err := json.Unmarshal(rawData, &event)
	if err != nil {
		t.Error(err.Error())
	}

	if _, ok := event.Request.(*LaunchRequest); !ok {
		t.Error(unmarshalErr)
	}
}
