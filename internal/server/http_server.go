package server

import (
	"encoding/json"
	"lgtv_ctl_clova_ext/internal/lgtv"
	"lgtv_ctl_clova_ext/internal/protocol"
	"log"
	"net/http"
)

var appMappings = map[string]string{
	"넷플릭스": "netflix",
	"유튜브": "youtube.leanback.v4",
	"티빙": "cj.eandm",
	"왓챠": "com.frograms.watchaplay.webos",
	"구글 플레이": "googleplaymovieswebos",
}

func indexHandler(w http.ResponseWriter, req *http.Request) {

	var cekMessage protocol.CEKMessage
	if err := json.NewDecoder(req.Body).Decode(&cekMessage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if cekMessage.Request.Type != "IntentRequest" {
		http.Error(w, "Unsupported intent type", http.StatusBadRequest)
		return
	}

	var appName, hasAppName = cekMessage.Request.Intent.Slots["app_name"]
	if !hasAppName {
		http.Error(w, "Unsupported intent type", http.StatusBadRequest)
		return
	}

	var value, hasValue = appName["value"]
	if !hasValue {
		http.Error(w, "Unsupported intent type", http.StatusBadRequest)
		return
	}

	switch cekMessage.Request.Intent.Name {
	case "LaunchApplication":
		log.Println("Launch:", value)
		appId, ok := appMappings[value]
		if !ok {
			http.Error(w, "Unknown app name", http.StatusBadRequest)
			return
		}
		lgtv.LaunchApp(appId)
	case "CloseApplication":
		log.Println("Close:", value)
		appId, ok := appMappings[value]
		if !ok {
			http.Error(w, "Unknown app name", http.StatusBadRequest)
			return
		}
		lgtv.CloseApp(appId)
	default:
		http.Error(w, "Unsupported intent name", http.StatusBadRequest)
		return
	}

	if resp, err := protocol.MakeMessageResponse(); err == nil {
		jsonResp, _ := json.Marshal(&resp)
		w.Write(jsonResp)
	}
}

func Serve() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
