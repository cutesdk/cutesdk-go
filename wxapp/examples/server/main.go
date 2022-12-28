package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/wxapp"
)

var (
	appid          = "xxx"
	secret         = "xxx"
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

	err := http.ListenAndServe(":8081", mux)
	log.Printf("server listen: %v", err)
}

func MsgNotifyHandler(appid string, resp http.ResponseWriter, req *http.Request) {
	server := getServer()
	client := getClient()

	err := server.Listen(req, resp, func(msg *wxapp.NotifyMsg) *wxapp.ReplyMsg {
		log.Printf("notify msg: %s\n", msg)

		msgType := msg.GetString("MsgType")

		if msgType == "text" {
			content := msg.GetString("Content")
			if content == "人工" {
				return msg.Transfer()
			}
		}

		openid := msg.GetString("FromUserName")
		content := msg.String()

		go sendTextMsg(client, openid, content)

		return nil
	})

	log.Printf("msg notify listen error: %v", err)
}

func sendTextMsg(client *wxapp.Client, openid string, content string) error {
	uri := "/cgi-bin/message/custom/send"
	params := map[string]interface{}{
		"touser":  openid,
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": content,
		},
	}

	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Printf("request api failed: %v\n", err)
		return err
	}

	if res.GetInt("errcode") != 0 {
		log.Printf("send service text msg failed: %s\n", res)
		return err
	}

	return nil
}

func getServer() *wxapp.Server {
	server, err := wxapp.NewServer(&wxapp.Options{
		VerifyToken:    verifyToken,
		EncodingAesKey: encodingAesKey,
		Debug:          true,
	})

	if err != nil {
		log.Fatalf("new wxapp server failed: %v\n", err)
	}

	return server
}

func getClient() *wxapp.Client {
	client, err := wxapp.NewClient(&wxapp.Options{
		Appid:  appid,
		Secret: secret,
		Debug:  true,
		Cache: &cache.FileOptions{
			Dir: "../cache",
		},
	})

	if err != nil {
		log.Fatalf("new wxapp client failed: %v\n", err)
	}

	return client
}
