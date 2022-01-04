package wxopen

import "testing"

func TestGetComponentAccessToken(t *testing.T) {
	w := getWxOpen()

	res, err := w.GetComponentAccessToken()
	if err != nil {
		t.Fatalf("get component access_token error: %v\n", err)
	}

	t.Errorf("%s", res)
}
