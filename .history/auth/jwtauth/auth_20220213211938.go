package jwtauth

import (
	"context"
	"gdemo1/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const defaultKey = "gin-admin"

type options struct {
	signingMethod jwt.SigningMethod
	signingKey    interface{}
	tokenType     string
	expired       int
	keyfunc       jwt.Keyfunc
}

var defaultOptions = options{
	tokenType:     "Bearer",
	expired:       7200,
	signingMethod: jwt.SigningMethodHS512,
	signingKey:    []byte(defaultKey),
	keyfunc: func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(defaultKey), nil
	},
}

type Option func(*options)

//设置签名方式
func SetSigningMethod(method jwt.SigningMethod) func(*options) {
	return func(o *options) {
		o.signingMethod = method
	}
}

//设定签名key
func SetSigningKey(key interface{}) func(*options) {
	return func(o *options) {
		o.signingKey = key
	}
}

//设置验证key的回调函数
func SetKeyfunc(keyFunc jwt.Keyfunc) func(*options) {
	return func(o *options) {
		o.keyfunc = keyFunc
	}
}

//设置令牌过期时长 单位秒，默认7200
func SetExpired(expired int) func(*options) {
	return func(o *options) {
		o.expired = expired
	}
}

type JWTAuth struct {
	opts *options
}

//生成令牌
func (jt *JWTAuth) GenerateToken(ctx context.Context, userID string) (auth.TokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(jt.opts.expired) * time.Second).Unix()
	token := jwt.NewWithClaims(jt.opts.signingMethod, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject:   userID,
	})
	tokenString, err := token.SignedString(jt.opts.signingKey)
	if err != nil {
		return nil, err
	}
}

//销毁
func (jt *JWTAuth) DestroyToken(ctx context.Context, accessToken string) error {

}

//解析用户ID
func (jt *JWTAuth) ParseUserID(ctx context.Context, accessToken string) (string, error) {

}

//释放资源
func (jt *JWTAuth) Release() error {

}
