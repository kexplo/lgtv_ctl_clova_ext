package protocol_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"lgtv_ctl_clova_ext/internal/protocol"
	"testing"
)

func TestUnmarshalCEKMessage(t *testing.T) {
	testPayload := `
{
    "version": "0.1.0",
    "session": {
        "sessionId": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        "sessionAttributes": {},
        "user": {
            "userId": "xxxxxxxxxxxxxxxxxxxxxx",
            "accessToken": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
        },
        "new": true
    },
    "context": {
        "System": {
            "application": {
                "applicationId": "xxx.xxxxxxx.xxxx"
            },
            "user": {
                "userId": "xxxxxxxxxxxxxxxxxxxxxx",
                "accessToken": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
            },
            "device": {
                "deviceId": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
                "display": {
                    "size": "l100",
                    "orientation": "landscape",
                    "dpi": 96,
                    "contentLayer": {
                        "width": 640,
                        "height": 360
                    }
                }
            }
        }
    },
    "request": {
        "type": "IntentRequest",
        "intent": {
            "name": "LaunchApplication",
            "slots": {
                "app_name": {
                    "name": "app_name",
                    "value": "테스트"
                }
            }
        }
    }
}
`
	var cekMessage protocol.CEKMessage
	if err := json.Unmarshal([]byte(testPayload), &cekMessage); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, cekMessage.Session.Id, "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
	assert.Equal(t, cekMessage.Session.User.Id, "xxxxxxxxxxxxxxxxxxxxxx")
	assert.Equal(t, cekMessage.Session.User.AccessToken, "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
	assert.Equal(t, cekMessage.Session.New, true)

	assert.Equal(t, cekMessage.Request.Type, "IntentRequest")
	assert.Equal(t, cekMessage.Request.Intent.Name, "LaunchApplication")
	assert.Equal(t, cekMessage.Request.Intent.Slots["app_name"]["name"], "app_name")
	assert.Equal(t, cekMessage.Request.Intent.Slots["app_name"]["value"], "테스트")
}

func TestMakeMessageResponse(t *testing.T) {
	resp, _ := protocol.MakeMessageResponse()
	marshal, _ := json.Marshal(&resp)
	fmt.Println(string(marshal))
}