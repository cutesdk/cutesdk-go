package examples

import (
	"fmt"
	"log"
)

func ExampleDecryptUserInfo() {
	cli := getClient()

	sessionKey := "xxx"
	encryptedData := "xxx"
	iv := "xxx"

	res, err := cli.DecryptUserInfo(sessionKey, encryptedData, iv)
	if err != nil {
		log.Fatalf("decrypt user info failed: %v\n", err)
	}

	if res.GetString("nickName") == "" {
		log.Fatalf("decrypt user info failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
