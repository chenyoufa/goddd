package auth

//令牌信息
type TokenInfo interface {
	GetTokenType() string
	GetExpiresAt() int64
	GetAccessToken() string
	EncodeToJSON() ([]byte, error)
}
