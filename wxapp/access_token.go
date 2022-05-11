package wxapp

import (
	"fmt"
	"time"
)

// AccessToken: default access_token handler
type AccessToken struct {
	client *Client
}

// NewAccessToken: init access_token handler
func NewAccessToken(c *Client) *AccessToken {
	return &AccessToken{c}
}

// GetToken: get access_token, from cache or api
func (t *AccessToken) GetToken() (string, error) {
	cacheKey := t.client.GetAccessTokenCacheKey()

	cache := t.client.GetCacheHandler()

	// get access_token from cache
	if v, err := cache.Get(cacheKey); err == nil && v != nil {
		if token, ok := v.(string); ok {
			return token, nil
		}
	}

	// get access_token from api
	return t.RefreshToken()
}

// RefreshToken: refresh access_token
func (t *AccessToken) RefreshToken() (string, error) {
	cacheKey := t.client.GetAccessTokenCacheKey()

	cache := t.client.GetCacheHandler()

	// get access_token from api
	res, err := t.client.FetchAccessToken()
	if err != nil {
		return "", fmt.Errorf("fetch token failed: %v", err)
	}

	pres := res.Parsed()
	accessToken := pres.Get("access_token").String()
	if accessToken == "" {
		return "", fmt.Errorf("invalid access_token")
	}

	expire := (pres.Get("expires_in").Int() - 300) * int64(time.Second)

	// set access_token to cache
	cache.Set(cacheKey, accessToken, time.Duration(expire))

	return accessToken, nil
}

// SetToken: set access_token to cache
func (t *AccessToken) SetToken(token string, expire time.Duration) error {
	cacheKey := t.client.GetAccessTokenCacheKey()

	cache := t.client.GetCacheHandler()

	// set access_token to cache
	return cache.Set(cacheKey, token, expire)
}
