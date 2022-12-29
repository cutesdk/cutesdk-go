package wxmp

import (
	"fmt"
	"time"
)

// JsapiTicket: default jsapi_ticket handler
type JsapiTicket struct {
	cli *Client
}

// NewJsapiTicket: init jsapi_ticket handler
func NewJsapiTicket(cli *Client) *JsapiTicket {
	return &JsapiTicket{cli}
}

// GetToken: get jsapi_ticket, from cache or api
func (t *JsapiTicket) GetToken() (string, error) {
	cacheKey := t.cli.GetJsapiTicketCacheKey()

	cache := t.cli.GetCacheHandler()

	// get jsapi_ticket from cache
	if v, err := cache.Get(cacheKey); err == nil && v != nil {
		if ticket, ok := v.(string); ok {
			return ticket, nil
		}
	}

	// get jsapi_ticket from api
	return t.RefreshToken()
}

// RefreshToken: get jsapi_ticket
func (t *JsapiTicket) RefreshToken() (string, error) {
	cacheKey := t.cli.GetJsapiTicketCacheKey()

	cache := t.cli.GetCacheHandler()

	// get jsapi_ticket from api
	res, err := t.cli.FetchJsapiTicket()
	if err != nil {
		return "", fmt.Errorf("fetch ticket failed: %v", err)
	}

	pres := res.Parsed()
	jsapiTicket := pres.Get("ticket").String()
	if jsapiTicket == "" {
		return "", fmt.Errorf("invalid ticket")
	}

	expire := (pres.Get("expires_in").Int() - 300) * int64(time.Second)

	// set jsapi_ticket to cache
	cache.Set(cacheKey, jsapiTicket, time.Duration(expire))

	return jsapiTicket, nil
}

// SetToken: set jsapi_ticket to cache
func (t *JsapiTicket) SetToken(token string, expire time.Duration) error {
	cacheKey := t.cli.GetJsapiTicketCacheKey()

	cache := t.cli.GetCacheHandler()

	// set jsapi_ticket to cache
	return cache.Set(cacheKey, token, expire)
}
