package jwtauth

import (
	"context"
	"fmt"
	"gdemo1/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const defaultkey = "fage-good"

type options struct {
	signingMethod jwt.SigningMethod
	sIgningKey    interface{}
	expiresAt     int
	tokenType     string
	keyfunc       jwt.Keyfunc
}
type JWTAuth struct {
	opt   *options
	store Storer
}

var defaultOPtions = options{
	tokenType:     "Bearer",
	expiresAt:     7200,
	signingMethod: jwt.SigningMethodHS512,
	sIgningKey:    []byte(defaultkey),
	keyfunc: func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(defaultkey), nil
	},
}

type Options func(*options)

func SetTokenType(stype string) Options {
	return func(o *options) {
		o.tokenType = stype
	}
}
func SetExpiresAt(val int) Options {
	return func(o *options) {
		o.expiresAt = val
	}
}
func SetSigningMethod(method jwt.SigningMethod) Options {
	return func(o *options) {
		o.signingMethod = method
	}
}
func SetSIgningKey(val interface{}) Options {
	return func(o *options) {
		o.sIgningKey = val
	}
}

func SetKeyfunc(val jwt.Keyfunc) Options {
	return func(o *options) {
		o.keyfunc = val
	}
}

func (a *JWTAuth) GeneraToken(ctx context.Context, userID string) (auth.TokenInfoer, error) {

	now := time.Now()
	expiresAt := now.Add(time.Duration(a.opt.expiresAt)).Unix()
	token := jwt.NewWithClaims(a.opt.signingMethod, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject:   userID,
	})
	tokenString, err := token.SignedString(a.opt.sIgningKey)
	if err != nil {
		return nil, err
	}
	return &Tokeninfo{
		AccessToken: tokenString,
		TokenType:   a.opt.tokenType,
		ExpiresAt:   expiresAt,
	}, nil
}

func (a *JWTAuth) callStore(fn func(store Storer) error) error {
	if store := a.store; store != nil {
		return fn(store)
	}
	return nil
}

func (a *JWTAuth) parseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.StandardClaims{}, a.opt.keyfunc)
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(*jwt.StandardClaims), nil
}

func (a *JWTAuth) DestoryToken(ctx context.Context, tokenString string) error {
	claims, err := a.parseToken(tokenString)
	if err != nil {
		return err
	}

	return a.callStore(func(store Storer) error {
		expired := time.Unix(claims.ExpiresAt, 0).Sub(time.Now())
		fmt.Println("expired,", expired)
		return store.Set(ctx, tokenString, expired)
	})
}

func (a *JWTAuth) ParseUserId(ctx context.Context, tokenString string) (string, error) {
	if tokenString != "" {
		return "", auth.ErrInvalidToken
	}
	claims, err := a.parseToken(tokenString)
	if err != nil {
		return "", err
	}

	err = a.callStore(func(store Storer) error {
		if exists, err := store.Check(ctx, tokenString); err != nil {
			return err
		} else if exists {
			return auth.ErrInvalidToken
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return claims.Subject, nil

}
func (a *JWTAuth) Release() error {
	return a.callStore(func(store Storer) error {
		return store.Close()
	})
}

func New(store Storer, opts ...Options) *JWTAuth {
	o := defaultOPtions
	for _, opt := range opts {
		opt(&o)
	}
	return &JWTAuth{
		opt:   &o,
		store: store,
	}
}
