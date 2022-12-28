package examples

import (
	"fmt"
	"log"
)

func ExampleCode2Session() {
	client := getClient()

	code := "Q2d-wAQcpdMPvjX_r_5G6FUr-5vLyu-fnZQwZA2X66Llhd7Sbi7MBhIX7fOiCziPMIZtS8I7LmuTCqLEUPl8yPPG_GkCnZ_SNxB5nE3lEyw_kIEBjnlVLzVGzrc"
	// anonymousCode := "hD8F_FnsLEitJ392rtaR27iyPhqf4Hb3izWLiBg7IdhIE4PUxgIvdWxWtYdOJs8qpT1GFp-Xajf3JAIvKyGELtuTxPBmP4w4aloBH5er6d8MSdJNTGf7-zac6jc"

	uri := "/api/apps/v2/jscode2session"
	params := map[string]interface{}{
		"appid":  client.GetAppid(),
		"secret": client.GetSecret(),
		"code":   code,
		// "anonymous_code": anonymousCode,
	}

	res, err := client.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	openid := res.GetString("data.openid")
	if openid == "" {
		log.Fatalf("code2session failed: %s\n", res)
	}

	fmt.Println(openid)
	// Output: xxx
}
