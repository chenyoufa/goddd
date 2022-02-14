package auth

type TokenInfo interface {
	GetAccessToken() string
	GetTokenType() string
	GetExitTime() int64
	EncodeJSON() (byte, error)
}
