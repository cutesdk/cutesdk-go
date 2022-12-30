package wxopen

import (
	"fmt"
	"time"
)

// AccessToken: default component_access_token handler
type AccessToken struct {
	cli *Client
}

// NewAccessToken: init component_access_token handler
func NewAccessToken(cli *Client) *AccessToken {
	return &AccessToken{cli}
}

// GetToken: get component_access_token, from cache or api
func (t *AccessToken) GetToken() (string, error) {
	cacheKey := t.cli.GetAccessTokenCacheKey()

	cache := t.cli.GetCacheHandler()

	// get component_access_token from cache
	if v, err := cache.Get(cacheKey); err == nil && v != nil {
		if token, ok := v.(string); ok {
			return token, nil
		}
	}

	// get component_access_token from api
	return t.RefreshToken()
}

// RefreshToken: refresh component_access_token
func (t *AccessToken) RefreshToken() (string, error) {
	cacheKey := t.cli.GetAccessTokenCacheKey()

	cache := t.cli.GetCacheHandler()

	// get component_access_token from api
	res, err := t.cli.FetchAccessToken()
	if err != nil {
		return "", fmt.Errorf("fetch token failed: %v", err)
	}

	pres := res.Parsed()
	accessToken := pres.Get("component_access_token").String()
	if accessToken == "" {
		return "", fmt.Errorf("invalid component_access_token")
	}

	expire := (pres.Get("expires_in").Int() - 300) * int64(time.Second)

	// set component_access_token to cache
	cache.Set(cacheKey, accessToken, time.Duration(expire))

	return accessToken, nil
}

// SetToken: set component_access_token to cache
func (t *AccessToken) SetToken(token string, expire time.Duration) error {
	cacheKey := t.cli.GetAccessTokenCacheKey()

	cache := t.cli.GetCacheHandler()

	// set component_access_token to cache
	return cache.Set(cacheKey, token, expire)
}
