package wxapp

import "testing"

func TestCode2Session(t *testing.T) {
	sdk := getSdk()

	code := "021zYc0w3kfkMX2nXw1w3NA1Bb3zYc0e"

	res, err := sdk.Code2Session(code)

	// {"session_key":"wgTAh4KDkqKqMO0Xs\/8jXw==","openid":"oLW495c2KVrduEpiSGDpHp7qKqCc","unionid":"oCLLZ5i-U9fI5FGtJulN5qDYn-Is"}
	if err != nil {
		t.Fatalf("request code2session error: %v\n", err)
	}

	_ = res
}
