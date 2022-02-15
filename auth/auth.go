package auth

import (
	"context"
	"errors"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

//接口信息接口
type TokenInfoer interface {
	GetAccessToken() string
	GetTokenType() string
	GetExitAT() int64
	EncodeToJson() ([]byte, error)
}

type Auther interface {
	GeneraToken(ctx context.Context, userID string) (TokenInfoer, error)
	DestoryToken(ctx context.Context, tokenString string) error
	ParseUserId(ctx context.Context, tokenString string) (int64, error)
	Release() error
}
