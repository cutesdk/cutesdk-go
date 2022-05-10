package wxopen

import (
	"fmt"
)

// ComponentVerifyTicket: default component_verify_ticket handler
type ComponentVerifyTicket struct {
	client *Client
}

// NewComponentVerifyTicket: init component_verify_ticket handler
func NewComponentVerifyTicket(c *Client) *ComponentVerifyTicket {
	return &ComponentVerifyTicket{c}
}

// GetToken: get component_verify_ticket, from cache
func (t *ComponentVerifyTicket) GetToken() (string, error) {
	cacheKey := t.client.GetComponentVerifyTicketCacheKey()

	cache := t.client.GetCacheHandler()

	// get component_verify_ticket from cache
	if v, err := cache.Get(cacheKey); err == nil && v != nil {
		if token, ok := v.(string); ok {
			return token, nil
		}
	}

	return "", fmt.Errorf("invalid component_verify_ticket")
}

// RefreshToken: refresh component_verify_ticket
func (t *ComponentVerifyTicket) RefreshToken() (string, error) {
	cacheKey := t.client.GetComponentVerifyTicketCacheKey()

	cache := t.client.GetCacheHandler()

	cache.Delete(cacheKey)

	return "", fmt.Errorf("invalid component_verify_ticket")
}
