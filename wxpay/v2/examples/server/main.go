package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/cutesdk/cutesdk-go/wxpay/v2"
)

var (
	mchId  = "1498014222"
	apiKey = "WgGKrpgQwO7CuFLhFm8hInp79NzgOg4q"
)

type Mux struct {
}

func (m *Mux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	route := req.URL.Path

	log.Printf("route: %s\n", route)

	if strings.HasPrefix(route, "/pay-notify/") {
		mchId := strings.Replace(route, "/pay-notify/", "", 1)
		log.Printf("pay notify mchid: %s\n", mchId)

		PayNotifyHandler(mchId, resp, req)
		return
	}

	if strings.HasPrefix(route, "/refund-notify/") {
		mchId := strings.Replace(route, "/refund-notify/", "", 1)
		log.Printf("refund notify mchid: %s\n", mchId)

		RefundNotifyHandler(mchId, resp, req)
		return
	}

	http.NotFound(resp, req)
}

func main() {
	mux := new(Mux)

	err := http.ListenAndServe(":8084", mux)
	log.Printf("server listen: %v", err)
}

func PayNotifyHandler(mchId string, resp http.ResponseWriter, req *http.Request) {
	svr := getServer()

	err := svr.Listen(req, resp, func(msg *wxpay.NotifyMsg) *wxpay.ReplyMsg {
		log.Printf("pay notify order: %s\n", msg)

		outTradeNo := msg.GetString("out_trade_no")
		log.Printf("order out_trade_no: %s\n", outTradeNo)

		return msg.ReplySuccess()
	})

	log.Printf("pay notify listen error: %v", err)
}

func RefundNotifyHandler(mchId string, resp http.ResponseWriter, req *http.Request) {
	svr := getServer()

	err := svr.Listen(req, resp, func(msg *wxpay.NotifyMsg) *wxpay.ReplyMsg {
		log.Printf("refund notify order: %s\n", msg)

		outRefundNo := msg.GetString("out_refund_no")
		log.Printf("order out_refund_no: %s\n", outRefundNo)

		return msg.ReplySuccess()
	})

	log.Printf("refund notify listen error: %v", err)
}

func getServer() *wxpay.Server {
	svr, err := wxpay.NewServer(&wxpay.Options{
		MchId:  mchId,
		ApiKey: apiKey,
	})

	if err != nil {
		log.Fatalf("new wxpay server failed: %v\n", err)
	}

	return svr
}
