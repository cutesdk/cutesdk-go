package token

import "time"

// IToken: token interface
type IToken interface {
	GetToken() (string, error)
	RefreshToken() (string, error)
	SetToken(string, time.Duration) error
}
