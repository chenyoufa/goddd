package jwtauth

type tokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   int64  `json:"expires_at"`
}

//获取令牌类型
func (t *tokenInfo) GetTokenType() string {
	return ""
}

//获取令牌过期时间
func (t *tokenInfo) GetExpiresAt() int64 {
	return 0
}

//获取令牌
func (t *tokenInfo) GetAccessToken() string {
	return ""
}

//json 编码
func (t *tokenInfo) EncodeToJSON() ([]byte, error) {

}
