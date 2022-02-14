package auth

//token 信息
type TokenInfo interface {
	GetType() string
	GetaccessToke() string
	GetExitTime() int64
	EncodJSON() (byte, error)
}
