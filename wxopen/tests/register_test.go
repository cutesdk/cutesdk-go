package tests

import (
	"fmt"
	"testing"
)

func TestRegisterCompanyWxapp(t *testing.T) {
	ins := getIns()

	componentAccessToken, err := ins.GetComponentAccessToken()
	if err != nil {
		t.Fatalf("get component_access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/component/fastregisterweapp?action=create&component_access_token=%s", componentAccessToken)

	params := map[string]interface{}{
		"name":                 "xxx",
		"code":                 "123",
		"code_type":            1,
		"legal_persona_wechat": "123",
		"legal_persona_name":   "xxx",
		"component_phone":      "1234567",
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetInt("errcode"), err)
}

func TestRegisterPesonalWxapp(t *testing.T) {
	ins := getIns()

	componentAccessToken, err := ins.GetComponentAccessToken()
	if err != nil {
		t.Fatalf("get component_access_token failed: %v", err)
	}

	name := "xxx"
	wechat := "xxx"
	phone := "xxx"

	uri := fmt.Sprintf("/wxa/component/fastregisterpersonalweapp?action=create&component_access_token=%s", componentAccessToken)

	params := map[string]interface{}{
		"idname":          name,
		"wxuser":          wechat,
		"component_phone": phone,
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetInt("errcode"), err)
}
