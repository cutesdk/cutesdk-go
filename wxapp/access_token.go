package wxapp

import (
	"fmt"
	"time"

	"github.com/idoubi/goz"
	"github.com/muesli/cache2go"
)

// GetAccessToken 获取access_token
func (w *Wxapp) GetAccessToken() (Result, error) {
	apiURL := fmt.Sprintf(apiBase+"/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", w.opts.Appid, w.opts.Appsecret)
	resp, err := goz.Get(apiURL, goz.Options{
		Debug: w.opts.Debug,
	})

	if err != nil {
		return nil, err
	}

	body, err := resp.GetBody()

	return Result(body), err
}

// getAccessToken 获取access_token 内部调用
func (w *Wxapp) getAccessToken() string {
	// 执行自定义的缓存逻辑
	if w.opts.GetToken != nil {
		return w.opts.GetToken(w)
	}

	// 默认缓存
	cache := cache2go.Cache("tokenCache")

	cacheKey := fmt.Sprintf("wxa:access_token:%s", w.opts.Appid)
	cacheData, err := cache.Value(cacheKey)
	if err == nil {
		return cacheData.Data().(string)
	}

	res, err := w.GetAccessToken()
	if err != nil {
		return ""
	}

	pres := res.Parsed()
	accessToken := pres.Get("access_token").String()

	if accessToken != "" {
		expiresIn := pres.Get("expires_in").Int()
		cache.Add(cacheKey, time.Duration(expiresIn-180)*time.Second, accessToken)
	}

	return accessToken
}
