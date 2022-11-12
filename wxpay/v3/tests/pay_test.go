package tests

import (
	"fmt"
	"testing"

	"github.com/idoubi/goutils"
)

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

func TestQueryOrder(t *testing.T) {
	ins := getPayIns()

	outTradeNo := "1578368941695176704"

	uri := fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s?mchid=%s", outTradeNo, ins.GetOptions().MchId)

	res, err := ins.Get(uri)

	t.Error(res, err)
}
