package auth

import "context"

type TokenInfo interface {
	GetAccessToken() string
	GetTokenType() string
	GetExpiresAt() int64
	EncodeJSON() (byte, error)
}

type Auther interface {
	GenerateToken(ctx context.Context) (TokenInfo, error)
	DestroyToken(ctx context.Context, accessToken string) error
	ParseUserID(ctx context.Context, accessToken string) (string, error)
	Release() error
}
