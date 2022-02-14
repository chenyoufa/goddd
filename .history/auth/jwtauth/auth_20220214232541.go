package jwtauth

import (
	"context"
	"gdemo1/auth"
	"time"
)

type Options struct {
	expiretime int
}

type JWTAuth struct {
	opts Options
}

func (a *JWTAuth) GenerateToken(ctx context.Context) (auth.TokenInfo, error) {
	now := time.Now()
	exittime := now.Add(time.Duration(a.opts.expiretime)).Unix()
}
func (a *JWTAuth) DestroyToken(ctx context.Context, accessToken string) error {

}
func (a *JWTAuth) ParseUserID(ctx context.Context, accessToken string) (string, error) {

}
func (a *JWTAuth) Release() error {

}
