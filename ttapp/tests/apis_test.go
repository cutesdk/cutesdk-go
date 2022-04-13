package tests

import (
	"os"
	"testing"
)

func TestCode2Session(t *testing.T) {
	client := getClient()

	code := "pox7TA-_rr3DrPLDaeuyWRGBnJsn9vWTxaLa6tPYJVMiBxaZ_cc47wGLVDlJyCsklQw5E3eRsqthmLVRzco_rBodQvGIUeujOgCjKcIkf-7HiqyPritSmlleOOA"
	anonymousCode := "KTqKa6x31on0X6bf7E3SlwU4MSF3BNizc64zsNBd6KGDfkkSvCjjysmCad04rfUqLLcUPUfMGyQvpkeGW1Jjt1Q6Yrjp5OCX7rkfKDtOduqf2Y_1UQmlIx4YMdA"

	res, err := client.Code2Session(code, anonymousCode)

	if err != nil {
		t.Fatalf("request code2session error: %v\n", err)
	}

	t.Error(res, err)
}

func TestCreateQrcode(t *testing.T) {
	client := getClient()

	data := map[string]interface{}{
		"appname":    "douyin",
		"path":       "",
		"width":      430,
		"set_icon":   true,
		"line_color": map[string]int{"r": 36, "g": 8, "b": 253},
		"background": map[string]int{"r": 253, "g": 43, "b": 85},
	}

	res, err := client.CreateQrcode(data)

	if err != nil {
		t.Fatalf("request api error: %v\n", err)
	}

	f, err := os.Create("./qrcode.jpeg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(res)
}

func TestSendSubscribeMsg(t *testing.T) {
	client := getClient()

	tplId := "MSG5877801d55d01588355646371eaaca417d88f15682"
	openid := "L1p8S7Bwp2kweSGr"
	data := map[string]string{
		"打卡名称": "今日打卡提醒",
		"备注":   "快去打卡吧~",
	}
	page := ""

	res, err := client.SendSubscribeMsg(tplId, openid, data, page)

	t.Error(res, err)
}
