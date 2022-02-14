package jwtauth

import (
	"context"
	"gdemo1/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Options struct {
	expiretime    int
	SigningMethod jwt.SigningMethod
}

type JWTAuth struct {
	opts Options
}

func (a *JWTAuth) GenerateToken(ctx context.Context, userID string) (auth.TokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(a.opts.expiretime)).Unix()
	token := jwt.NewWithClaims(a.opts.SigningMethod, jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
		ExpiresAt: expiresAt,
		Subject:   userId,
	})
}
func (a *JWTAuth) DestroyToken(ctx context.Context, accessToken string) error {

}
func (a *JWTAuth) ParseUserID(ctx context.Context, accessToken string) (string, error) {

}
func (a *JWTAuth) Release() error {

}
