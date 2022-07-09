package tests

import (
	"testing"
)

func TestSendSubscribeMsg(t *testing.T) {
	ins := getIns()

	uri := "/cgi-bin/message/subscribe/bizsend"

	params := map[string]interface{}{
		"touser":      "xxx",
		"template_id": "xxx-xxx",
		"page":        "https://xxx.com",
		"data": map[string]interface{}{
			"thing5": map[string]interface{}{
				"value": "xxx",
			},
			"thing7": map[string]interface{}{
				"value": "content",
			},
			"time2": map[string]interface{}{
				"value": "2021-12-02 13:11",
			},
		},
		"miniprogram": map[string]interface{}{
			"appid":    "xxx",
			"pagepath": "pages/index/index?scene=123",
		},
	}

	res, err := ins.PostWithToken(uri, params)

	t.Error(res.GetInt("errcode"), err)
}
