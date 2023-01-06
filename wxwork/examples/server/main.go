package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/wxwork"
)

var (
	corpid         = "xxx"
	contactSecret  = "xxx"
	customerSecret = "xxx"
	serviceSecret  = "xxx"
	agentId        = "xxx"
	agentSecret    = "xxx"
	token          = "xxx"
	aesKey         = "xxx"
)

type Mux struct {
}

func (m *Mux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	route := req.URL.Path

	log.Printf("route: %s\n", route)

	if strings.HasPrefix(route, "/event-notify/") {
		appid := strings.Replace(route, "/event-notify/", "", 1)
		log.Printf("msg appid: %s\n", appid)

		var secret string
		if appid == "contact" {
			secret = contactSecret
		}
		if appid == "customer" {
			secret = customerSecret
		}
		if appid == "service" {
			secret = serviceSecret
		}
		if appid == agentId {
			secret = agentSecret
		}

		EventNotifyHandler(appid, secret, resp, req)
		return
	}

	http.NotFound(resp, req)
}

func main() {
	mux := new(Mux)

	err := http.ListenAndServe(":8086", mux)
	log.Printf("server listen: %v", err)
}

func EventNotifyHandler(appid, secret string, resp http.ResponseWriter, req *http.Request) {
	svr := getServer(appid, secret)

	err := svr.Listen(req, resp, func(msg *wxwork.NotifyMsg) *wxwork.ReplyMsg {
		log.Printf("%s notify msg: %s\n", appid, msg)

		msgType := msg.GetString("MsgType")
		event := msg.GetString("Event")
		if !(msgType == "event" && event == "LOCATION") {
			return msg.ReplyText(msg.String())
		}

		return nil
	})

	log.Printf("event notify listen error: %v", err)
}

func getServer(appid, secret string) *wxwork.Server {
	svr, err := wxwork.NewServer(&wxwork.Options{
		Corpid: corpid,
		Appid:  appid,
		Secret: secret,
		Token:  token,
		AesKey: aesKey,
		Debug:  true,
		Cache: &cache.FileOptions{
			Dir: "../../cache",
		},
	})

	if err != nil {
		log.Fatalf("new wxwork server failed: %v\n", err)
	}

	return svr
}
