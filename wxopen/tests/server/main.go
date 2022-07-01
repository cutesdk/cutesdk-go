package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cutesdk/cutesdk-go/wxopen"
)

type Mux struct {
}

func (m *Mux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	route := req.URL.Path

	fmt.Printf("route: %s\n", route)

	if route == "/event-notify" {
		EventNotifyHandler(resp, req)
		return
	}

	if strings.HasPrefix(route, "/msg-notify/") {
		appid := strings.Replace(route, "/msg-notify/", "", 1)
		fmt.Printf("msg appid: %s\n", appid)

		MsgNotifyHandler(appid, resp, req)
		return
	}

	http.NotFound(resp, req)
}

func main() {
	mux := new(Mux)

	err := http.ListenAndServe(":8091", mux)
	fmt.Printf("server listen: %v", err)
}

func EventNotifyHandler(resp http.ResponseWriter, req *http.Request) {
	ins := getIns()

	err := ins.Listen(req, resp, func(msg *wxopen.NotifyMsg) *wxopen.ReplyMsg {
		infoType := msg.GetString("InfoType")
		fmt.Printf("notify info type: %s\n", infoType)

		if infoType == "component_verify_ticket" {
			ticket := msg.GetString("ComponentVerifyTicket")
			fmt.Printf("ticket: %s\n", ticket)
			ins.SetComponentVerifyTicket(ticket, 2*time.Hour)

			return nil
		}

		fmt.Printf("notify info: %s\n", msg)

		return nil
	})
	fmt.Printf("event notify listen error: %v", err)
}

func MsgNotifyHandler(appid string, resp http.ResponseWriter, req *http.Request) {
	ins := getIns()

	err := ins.Listen(req, resp, func(msg *wxopen.NotifyMsg) *wxopen.ReplyMsg {
		appid := msg.ReceiveId()
		fmt.Printf("notify receive appid: %s\n", appid)

		msgType := msg.GetString("MsgType")
		fmt.Printf("notify msg type: %s\n", msgType)

		fmt.Printf("notify msg: %s\n", msg)

		switch msgType {
		case "image":
			return msg.ReplyImage(msg.GetString("MediaId"))
		case "voice":
			return msg.ReplyVoice(msg.GetString("MediaId"))
		case "video":
			return msg.ReplyVideo(msg.GetString("MediaId"), "video title", "about this video")
		case "link":
			return msg.ReplyNews(msg.GetString("Title"), msg.GetString("Description"), msg.GetString("Url"), "")
		}

		return msg.ReplyText(msg.String())
	})
	fmt.Printf("msg notify listen error: %v", err)
}

func getIns() *wxopen.Instance {
	opts := &wxopen.Options{
		ComponentAppid:     "xxx",
		ComponentAppsecret: "xxx",
		VerifyToken:        "xxx",
		EncodingAesKey:     "xxx",
		Debug:              true,
	}

	ins, err := wxopen.New(opts)

	if err != nil {
		panic(err)
	}

	return ins
}
