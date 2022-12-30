package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/wxopen"
)

var (
	appid          = "wxd19865e64822ff7a"
	secret         = "d347fdaa99f879099264ad0c072e0181"
	verifyToken    = "ga55qnibf14sj3bs55xwzekzmuhp52va"
	encodingAesKey = "lq94ZOugkS4YHsy4REbC4irHTUjX17TmWpbP55gO40j"
)

type Mux struct {
}

func (m *Mux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	route := req.URL.Path

	log.Printf("route: %s\n", route)

	if route == "/event-notify" {
		EventNotifyHandler(resp, req)
		return
	}

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

	err := http.ListenAndServe(":8083", mux)
	log.Printf("server listen: %v", err)
}

func EventNotifyHandler(resp http.ResponseWriter, req *http.Request) {
	svr := getServer()

	err := svr.Listen(req, resp, func(msg *wxopen.NotifyMsg) *wxopen.ReplyMsg {
		log.Printf("notify msg: %s\n", msg)

		return nil
	})
	log.Printf("event notify listen error: %v", err)
}

func MsgNotifyHandler(appid string, resp http.ResponseWriter, req *http.Request) {
	svr := getServer()

	err := svr.Listen(req, resp, func(msg *wxopen.NotifyMsg) *wxopen.ReplyMsg {
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

func getServer() *wxopen.Server {
	opts := &wxopen.Options{
		Appid:          appid,
		Secret:         secret,
		VerifyToken:    verifyToken,
		EncodingAesKey: encodingAesKey,
		Debug:          true,
		Cache: &cache.FileOptions{
			Dir: "../cache",
		},
	}

	svr, err := wxopen.NewServer(opts)

	if err != nil {
		log.Fatalf("new wxopen server failed: %v\n", err)
	}

	return svr
}
