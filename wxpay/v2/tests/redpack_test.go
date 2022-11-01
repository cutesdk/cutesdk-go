package tests

import (
	"testing"

	"github.com/idoubi/goutils"
)

func TestRedpack(t *testing.T) {
	ins := getPayIns()

	uri := "/mmpaymkttransfers/sendredpack"

	params := map[string]interface{}{
		"mch_billno":   goutils.GenSnowid(),
		"wxappid":      ins.GetOptions().Appid,
		"send_name":    "心跳实验室",
		"re_openid":    "orNyi07T5GeE2nwVV7b1dJ3xGnPM",
		"total_amount": 20000,
		"total_num":    1,
		"wishing":      "快来参加大冒险软件促销活动吧~",
		"client_ip":    "127.0.0.1",
		"act_name":     "大冒险软件促销",
		"remark":       "大冒险软件促销活动",
		"scene_id":     "PRODUCT_1",
	}

	res, err := ins.PostWithCert(uri, params)

	t.Error(res.GetString("trade_state"), err)
}
