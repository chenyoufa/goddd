package jwtauth

func GetAccessToken() string

// 获取令牌类型
func GetTokenType() string

// 获取令牌到期时间戳
func GetExpiresAt() int64

// JSON编码
func EncodeToJSON() ([]byte, error)
