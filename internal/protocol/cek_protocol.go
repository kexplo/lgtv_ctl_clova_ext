package protocol

/*
{
  "context": {
    "AudioPlayer": {
      "offsetInMilliseconds": {{number}},
      "playerActivity": {{string}},
      "stream": {{AudioStreamInfoObject}},
      "totalInMilliseconds": {{number}},
    },
    "System": {
      "application": {
        "applicationId": {{string}}
      },
      "device": {
        "deviceId": {{string}},
        "display": {
          "contentLayer": {
            "width": {{number}},
            "height": {{number}}
          },
          "dpi": {{number}},
          "orientation": {{string}},
          "size": {{string}}
        }
      },
      "user": {
        "userId": {{string}},
        "accessToken": {{string}}
      }
    }
  },
  "request": {{object}},
  "session": {
    "new": {{boolean}},
    "sessionAttributes": {{object}},
    "sessionId": {{string}},
    "user": {
      "userId": {{string}},
      "accessToken": {{string}}
    }
  },
  "version": {{string}}
}
*/

type CEKSession struct {
	New        bool        `json:"new"`
	Attributes interface{} `json:"sessionAttributes"`
	Id         string      `json:"sessionId"`
	User       struct {
		Id          string `json:"userId"`
		AccessToken string `json:"accessToken"`
	} `json:"user"`
}

type CEKRequest struct {
	Type   string `json:"type"`
	Intent struct {
		Name  string                       `json:"name"`
		Slots map[string]map[string]string `json:"slots"`
	} `json:"intent"`
}

type CEKMessage struct {
	CEKContext interface{} `json:"context"`
	Request    CEKRequest  `json:"request"`
	Session    CEKSession  `json:"session"`
	Version    string      `json:"version"`
}

// ----------------------------------------------------------------------------

/*
{
  "response": {
    "card": {{object}},
    "directives": [
      {
        "header": {
          "messageId": {{string}},
          "name": {{string}},
          "namespace": {{string}}
        },
        "payload": {{object}}
      }
    ],
    "outputSpeech": {
      "type": {{string}},
      "values": {{SpeechInfoObject|SpeechInfoObject array}},
      "brief": {{SpeechInfoObject}},
      "verbose": {
        "type": {{string}},
        "values": {{SpeechInfoObject|SpeechInfoObject array}},
      }
    },
    "reprompt": {
      "outputSpeech": {
        "type": {{string}},
        "values": {{SpeechInfoObject|SpeechInfoObject array}},
        "brief": {{SpeechInfoObject}},
        "verbose": {
          "type": {{string}},
          "values": {{SpeechInfoObject|SpeechInfoObject array}},
        }
      }
    },
    "shouldEndSession": {{boolean}},
  },
  "sessionAttributes": {{object}},
  "version": {{string}}
}
*/

type CEKResponse struct {
	Card             struct{}   `json:"card"`
	Directives       []struct{} `json:"directives"`
	OutputSpeech     struct{}   `json:"outputSpeech"`
	ShouldEndSession bool       `json:"shouldEndSession"`
}

type CEKMessageResponse struct {
	Response          CEKResponse `json:"response"`
	SessionAttributes struct{}    `json:"sessionAttributes"`
	Version           string      `json:"version"`
}

func MakeMessageResponse() (CEKMessageResponse, error) {
	response := CEKMessageResponse{
		Response: CEKResponse{
			Directives:       []struct{}{},
			ShouldEndSession: true,
		},
		Version: "0.1.0",
	}
	return response, nil
}
