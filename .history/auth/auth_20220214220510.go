package auth

import "context"

//token 信息接口
type TokenInfo interface {
	GetType() string
	GetaccessToke() string
	GetExitTime() int64
	EncodJSON() (byte, error)
}

type Auther interface {
	GenerAccessToken(ctx context.Context, userID string) (TokenInfo, error)
	DeleteToken(ctx context.Context, accessToken string) error
	ParseUserID(ctx context.Context, accessToken string) (string, error)
	Release() error
}
