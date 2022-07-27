package wxmp

import (
	"fmt"
	"time"

	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/cutesdk/cutesdk-go/wxopen"
)

// SetAuthorizer: set authorizer info
func (ins *Instance) SetAuthorizer(provider *wxopen.Instance, refreshToken string) error {
	// set default authorizer_access_token cache key
	accessTokenCacheKey := fmt.Sprintf("wxmp.authorizer_access_token.%s.%s", ins.GetAppid(), provider.GetComponentAppid())
	ins.SetAuthorizerAccessTokenCacheKey(accessTokenCacheKey)

	// set default authorizer_access_token handler
	ins.SetAuthorizerAccessTokenHandler(NewAuthorizerAccessToken(ins))

	return nil
}

// GetAuthorizerAccessToken: get authorizer_access_token from cache or api
func (ins *Instance) GetAuthorizerAccessToken() (string, error) {
	return ins.GetAuthorizerAccessTokenHandler().GetToken()
}

// RefreshAuthorizerAccessToken: refresh authorizer_access_token
func (ins *Instance) RefreshAuthorizerAccessToken() (string, error) {
	return ins.GetAuthorizerAccessTokenHandler().RefreshToken()
}

// SetAuthorizerAccessToken: set authorizer_access_token
func (ins *Instance) SetAuthorizerAccessToken(token string, expire time.Duration) error {
	return ins.GetAuthorizerAccessTokenHandler().SetToken(token, expire)
}

// GetAuthorizerRefreshToken: get authorizer_refresh_token
func (ins *Instance) GetAuthorizerRefreshToken() string {
	return ins.opts.AuthorizerRefreshToken
}

// GetAuthorizerProvider: get authorizer provider
func (ins *Instance) GetAuthorizerProvider() *wxopen.Instance {
	return ins.opts.AuthorizerProvider
}

// GetAuthorizerAccessTokenCacheKey: get authorizer_access_token cache key
func (ins *Instance) GetAuthorizerAccessTokenCacheKey() string {
	return ins.authorizerAccessTokenCacheKey
}

// SetAuthorizerAccessTokenCacheKey: set authorizer_access_token cache key
func (ins *Instance) SetAuthorizerAccessTokenCacheKey(cacheKey string) error {
	ins.authorizerAccessTokenCacheKey = cacheKey

	return nil
}

// GetAuthorizerAccessTokenHandler: get authorizer_access_token handler
func (ins *Instance) GetAuthorizerAccessTokenHandler() token.IToken {
	return ins.authorizerAccessTokenHandler
}

// SetAuthorizerAccessTokenHandler: set authorizer_access_token handler
func (ins *Instance) SetAuthorizerAccessTokenHandler(handler token.IToken) error {
	ins.authorizerAccessTokenHandler = handler

	return nil
}
