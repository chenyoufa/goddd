package jwtauth

type tokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   int64  `json:"expires_at"`
}

//获取令牌类型
func (t *tokenInfo) GetTokenType() string
//获取令牌过期时间
GetExpiresAt() int64
//获取令牌
GetAccessToken() string
//json 编码
EncodeToJSON() ([]byte, error)