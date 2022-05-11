package wxopen

import (
	"fmt"
	"time"
)

// ComponentAccessToken: default component_access_token handler
type ComponentAccessToken struct {
	client *Client
}

// NewComponentAccessToken: init component_access_token handler
func NewComponentAccessToken(c *Client) *ComponentAccessToken {
	return &ComponentAccessToken{c}
}

// GetToken: get component_access_token, from cache or api
func (t *ComponentAccessToken) GetToken() (string, error) {
	cacheKey := t.client.GetComponentAccessTokenCacheKey()

	cache := t.client.GetCacheHandler()

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
func (t *ComponentAccessToken) RefreshToken() (string, error) {
	cacheKey := t.client.GetComponentAccessTokenCacheKey()

	cache := t.client.GetCacheHandler()

	// get component_access_token from api
	res, err := t.client.FetchComponentAccessToken()
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
func (t *ComponentAccessToken) SetToken(token string, expire time.Duration) error {
	cacheKey := t.client.GetComponentAccessTokenCacheKey()

	cache := t.client.GetCacheHandler()

	// set component_access_token to cache
	return cache.Set(cacheKey, token, expire)
}
