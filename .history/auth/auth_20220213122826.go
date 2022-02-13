package auth

//令牌信息
type accessinfo interface {
	GetTokenType() string
	GetExpiresAt() int64
	GetAccessToken() string
	EncodeToJSON() ([]byte, error)
}
