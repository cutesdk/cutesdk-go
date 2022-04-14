package tests

import (
	"testing"
)

func TestSendServiceMsg(t *testing.T) {
	client := getClient()

	openid := "orNyi07T5GeE2nwVV7b1dJ3xGnPM"
	msgtype := "text"

	params := map[string]interface{}{
		"content": "hello world",
	}

	res, err := client.SendServiceMsg(openid, msgtype, params)

	t.Error(res, err)
}
