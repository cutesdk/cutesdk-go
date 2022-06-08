package tests

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/ttapp"
	"github.com/idoubi/goutils"
)

func TestSetFileCache(t *testing.T) {
	opts := &ttapp.Options{
		Appid:  appid,
		Secret: secret,
		Cache: &cache.Options{
			Driver: "file",
			Conf: map[string]interface{}{
				"dir": "./ttappcache",
			},
		},
	}
	ins, err := ttapp.New(opts)

	if err != nil {
		t.Fatalf("new client error: %v", err)
	}

	res, err := ins.GetAccessToken()
	t.Error(res, err)
}

func TestSetRedisCache(t *testing.T) {
	opts := &ttapp.Options{
		Appid:  appid,
		Secret: secret,
		Cache: &cache.Options{
			Driver: "redis",
			Conf: map[string]interface{}{
				"dsn":     "redis://:@127.0.0.1:6379/1",
				"timeout": "3s",
			},
		},
	}
	ins, err := ttapp.New(opts)

	if err != nil {
		t.Fatalf("new client error: %v", err)
	}

	res, err := ins.GetAccessToken()

	t.Error(res, err)
}

func TestSetAccessTokenCacheKey(t *testing.T) {
	opts := &ttapp.Options{
		Appid:  appid,
		Secret: secret,
		Cache: &cache.Options{
			Driver: "redis",
			Conf: map[string]interface{}{
				"dsn":     "redis://:@127.0.0.1:6379/1",
				"timeout": "3s",
			},
		},
	}
	ins, err := ttapp.New(opts)

	if err != nil {
		t.Fatalf("new client error: %v", err)
	}

	keyFields := struct {
		GrantType string `json:"grant_type"`
		Appid     string `json:"appid"`
		Secret    string `json:"secret"`
	}{"client_credential", ins.GetAppid(), ins.GetSecret()}

	jsonByte, _ := json.Marshal(keyFields)

	cacheKey := fmt.Sprintf("easywechat.kernel.access_token.%s", goutils.MD5(string(jsonByte)))

	ins.SetAccessTokenCacheKey(cacheKey)

	res, err := ins.GetAccessToken()

	t.Error(res, err)
}

func TestSetAccessTokenHandler(t *testing.T) {
	ins := getIns()

	handler := newCustomAccessTokenHandler(ins)

	ins.SetAccessTokenHandler(handler)

	res, err := ins.GetAccessToken()

	t.Error(res, err)
}

type customAccessTokenHandler struct {
	ins *ttapp.Instance
}

func newCustomAccessTokenHandler(ins *ttapp.Instance) *customAccessTokenHandler {
	return &customAccessTokenHandler{ins}
}

func (c *customAccessTokenHandler) GetToken() (string, error) {
	cacheKey := c.ins.GetAccessTokenCacheKey() + "custom"

	cache := c.ins.GetCacheHandler()

	// get access_token from cache
	if v, err := cache.Get(cacheKey); err == nil && v != nil {
		return v.(string), nil
	}

	return c.RefreshToken()
}

func (c *customAccessTokenHandler) RefreshToken() (string, error) {
	cacheKey := c.ins.GetAccessTokenCacheKey() + "custom"

	cache := c.ins.GetCacheHandler()

	// get access_token from cache
	if v, err := cache.Get(cacheKey); err == nil && v != nil {
		return v.(string), nil
	}

	// get access_token from api
	res, err := c.ins.FetchAccessToken()

	if err != nil || res.Get("data.access_token").String() == "" {
		return "", fmt.Errorf("fetch access_token failed: %v, %v", res, err)
	}

	// set access_token to cache
	cache.Set(cacheKey, res.Get("data.access_token").String(), 3*time.Second)

	return res.Get("data.access_token").String(), nil
}

func (c *customAccessTokenHandler) SetToken(token string, expire time.Duration) error {
	cacheKey := c.ins.GetAccessTokenCacheKey() + "custom"

	cache := c.ins.GetCacheHandler()

	// set access_token to cache
	return cache.Set(cacheKey, token, expire)
}
