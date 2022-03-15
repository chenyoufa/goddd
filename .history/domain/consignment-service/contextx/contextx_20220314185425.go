package contextx

import (
	"context"
)

type (
	transCtx     struct{}
	noTransCtx   struct{}
	transLockCtx struct{}
	userIDCtx    struct{}
	userNameCtx  struct{}
	traceIDCtx   struct{}
)

func NewTrans(ctx context.Context, trans interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}
func FromTrans(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(transCtx{})
	return v, v != nil
}

func NewNoTrans(ctx context.Context) context.Context {
	return context.WithValue(ctx, noTransCtx{}, true)
}
func FromNoTrans(ctx context.Context) bool {
	v := ctx.Value(noTransCtx{})
	return v != nil && v.(bool)
}
func NewTransLock(ctx context.Context) context.Context {
	return context.WithValue(ctx, transLockCtx{}, true)
}
func FromTransLock(ctx context.Context) bool {
	v := ctx.Value(transLockCtx{})
	return v != nil && v.(bool)
}
func NewUserID(ctx context.Context, userID uint64) context.Context {
	return context.WithValue(ctx, userIDCtx{}, userID)
}
func FromUserID(ctx context.Context) uint64 {
	v := ctx.Value(userIDCtx{})
	if v != nil {
		if s, ok := v.(uint64); ok {
			return s
		}
	}
}
