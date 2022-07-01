package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cutesdk/cutesdk-go/wxmp"
)

type Mux struct {
}

func (m *Mux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	route := req.URL.Path

	fmt.Printf("route: %s\n", route)

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

	err := http.ListenAndServe(":8092", mux)
	fmt.Printf("server listen: %v", err)
}

func MsgNotifyHandler(appid string, resp http.ResponseWriter, req *http.Request) {
	ins := getIns()

	err := ins.Listen(req, resp, func(msg *wxmp.NotifyMsg) *wxmp.ReplyMsg {
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

		return msg.ReplyText("wxmp" + msg.String())
	})
	fmt.Printf("msg notify listen error: %v", err)
}

func getIns() *wxmp.Instance {
	opts := &wxmp.Options{
		Appid:          "xxx",
		Secret:         "xxx",
		VerifyToken:    "xxx",
		EncodingAesKey: "xxx",
		Debug:          true,
	}

	ins, err := wxmp.New(opts)

	if err != nil {
		panic(err)
	}

	return ins
}
