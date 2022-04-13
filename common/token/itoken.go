package token

// IToken: token interface
type IToken interface {
	GetToken() (string, error)
}
