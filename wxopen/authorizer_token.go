package wxopen

import (
	"fmt"
	"time"
)

// AuthorizerToken: default authorizer_access_token handler
type AuthorizerToken struct {
	cli          *Client
	appid        string
	refreshToken string
}

// NewAuthorizerToken: init authorizer_access_token handler
func NewAuthorizerToken(cli *Client, appid, refreshToken string) *AuthorizerToken {
	return &AuthorizerToken{cli, appid, refreshToken}
}

// GetToken: get authorizer_access_token, from cache or api
func (t *AuthorizerToken) GetToken() (string, error) {
	cacheKey := t.cli.GetAuthorizerTokenCacheKey()

	cache := t.cli.GetCacheHandler()

	// get authorizer_access_token from cache
	if v, err := cache.Get(cacheKey); err == nil && v != nil {
		if token, ok := v.(string); ok {
			return token, nil
		}
	}

	// get authorizer_access_token from api
	return t.RefreshToken()
}

// RefreshToken: refresh authorizer_access_token
func (t *AuthorizerToken) RefreshToken() (string, error) {
	cacheKey := t.cli.GetAuthorizerTokenCacheKey()

	cache := t.cli.GetCacheHandler()

	// get authorizer_access_token from api
	res, err := t.cli.FetchAuthorizerToken(t.appid, t.refreshToken)
	if err != nil {
		return "", fmt.Errorf("fetch token failed: %v", err)
	}

	pres := res.Parsed()
	authorizerToken := pres.Get("authorizer_access_token").String()
	if authorizerToken == "" {
		return "", fmt.Errorf("%v", pres)
	}

	expire := (pres.Get("expires_in").Int() - 300) * int64(time.Second)

	// set authorizer_access_token to cache
	cache.Set(cacheKey, authorizerToken, time.Duration(expire))

	return authorizerToken, nil
}

// SetToken: set authorizer_access_token to cache
func (t *AuthorizerToken) SetToken(token string, expire time.Duration) error {
	cacheKey := t.cli.GetAuthorizerTokenCacheKey()

	cache := t.cli.GetCacheHandler()

	// set authorizer_access_token to cache
	return cache.Set(cacheKey, token, expire)
}
