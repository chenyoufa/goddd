package jwtauth

import (
	"context"
	"gdemo1/auth"
)

type JWTAuth struct {
}

func GenerateToken(ctx context.Context) (auth.TokenInfo, error) {

}
func DestroyToken(ctx context.Context, accessToken string) error {

}
func ParseUserID(ctx context.Context, accessToken string) (string, error) {

}
func Release() error {

}
