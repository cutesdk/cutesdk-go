package wxmp

import (
	"fmt"
	"time"
)

// AuthorizerAccessToken: default authorizer_access_token handler
type AuthorizerAccessToken struct {
	ins *Instance
}

// NewAuthorizerAccessToken: init authorizer_access_token handler
func NewAuthorizerAccessToken(ins *Instance) *AuthorizerAccessToken {
	return &AuthorizerAccessToken{ins}
}

// GetToken: get authorizer_access_token, from cache or api
func (t *AuthorizerAccessToken) GetToken() (string, error) {
	cacheKey := t.ins.GetAuthorizerAccessTokenCacheKey()

	cache := t.ins.GetCacheHandler()

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
func (t *AuthorizerAccessToken) RefreshToken() (string, error) {
	cacheKey := t.ins.GetAuthorizerAccessTokenCacheKey()

	cache := t.ins.GetCacheHandler()

	// get authorizer_access_token from api
	res, err := t.ins.FetchAuthorizerAccessToken()
	if err != nil {
		return "", fmt.Errorf("fetch token failed: %v", err)
	}

	pres := res.Parsed()
	AuthorizeraccessToken := pres.Get("authorizer_access_token").String()
	if AuthorizeraccessToken == "" {
		return "", fmt.Errorf("%v", pres)
	}

	expire := (pres.Get("expires_in").Int() - 300) * int64(time.Second)

	// set authorizer_access_token to cache
	cache.Set(cacheKey, AuthorizeraccessToken, time.Duration(expire))

	return AuthorizeraccessToken, nil
}

// SetToken: set authorizer_access_token to cache
func (t *AuthorizerAccessToken) SetToken(token string, expire time.Duration) error {
	cacheKey := t.ins.GetAuthorizerAccessTokenCacheKey()

	cache := t.ins.GetCacheHandler()

	// set authorizer_access_token to cache
	return cache.Set(cacheKey, token, expire)
}
