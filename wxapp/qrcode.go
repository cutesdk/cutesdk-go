package wxapp

import (
	"errors"
	"fmt"
	"strings"

	"github.com/idoubi/goz"
)

// GetUnlimitCode 获取小程序码 B接口
func (w *WxApp) GetUnlimitCode(params map[string]interface{}) (Result, error) {
	apiURL := fmt.Sprintf(apiBase+"/wxa/getwxacodeunlimit?access_token=%s", "")
	resp, err := goz.Post(apiURL, goz.Options{
		JSON:  params,
		Debug: w.opts.Debug,
	})

	if err != nil {
		return nil, err
	}

	body, err := resp.GetBody()
	if err != nil {
		return nil, err
	}

	contentType := resp.GetHeaderLine("Content-Type")
	// 返回的是二进制图片
	if strings.Contains(contentType, "image/") {
		return Result(body), nil
	}

	return nil, errors.New(string(body))
}
