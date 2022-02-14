package auth

import "context"

type TokenInfo interface {
	GetAccessToken() string
	GetTokenType() string
	GetExitTime() int64
	EncodeJSON() (byte, error)
}

type Auther interface {
	GennerTokenInfo(ctx context.Context) (TokenInfo, error)
	DeleteToken(ctx context.Context, accessToken string) error
	ParseUserID(ctx context.Context, accessToken string) (string, error)
	Release(ctx context.Context) error
}
