package examples

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func ExampleCreateQrcode() {
	cli := getClient()

	uri := "/cgi-bin/qrcode/create"
	params := map[string]interface{}{
		"expire_seconds": 300,
		"action_name":    "QR_STR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": "id=3",
			},
		},
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") > 0 {
		log.Fatalf("create qrcode failed: %s\n", res)
	}

	ticket := res.GetString("ticket")
	picurl := fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s", ticket)
	log.Printf("picurl: %s\n", picurl)

	ret, err := http.Get(picurl)
	if err != nil {
		log.Fatalf("fetch qrcode failed: %v\n", err)
	}
	defer ret.Body.Close()
	body, _ := io.ReadAll(ret.Body)

	filename := "./qrcode.jpg"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("create qrcode file failed: %v\n", err)
	}
	defer f.Close()

	if _, err := f.Write(body); err != nil {
		log.Fatalf("save qrcode data failed: %v\n", err)
	}

	fmt.Println("ok")
	// Output: ok
}
