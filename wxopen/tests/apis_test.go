package tests

import "testing"

func TestPushTicket(t *testing.T) {
	client := getClient()

	res, err := client.StartPushTicket()

	t.Error(res, err)
}

func TestFetchComponentAccessToken(t *testing.T) {
	client := getClient()

	res, err := client.FetchComponentAccessToken()

	t.Error(res, err)
}

func TestCreatePreauthCode(t *testing.T) {
	client := getClient()

	res, err := client.CreatePreauthCode()

	t.Error(res, err)
}

func TestRegisterCompanyWxapp(t *testing.T) {
	client := getClient()

	params := map[string]interface{}{
		"name":                 "tencent",
		"code":                 "123",
		"code_type":            1,
		"legal_persona_wechat": "123",
		"legal_persona_name":   "candy",
		"component_phone":      "1234567",
	}

	res, err := client.RegisterCompanyWxapp(params)

	t.Error(res, err)
}

func TestRegisterPesonalWxapp(t *testing.T) {
	client := getClient()

	name := "idoubi"
	wechat := "xxx"
	phone := "xxx"

	res, err := client.RegisterPersonalWxapp(name, wechat, phone)

	t.Error(res, err)
}
