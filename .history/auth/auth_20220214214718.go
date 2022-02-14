package auth

//token 信息接口
type TokenInfo interface {
	GetType() string
	GetaccessToke() string
	GetExitTime() int64
	EncodJSON() (byte, error)
}
