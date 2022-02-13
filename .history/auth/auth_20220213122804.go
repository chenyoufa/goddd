package auth

//令牌信息
type accessinfo interface {
	GetTokenType()
	GetExpiresAt()
	GetAccessToken()
	EncodeToJSON()
}
