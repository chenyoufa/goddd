package auth

type TokenInfo interface {
	getType() string
	getaccessToke() string
	GetExitTime() int64
	EncodJSON() (byte, error)
}
