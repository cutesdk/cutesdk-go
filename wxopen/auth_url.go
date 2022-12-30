package wxopen

import (
	"fmt"
	"net/url"
)

// GetRegisterAuthUrl: get register app auth url
func (cli *Client) GetRegisterAuthUrl(appid, redirectUri, copyWxVerify string) (string, error) {
	authUrl := fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/fastregisterauth?appid=%s&component_appid=%s&copy_wx_verify=%s&redirect_uri=%s", appid, cli.GetAppid(), copyWxVerify, url.QueryEscape(redirectUri))

	return authUrl, nil
}

// GetRebindAuthUrl: get rebind admin auth url
func (cli *Client) GetRebindAuthUrl(appid, redirectUri string) (string, error) {
	authUrl := fmt.Sprintf("https://mp.weixin.qq.com/wxopen/componentrebindadmin?appid=%s&component_appid=%s&redirect_uri=%s", appid, cli.GetAppid(), url.QueryEscape(redirectUri))

	return authUrl, nil
}

// GetPcAuthUrl: get pc auth url
func (cli *Client) GetPcAuthUrl(preAuthCode, redirectUri, authType string, extra map[string]string) (string, error) {
	authUrl := fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s", cli.GetAppid(), preAuthCode, url.QueryEscape(redirectUri), authType)
	for k, v := range extra {
		authUrl += fmt.Sprintf("&%s=%s", k, v)
	}

	return authUrl, nil
}

// GetH5AuthUrl: get pc auth url
func (cli *Client) GetH5AuthUrl(preAuthCode, redirectUri, authType string, extra map[string]string) (string, error) {
	authUrl := fmt.Sprintf("https://open.weixin.qq.com/wxaopen/safe/bindcomponent?action=bindcomponent&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s", cli.GetAppid(), preAuthCode, url.QueryEscape(redirectUri), authType)
	for k, v := range extra {
		authUrl += fmt.Sprintf("&%s=%s", k, v)
	}
	authUrl += "#wechat_redirect"

	return authUrl, nil
}
