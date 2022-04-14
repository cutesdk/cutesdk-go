package wxmp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/app"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// Client: wxmp client
type Client struct {
	*app.Client
}

// NewClient create wxmp client
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
		BaseUri: "https://api.weixin.qq.com",
	})

	// set default access_token cache key
	accessTokenCacheKey := fmt.Sprintf("wxmp.access_token.%s", c.GetAppid())
	c.SetAccessTokenCacheKey(accessTokenCacheKey)

	// set default access_token handler
	c.SetAccessTokenHandler(NewAccessToken(c))

	// set default jsapi_ticket cacheKey
	jsapiTicketCacheKey := fmt.Sprintf("wxmp.jsapi_ticket.%s", c.GetAppid())
	c.SetJsapiTicketCacheKey(jsapiTicketCacheKey)

	// set default jsapi_ticket handler
	c.SetJsapiTicketHandler(NewJsapiTicket(c))

	return c, nil
}
