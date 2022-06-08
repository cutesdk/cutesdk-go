package tests

import (
	"testing"
)

func TestSendSubscribeMsg(t *testing.T) {
	ins := getIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := "/api/apps/subscribe_notification/developer/v1/notify"

	tplId := "MSG5877801d55d01588355646371eaaca417d88f15682"
	openid := "L1p8S7Bwp2kweSGr"
	data := map[string]interface{}{
		"打卡名称": "今日打卡提醒",
		"备注":   "快去打卡吧~",
	}
	page := ""

	postData := map[string]interface{}{
		"access_token": accessToken,
		"app_id":       ins.GetAppid(),
		"tpl_id":       tplId,
		"open_id":      openid,
		"data":         data,
		"page":         page,
	}

	res, err := ins.Post(uri, postData)

	t.Error(res.GetInt("err_no"), err)
}
