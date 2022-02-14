package jwtauth

import "context"

func GenerateToken(ctx context.Context) (TokenInfo, error)
func DestroyToken(ctx context.Context, accessToken string) error
func ParseUserID(ctx context.Context, accessToken string) (string, error)
func Release() error
