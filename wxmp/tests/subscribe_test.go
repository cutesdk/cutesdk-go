package tests

import "testing"

func TestSendSubscribeMsg(t *testing.T) {
	client := getClient()

	tplId := "LTA5XpzSBWQQ3YI0hOw9Ucj-U1bvZIB_Ut9kHAhNBEk"
	openid := "orNyi07T5GeE2nwVV7b1dJ3xGnPM"
	page := "https://idoubi.cc"

	data := map[string]interface{}{
		"thing5": map[string]interface{}{
			"value": "idoubi",
		},
		"thing7": map[string]interface{}{
			"value": "content",
		},
		"time2": map[string]interface{}{
			"value": "2021-12-02 13:11",
		},
	}

	miniprogram := map[string]interface{}{
		"appid":    "wx25da2eca8fa3f4ef",
		"pagepath": "pages/index/index?scene=123",
	}

	res, err := client.SendSubscribeMsg(tplId, openid, data, page, miniprogram)

	t.Error(res, err)
}
