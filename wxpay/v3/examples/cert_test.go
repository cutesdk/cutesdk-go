package examples

import (
	"fmt"
	"log"
)

func ExampleGetCertificates() {
	cli := getClient()

	uri := "/v3/certificates"

	res, err := cli.Get(uri)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	for _, v := range res.GetArray("data") {
		serialNo := v.Get("serial_no").String()
		expireTime := v.Get("expire_time").String()

		encryptedData := v.Get("encrypt_certificate")
		cipertext := encryptedData.Get("ciphertext").String()
		nonce := encryptedData.Get("nonce").String()
		associatedData := encryptedData.Get("associated_data").String()
		res, err := cli.DecryptData(cipertext, nonce, associatedData)
		if err != nil {
			log.Fatalf("decrypt data failed: %v\n", err)
		}

		fmt.Printf("%s, %s, \n%s", serialNo, expireTime, res)
	}
	// Output: xxx
}
