package jwtauth

import (
	"context"
	"go/token"
)

func GenerateToken(ctx context.Context) (token.TokenInfo, error)
func DestroyToken(ctx context.Context, accessToken string) error
func ParseUserID(ctx context.Context, accessToken string) (string, error)
func Release() error
