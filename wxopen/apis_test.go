package wxopen

import "testing"

func TestPushTicket(t *testing.T) {
	sdk := getWxOpen()

	res, err := sdk.PushTicket()

	if err != nil {
		t.Fatalf("request api error: %v\n", err)
	}

	if res.Parsed().Get("errcode").Int() != 0 {
		t.Fatalf("request api failed: %s", res)
	}
}
