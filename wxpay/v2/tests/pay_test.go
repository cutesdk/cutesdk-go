package tests

import (
	"testing"

	"github.com/idoubi/goutils"
)

func TestUnifiedOrder(t *testing.T) {
	client := getClient()

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

	res, err := client.UnifiedOrder(params)
	x := res.XmlParsed()
	t.Error(x.Get("prepay_id").String(), err)
}

func TestGetPayParams(t *testing.T) {
	client := getClient()

	params := map[string]interface{}{
		"prepay_id": "wx181643135445953d9d84c82e51f2160000",
		"sign_type": "HMAC-SHA256",
	}

	res, err := client.GetPayParams(params)

	t.Error(res, err)
}
