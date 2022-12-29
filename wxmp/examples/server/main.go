package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/cutesdk/cutesdk-go/wxmp"
)

var (
	verifyToken    = "xxx"
	encodingAesKey = "xxx"
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
	server := getServer()

	err := server.Listen(req, resp, func(msg *wxmp.NotifyMsg) *wxmp.ReplyMsg {
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
	server, err := wxmp.NewServer(&wxmp.Options{
		VerifyToken:    verifyToken,
		EncodingAesKey: encodingAesKey,
		Debug:          true,
	})

	if err != nil {
		log.Fatalf("new wxmp server failed: %v\n", err)
	}

	return server
}
