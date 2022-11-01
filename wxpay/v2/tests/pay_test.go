package tests

import (
	"encoding/json"
	"testing"

	"github.com/cutesdk/cutesdk-go/wxpay/v2"
	"github.com/idoubi/goutils"
)

func getPayIns() *wxpay.Instance {
	opts := &wxpay.Options{
		MchId:      "xxx",
		Appid:      "xxx",
		ApiKey:     "xxx",
		CertKey:    ``,
		PrivateKey: ``,
		Debug:      true,
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
		"out_trade_no":     goutils.GenSnowid(),
		"total_fee":        3,
		"spbill_create_ip": "127.0.0.1",
		"notify_url":       "https://xxx.com/wxpay/callback",
		"trade_type":       "JSAPI",
		"openid":           "xxx",
		"sign_type":        "HMAC-SHA256",
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetString("prepay_id"), err)
}

func TestGetPayParams(t *testing.T) {
	ins := getPayIns()

	params := map[string]interface{}{
		"prepay_id": "xxx",
		"sign_type": "HMAC-SHA256",
	}

	res, err := ins.GetPayParams(params)

	j, _ := json.Marshal(res)

	t.Error(string(j), err)
}
