package wxapp

import (
	"fmt"

	"github.com/idoubi/goz"
)

// SendSubmsg 发送订阅通知
func (w *Wxapp) SendSubmsg(params map[string]interface{}) (Result, error) {
	apiURL := fmt.Sprintf(apiBase+"/cgi-bin/message/subscribe/send?access_token=%s", "")
	resp, err := goz.Post(apiURL, goz.Options{
		Debug: w.opts.Debug,
		JSON:  params,
	})

	if err != nil {
		return nil, err
	}

	body, err := resp.GetBody()

	return Result(body), err
}
