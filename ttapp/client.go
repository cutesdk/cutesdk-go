package ttapp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/app"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// Client toutiao app client
type Client struct {
	*app.Client
}

// NewClient create toutiao client
func NewClient(appid, secret string, optfuncs ...app.OptionFunc) (*Client, error) {
	baseClient, err := app.NewClient(appid, secret, optfuncs...)
	if err != nil {
		return nil, fmt.Errorf("init client failed: %v", err)
	}

	c := &Client{baseClient}

	// set default request options
	c.SetRequestOptions(request.Options{
		Debug:   c.GetRequestOptions().Debug,
		Timeout: c.GetRequestOptions().Timeout,
		BaseUri: "https://developer.toutiao.com",
	})

	// set default access_token cache key
	cacheKey := fmt.Sprintf("ttapp.access_token.%s", c.GetAppid())
	c.SetAccessTokenCacheKey(cacheKey)

	// set default access_token handler
	c.SetAccessTokenHandler(NewAccessToken(c))

	return c, nil
}
