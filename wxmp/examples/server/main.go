package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/wxmp"
)

var (
	appid  = "xxx"
	secret = "xxx"
	token  = "xxx"
	aesKey = "xxx"
)

type Mux struct {
}

func (m *Mux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	route := req.URL.Path

	log.Printf("route: %s\n", route)

	if strings.HasPrefix(route, "/msg-notify/") {
		appid := strings.Replace(route, "/msg-notify/", "", 1)
		log.Printf("msg appid: %s\n", appid)

		MsgNotifyHandler(appid, resp, req)
		return
	}

	http.NotFound(resp, req)
}

func main() {
	mux := new(Mux)

	err := http.ListenAndServe(":8082", mux)
	log.Printf("server listen: %v", err)
}

func MsgNotifyHandler(appid string, resp http.ResponseWriter, req *http.Request) {
	svr := getServer()

	err := svr.Listen(req, resp, func(msg *wxmp.NotifyMsg) *wxmp.ReplyMsg {
		log.Printf("notify msg: %s\n", msg)

		msgType := msg.GetString("MsgType")

		switch msgType {
		case "image":
			return msg.ReplyImage(msg.GetString("MediaId"))
		case "voice":
			return msg.ReplyVoice(msg.GetString("MediaId"))
		case "link":
			return msg.ReplyNews(msg.GetString("Title"), msg.GetString("Description"), msg.GetString("Url"), "")
		}

		return msg.ReplyText(msg.String())
	})

	log.Printf("msg notify listen error: %v", err)
}

func getServer() *wxmp.Server {
	svr, err := wxmp.NewServer(&wxmp.Options{
		Appid:  appid,
		Secret: secret,
		Token:  token,
		AesKey: aesKey,
		Debug:  true,
		Cache: &cache.FileOptions{
			Dir: "../cache",
		},
	})

	if err != nil {
		log.Fatalf("new wxmp server failed: %v\n", err)
	}

	return svr
}
