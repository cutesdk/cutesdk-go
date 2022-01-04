package wxopen

import (
	"github.com/idoubi/goz"
)

const (
	ComponentAccessToken = "COMPONENT_ACCESS_TOKEN"
)

// GetComponentAccessToken 获取 component_access_token
func (w *WxOpen) GetComponentAccessToken() (Result, error) {
	apiUrl := apiBase + "/cgi-bin/component/api_component_token"

	ticket, err := w.getComponentVerifyTicket()
	if err != nil {
		return nil, err
	}

	resp, err := goz.Post(apiUrl, goz.Options{
		Debug: w.opts.Debug,
		JSON: map[string]string{
			"component_appid":         w.opts.Appid,
			"component_appsecret":     w.opts.AppSecret,
			"component_verify_ticket": ticket,
		},
	})

	if err != nil {
		return nil, err
	}

	body, err := resp.GetBody()

	return Result(body), err
}

// getComponentAccessToken 获取 component_access_token
func (w *WxOpen) getComponentAccessToken() (string, error) {
	return "52_QLd-4ZVQLsHo-Te8LwDE_gdT4B_taPNBQgSqxvwVKTa9rZqDGW7MB0zC0qo7V3c8Jv49pOOzhvAp7AwCFTJxXbfbIB08OXuxZjn3guWmaF8DaUB4LzjZYjd9hXSGCykQiyy-FvgRGehc4JIFIAPiAFANFU", nil
}
