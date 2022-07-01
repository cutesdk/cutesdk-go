package tests

import (
	"testing"
)

func TestDecryptUserInfo(t *testing.T) {
	ins := getIns()

	sessionKey := "xxx=="
	encryptedData := "xxx"
	iv := "xxx/xxx+xxx=="

	res, err := ins.DecryptUserInfo(sessionKey, encryptedData, iv)

	if err != nil {
		t.Fatalf("decrypt user info failed: %v", err)
	}

	t.Error(res.GetString("nickName"), err)
}
