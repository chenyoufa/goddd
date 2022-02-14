package jwtauth

import (
	"context"
	"gdemo1/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Options struct {
	expiretime    int
	signingMethod jwt.SigningMethod
	signingKey    interface{}
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
		Subject:   userID,
	})
	tokeinfo := token.SigningString(a.o)
}
func (a *JWTAuth) DestroyToken(ctx context.Context, accessToken string) error {

}
func (a *JWTAuth) ParseUserID(ctx context.Context, accessToken string) (string, error) {

}
func (a *JWTAuth) Release() error {

}
