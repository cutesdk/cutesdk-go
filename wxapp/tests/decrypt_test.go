package tests

import (
	"testing"
)

func TestDecryptUserInfo(t *testing.T) {
	ins := getIns()

	sessionKey := "xxx+xxx=="
	encryptedData := "xxx+igb4QgbTA+4te1+uEl/xxx=="
	iv := "xxx=="

	res, err := ins.DecryptUserInfo(sessionKey, encryptedData, iv)

	if err != nil {
		t.Fatalf("decrypt user info error: %v\n", err)
	}

	t.Error(res.GetString("nickName"), err)
}

func TestDecryptPhone(t *testing.T) {
	ins := getIns()

	sessionKey := "xxx+xxx=="
	encryptedData := "xxx=="
	iv := "xxx=="

	res, err := ins.DecryptPhone(sessionKey, encryptedData, iv)

	if err != nil {
		t.Fatalf("decrypt phone error: %v\n", err)
	}

	t.Error(res.GetString("phoneNumber"), err)
}
