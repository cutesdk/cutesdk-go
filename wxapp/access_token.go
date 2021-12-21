package wxapp

import (
	"fmt"
	"net/url"
	"time"

	"github.com/idoubi/goz"
	"github.com/muesli/cache2go"
)

// GetAccessToken 获取access_token
func (w *Wxapp) GetAccessToken() (Result, error) {
	queryParams := url.Values{}
	queryParams.Set("grant_type", "client_credential")
	queryParams.Set("appid", w.opts.Appid)
	queryParams.Set("secret", w.opts.Secret)

	apiUrl := apiBase + "/cgi-bin/token?" + queryParams.Encode()

	resp, err := goz.Get(apiUrl, goz.Options{
		Debug: w.opts.Debug,
	})

	if err != nil {
		return nil, err
	}

	body, err := resp.GetBody()

	return Result(body), err
}

// getAccessToken 获取access_token 内部调用
func (w *Wxapp) getAccessToken() (string, error) {
	// 执行自定义的缓存逻辑
	if w.opts.GetAccessToken != nil {
		return w.opts.GetAccessToken(w)
	}

	// 默认缓存
	cache := cache2go.Cache("tokenCache")

	cacheKey := fmt.Sprintf("wxa:access_token:%s", w.opts.Appid)
	cacheData, err := cache.Value(cacheKey)
	if err == nil {
		// 取缓存的access_token返回
		return cacheData.Data().(string), nil
	}

	// 请求接口获取access_token
	res, err := w.GetAccessToken()
	if err != nil {
		return "", err
	}

	pres := res.Parsed()
	accessToken := pres.Get("access_token").String()
	// 获取access_token失败
	if accessToken == "" {
		return "", fmt.Errorf("get access_token failed: %s\n", res)
	}

	// 缓存access_token
	expiresIn := pres.Get("expires_in").Int()
	cache.Add(cacheKey, time.Duration(expiresIn-180)*time.Second, accessToken)

	return accessToken, nil
}
