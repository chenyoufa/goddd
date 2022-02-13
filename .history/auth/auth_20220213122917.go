package auth

//令牌信息
type TokenInfo interface {
	//获取令牌类型
	GetTokenType() string
	//获取令牌过期时间
	GetExpiresAt() int64
	GetAccessToken() string
	EncodeToJSON() ([]byte, error)
}
