package jwtauth

import (
	"encoding/json"
)

type Tokeninfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"toke_type"`
	ExpiresAt   int64  `json:"exit_at"`
}

func (a *Tokeninfo) GetAccessToken() string {
	return a.AccessToken
}
func (a *Tokeninfo) GetTokenType() string {
	return a.TokenType
}
func (a *Tokeninfo) GetExitAT() int64 {
	return a.ExpiresAt
}
func (a *Tokeninfo) EncodeToJson() ([]byte, error) {
	return json.Marshal(a.AccessToken)
}
