package jwtauth

type Tokeninfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"toke_type"`
	ExpiresAt   int    `json:"expires_at"`
}

func GetAccessToken() string {

}

// 获取令牌类型
func GetTokenType() string {

}

// 获取令牌到期时间戳
func GetExpiresAt() int64 {

}

// JSON编码
func EncodeToJSON() ([]byte, error) {

}
