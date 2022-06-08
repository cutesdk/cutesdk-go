package tests

import (
	"testing"

	"github.com/cutesdk/cutesdk-go/wxpay/v2"
	"github.com/idoubi/goutils"
)

func getPayIns() *wxpay.Instance {
	opts := &wxpay.Options{
		MchId:  "1498014222",
		Appid:  "wx25da2eca8fa3f4ef",
		ApiKey: "Q5xQzBMvdvKQgn3Li3e26XVb4TNYuS13",
		Debug:  true,
	}

	ins, _ := wxpay.New(opts)

	return ins
}

func TestUnifiedOrder(t *testing.T) {
	ins := getPayIns()

	uri := "/pay/unifiedorder"

	params := map[string]interface{}{
		"body":             "支付测试",
		"attach":           "qid=3&aid=6",
		"out_trade_no":     goutils.GenOrderID(),
		"total_fee":        3,
		"spbill_create_ip": "127.0.0.1",
		"notify_url":       "https://demo.thinkwx.com/wxpay/callback",
		"trade_type":       "JSAPI",
		"openid":           "oLW495c2KVrduEpiSGDpHp7qKqCc",
		"sign_type":        "HMAC-SHA256",
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetString("prepay_id"), err)
}

func TestGetPayParams(t *testing.T) {
	ins := getPayIns()

	params := map[string]interface{}{
		"prepay_id": "wx181643135445953d9d84c82e51f2160000",
		"sign_type": "HMAC-SHA256",
	}

	res, err := ins.GetPayParams(params)

	t.Error(res, err)
}
