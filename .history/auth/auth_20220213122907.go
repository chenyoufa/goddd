package auth

//令牌信息
type TokenInfo interface {
	//获取令牌类型
	GetTokenType() string
	GetExpiresAt() int64
	GetAccessToken() string
	EncodeToJSON() ([]byte, error)
}
