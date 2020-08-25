package examples

import (
	"fmt"

	"github.com/idoubi/cutesdk/wxwork"
)

func ExampleBotSend() {
	bot := &wxwork.Bot{
		WebhookURL: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=8190e3ac-fde3-424e-8d88-4ffca9be7597",
	}

	res, err := bot.Text("hello bot").Send()
	fmt.Println(string(res), err)
	// Output: 1
}
