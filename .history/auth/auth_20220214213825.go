package auth

import (
	"context"
	"errors"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type TokenInfo interface {
	GetAccessToken() string
	GetTokenType() string
	GetExpiresAt() int64
	EncodeToJSON() ([]byte, error)
}

type Auther interface {
	GenerateToken(ctx context.Context, userID string) (TokenInfo, error)
}
