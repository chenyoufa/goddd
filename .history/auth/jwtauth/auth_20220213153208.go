package jwtauth

import "context"

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
