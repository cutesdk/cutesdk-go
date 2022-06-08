package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cutesdk/cutesdk-go/wxopen"
)

func main() {
	http.HandleFunc("/notify", NotifyHandler)
	http.ListenAndServe(":8091", nil)
}

func NotifyHandler(resp http.ResponseWriter, req *http.Request) {
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

		fmt.Printf("msg: %s\n", msg)

		return nil
	})
	fmt.Printf("notify listen error: %v", err)
}

func getIns() *wxopen.Instance {
	opts := &wxopen.Options{
		ComponentAppid:     "wxf2f955ce09390e6a",
		ComponentAppsecret: "d6e9032e5f5bcea2f96b66f2c4e1cab8",
		VerifyToken:        "OaNCoqFftJz7YkUD",
		EncodingAesKey:     "MNfsPhrt28W4dksbARCANqIHqLmzdbZvQH8WtGgGzHv",
		Debug:              true,
	}

	ins, err := wxopen.New(opts)

	if err != nil {
		panic(err)
	}

	return ins
}
