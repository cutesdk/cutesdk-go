package tests

import (
	"fmt"
	"testing"
)

func TestPushTicket(t *testing.T) {
	ins := getIns()

	uri := "/cgi-bin/component/api_start_push_ticket"

	params := map[string]interface{}{
		"component_appid":  ins.GetComponentAppid(),
		"component_secret": ins.GetComponentAppsecret(),
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetInt("errcode"), err)
}

func TestCreatePreauthCode(t *testing.T) {
	ins := getIns()

	componentAccessToken, err := ins.GetComponentAccessToken()
	if err != nil {
		t.Fatalf("get component_access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/component/api_create_preauthcode?component_access_token=%s", componentAccessToken)

	params := map[string]interface{}{
		"component_appid": ins.GetComponentAppid(),
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetString("pre_auth_code"), err)
}
