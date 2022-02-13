package jwtauth

import (
	"context"

	"github.com/dgrijalva/jwt-go"
)

const defaultKey = "gin-admin"

type options struct {
	signingMethod jwt.SigningMethod
	signingKey    interface{}
	tokenType     string
	expired       int
}

//生成令牌
func GenerateToken(ctx context.Context, userID string) (TokenInfo, error) {

}

//销毁
func DestroyToken(ctx context.Context, accessToken string) error {

}

//解析用户ID
func ParseUserID(ctx context.Context, accessToken string) (string, error) {

}

//释放资源
func Release() error {

}
