package lgtv

import (
	"crypto/tls"
	"github.com/gorilla/websocket"
	webos "github.com/kaperys/go-webos"
	"log"
	"net"
	"time"
)

const tvIp string = "<redacted>"
const clientKey string = "<redacted>"

func createTV() *webos.TV {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NetDial: (&net.Dialer{
			Timeout: time.Second * 5,
		}).Dial,
	}

	tv, err := webos.NewTV(&dialer, tvIp)
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	go tv.MessageHandler()

	if err = tv.AuthoriseClientKey(clientKey); err != nil {
		log.Fatalf("could not authorise using client key: %v", err)
	}
	return tv
}

func LaunchApp(appId string) error {
	tv := createTV()
	defer tv.Close()
	return tv.LaunchApp(appId)
}

func CloseApp(appId string) error {
	tv := createTV()
	defer tv.Close()
	return tv.CloseApp(appId)
}
