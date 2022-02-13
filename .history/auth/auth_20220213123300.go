package auth

import "context"

//令牌信息
type TokenInfo interface {
	//获取令牌类型
	GetTokenType() string
	//获取令牌过期时间
	GetExpiresAt() int64
	//获取令牌
	GetAccessToken() string
	//json 编码
	EncodeToJSON() ([]byte, error)
}

type Auther interface {
	//生成令牌
	GenerateToken(ctx context.Context, userID string) (TokenInfo, error)
	//销毁
	DestroyToken(ctx context.Context, accessToken string) error
	//解析用户ID
	ParseUserID(ctx context.Context, accessToken string) (string, error)
	//释放资源
	Release() error
}
