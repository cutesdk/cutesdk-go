package wxapp

import (
	"fmt"
	"time"

	"github.com/cutesdk/cutesdk-go/comm/cache"
	"github.com/idoubi/goutils"
)

const (
	AccessToken = "ACCESS_TOKEN"
)

// GetAccessToken api to get access_token
func (w *WxApp) GetAccessToken() (Result, error) {
	apiPath := "/cgi-bin/token"

	res, err := w.ApiGet(apiPath, map[string]string{
		"grant_type": "client_credential",
		"appid":      w.opts.Appid,
		"secret":     w.opts.AppSecret,
	})

	return res, err
}

// getAccessToken 获取 access_token
func (w *WxApp) getAccessToken() (string, error) {
	// custom handler to get access_token
	if w.opts.accessTokenHandler != nil {
		return w.opts.accessTokenHandler(w)
	}

	// 默认的获取 access_token 逻辑

	// 构造缓存 access_token 的 key
	str := fmt.Sprintf("grant_type=%s&appid=%s&secret=%s", "client_credential", w.opts.Appid, w.opts.AppSecret)
	cacheKey := fmt.Sprintf("cutesdk.wxapp.access_token.%s", goutils.MD5(str))

	fmt.Println(cacheKey)

	c := cache.New()

	c.Set("foo", "bar2343", 100*time.Second)
	vvv, ooo := c.Get("foo")
	fmt.Println(vvv, ooo)

	// 从缓存获取 access_token
	cacheData, err := c.Get(cacheKey)
	fmt.Println("cache", cacheData, err)

	// 取缓存里的 access_token 返回
	if err == nil && cacheData != nil {
		return cacheData.(string), nil
	}

	// 请求接口获取 access_token
	res, err := w.GetAccessToken()
	if err != nil {
		return "", err
	}

	pres := res.Parsed()
	accessToken := pres.Get("access_token").String()
	if accessToken == "" {
		return "", fmt.Errorf("invalid access_token: %s", res)
	}

	// 缓存 access_token
	expiresIn := pres.Get("expires_in").Int()
	expire := time.Duration(expiresIn-180) * time.Second

	c.Set(cacheKey, accessToken, expire)
	vv, oo := c.Get(cacheKey)
	fmt.Println(vv, oo, accessToken, expire)

	return accessToken, nil
}
