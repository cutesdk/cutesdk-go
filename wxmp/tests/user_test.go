package tests

import "testing"

func TestGetUserList(t *testing.T) {
	client := getClient()

	nextOpenid := ""

	res, err := client.GetUserList(nextOpenid)

	t.Error(res, err)
}

func TestGetUserInfo(t *testing.T) {
	client := getClient()

	openid := "orNyi07T5GeE2nwVV7b1dJ3xGnPM"

	res, err := client.GetUserInfo(openid)

	t.Error(res, err)
}
