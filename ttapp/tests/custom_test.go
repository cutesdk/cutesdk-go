package tests

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cutesdk/cutesdk-go/common/app"
	"github.com/cutesdk/cutesdk-go/ttapp"
	"github.com/idoubi/goutils"
)

func TestSetFileCache(t *testing.T) {
	client, err := ttapp.NewClient(appid, secret, app.WithCache("file", map[string]interface{}{
		"dir": "./cache123",
	}))

	if err != nil {
		t.Fatalf("new client error: %v", err)
	}

	res, err := client.GetAccessToken()
	t.Error(res, err)
}

func TestSetRedisCache(t *testing.T) {
	client, err := ttapp.NewClient(appid, secret, app.WithCache("redis", map[string]interface{}{
		"dsn":     "redis://:test123@127.0.0.1:6379/1",
		"timeout": "3s",
	}))

	if err != nil {
		t.Fatalf("new client error: %v", err)
	}

	res, err := client.GetAccessToken()

	t.Error(res, err)
}

func TestSetAccessTokenCacheKey(t *testing.T) {
	client, err := ttapp.NewClient(appid, secret, app.WithCache("redis", map[string]interface{}{
		"dsn":     "redis://:test123@127.0.0.1:6379/1",
		"timeout": "3s",
	}))

	if err != nil {
		t.Fatalf("new client error: %v", err)
	}

	keyFields := struct {
		GrantType string `json:"grant_type"`
		Appid     string `json:"appid"`
		Secret    string `json:"secret"`
	}{"client_credential", client.GetAppid(), client.GetSecret()}

	jsonByte, _ := json.Marshal(keyFields)

	cacheKey := fmt.Sprintf("easywechat.kernel.access_token.%s", goutils.MD5(string(jsonByte)))

	client.SetAccessTokenCacheKey(cacheKey)

	res, err := client.GetAccessToken()

	t.Error(res, err)
}

func TestSetAccessTokenHandler(t *testing.T) {
	client := getClient()

	handler := newCustomAccessTokenHandler(client)

	client.SetAccessTokenHandler(handler)

	res, err := client.GetAccessToken()

	t.Error(res, err)
}

type customAccessTokenHandler struct {
	client *ttapp.Client
}

func newCustomAccessTokenHandler(client *ttapp.Client) *customAccessTokenHandler {
	return &customAccessTokenHandler{client}
}

func (c *customAccessTokenHandler) GetToken() (string, error) {
	cacheKey := c.client.GetAccessTokenCacheKey() + "custom"

	cache := c.client.GetCacheHandler()

	// get access_token from cache
	if v, err := cache.Get(cacheKey); err == nil && v != nil {
		return v.(string), nil
	}

	// get access_token from api
	res, err := c.client.FetchAccessToken()

	if err != nil || res.Get("data.access_token").String() == "" {
		return "", fmt.Errorf("fetch access_token failed: %v, %v", res, err)
	}

	// set access_token to cache
	cache.Set(cacheKey, res.Get("data.access_token").String(), 3*time.Second)

	return res.Get("data.access_token").String(), nil
}
