package examples

import (
	"fmt"
	"log"
)

func ExampleDecryptUserInfo() {
	client := getClient()

	sessionKey := "xxx"
	encryptedData := "xxx"
	iv := "xxx"

	res, err := client.DecryptUserInfo(sessionKey, encryptedData, iv)
	if err != nil {
		log.Fatalf("decrypt user info failed: %v\n", err)
	}

	if res.GetString("nickName") == "" {
		log.Fatalf("decrypt user info failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleDecryptPhone() {
	client := getClient()

	sessionKey := "xxx"
	encryptedData := "xxx"
	iv := "xxx"

	res, err := client.DecryptPhone(sessionKey, encryptedData, iv)
	if err != nil {
		log.Fatalf("decrypt phone failed: %v\n", err)
	}

	if res.GetString("phoneNumber") == "" {
		log.Fatalf("decrypt phone failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
