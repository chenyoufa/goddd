package jwtauth

import (
	"context"
	"gdemo1/auth"
	"time"
)
type Options struct{
	expiretime time.Time
}


type JWTAuth struct {

}

func (a *JWTAuth)GenerateToken(ctx context.Context) (auth.TokenInfo, error) {
	now:=time.Now()
	exittime:=
}
func  (a *JWTAuth) DestroyToken(ctx context.Context, accessToken string) error {

}
func (a *JWTAuth) ParseUserID(ctx context.Context, accessToken string) (string, error) {

}
func (a *JWTAuth) Release() error {

}
