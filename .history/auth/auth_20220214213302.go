package auth

import (
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
