package jwtauth

type Tokeninfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"toke_type"`
	ExpiresAt   int    `json:"expires_at"`
}

func (a *Tokeninfo) GetAccessToken() string {
	return a.AccessToken
}

// 获取令牌类型
func (a *Tokeninfo) GetTokenType() string {

}

// 获取令牌到期时间戳
func (a *Tokeninfo) GetExpiresAt() int64 {

}

// JSON编码
func (a *Tokeninfo) EncodeToJSON() ([]byte, error) {

}
