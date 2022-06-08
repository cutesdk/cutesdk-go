package tests

import (
	"testing"
)

func TestDecryptUserInfo(t *testing.T) {
	ins := getIns()

	sessionKey := "OfPrAxDK5X0eUBKbL9u4fA=="
	encryptedData := "ImvujJDNowQ8oy5YLMKhdBwtRlkIuSm+8/qNvVIYwCtmt66mRlGFWznp6sfe0fIJSgBek/0Ur2Wk5WSQhmjGzi3M6hD/zZ4VtyZklA8Kj6VwJLmxJgktc/a6+gBz++dCB8BOUdhPwm0ZSXnBQTRw3RquCToxZm9HPa1S3phZdSeJTNwBqwPyr3iEimWNq0uN9L0zalWXpJbwNxyonbBw0K5VOE1wJr38f//Hp2rFxFyWZDvfDkWcwioeAcDI3HVUeR08UEUH44ifCInDAPijVKB9CAPOfJSckhyzv9PwuVAL1aKGuWY5htpVT3rnASqw3zxdaMIjdfsSZy0qpICV5I+9XyE59nrnbuMp8FL5yi3mKnMSzZJd8tTsen2A6UXuSRAXsdtf5GnRuF4AYChEy9eqWabeYIz0y/bOyX0zNduQyiVxU9rFS7JPzb1Yu684"
	iv := "pi0pj6W/tSP+1UUxFSDuHA=="

	res, err := ins.DecryptUserInfo(sessionKey, encryptedData, iv)

	if err != nil {
		t.Fatalf("decrypt user info error: %v\n", err)
	}

	t.Error(res.GetString("nickName"), err)
}
