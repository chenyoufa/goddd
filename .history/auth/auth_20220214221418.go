package auth

import "context"

type TokenInfo interface {
	GetAccessToken() string
	GetTokenType() string
	GetExitTime() int64
	EncodeJSON() (byte, error)
}

type Auther interface {
	GennerToken(ctx context.Context) (TokenInfo, error)
}
