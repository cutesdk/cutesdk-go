package wxopen

import (
	"fmt"
	"time"
)

// VerifyTicket: default component_verify_ticket handler
type VerifyTicket struct {
	cli *Client
}

// NewVerifyTicket: init component_verify_ticket handler
func NewVerifyTicket(cli *Client) *VerifyTicket {
	return &VerifyTicket{cli}
}

// GetToken: get component_verify_ticket, from cache
func (t *VerifyTicket) GetToken() (string, error) {
	cacheKey := t.cli.GetVerifyTicketCacheKey()

	cache := t.cli.GetCacheHandler()

	// get component_verify_ticket from cache
	if v, err := cache.Get(cacheKey); err == nil && v != nil {
		if token, ok := v.(string); ok {
			return token, nil
		}
	}

	return "", fmt.Errorf("invalid component_verify_ticket")
}

// RefreshToken: refresh component_verify_ticket
func (t *VerifyTicket) RefreshToken() (string, error) {
	cacheKey := t.cli.GetVerifyTicketCacheKey()

	cache := t.cli.GetCacheHandler()

	cache.Delete(cacheKey)

	return "", fmt.Errorf("invalid component_verify_ticket")
}

// SetToken: set component_verify_ticket to cache
func (t *VerifyTicket) SetToken(token string, expire time.Duration) error {
	cacheKey := t.cli.GetVerifyTicketCacheKey()

	cache := t.cli.GetCacheHandler()

	// set component_verify_ticket to cache
	return cache.Set(cacheKey, token, expire)
}
