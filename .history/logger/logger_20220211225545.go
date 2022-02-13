package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

//设置日志等级
//TraceID UserID  UserName Tag Stack 等关键字段的输出
//支持日志钩子写入到`Gorm`
//设置日子输出类型 json 、txt

func SetLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

func SetFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(new(logrus.JSONFormatter))
	default:
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}

}

func AddHook(hook logrus.Hook) {
	logrus.AddHook(hook)
}

type (
	traceIDKey  struct{}
	userIDKey   struct{}
	userNameKey struct{}
	tagkey      struct{}
	stackKey    struct{}
)

func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
func NewUserIDContext(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func FromUserIDContext(ctx context.Context) int64 {
	v := ctx.Value(userIDKey{})
	if v != nil {
		if s, ok := v.(int64); ok {
			return s
		}
	}
	return 0
}

func NewUserNameContext(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameKey{}, userName)
}

func FromUserNameContext(ctx context.Context) string {
	v := ctx.Value(userNameKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagkey{}, tag)
}

func FromTagContext(ctx context.Context) string {
	v := ctx.Value(tagkey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
func NewStackContext(ctx context.Context, stack error) context.Context {
	return context.WithValue(ctx, stackKey{}, stack)
}

func FromStackContext(ctx context.Context) error {
	v := ctx.Value(stackKey{})
	if v != nil {
		if s, ok := v.(error); ok {
			return s
		}
	}
	return nil
}

const (
	traceID_Key  = "trace_ID_Key"
	userID_Key   = "user_ID_Key"
	userName_Key = "userName_Key"
	tag_Key      = "tag_key"
	stack_Key    = "stack_Key"
)

func WithContext(ctx context.Context) logrus.Entry {
	fields := logrus.Fields{}
	if v := FromTraceIDContext(ctx); v != "" {
		fields[traceID_Key] = v
	}
	if v := FromUserIDContext(ctx); v != "" {
		fields[traceID_Key] = v
	}
	if v := FromUserNameContext(ctx); v != "" {
		fields[traceID_Key] = v
	}
	if v := FromTagContext(ctx); v != "" {
		fields[traceID_Key] = v
	}
	if v := FromStackContext(ctx); v != "" {
		fields[traceID_Key] = v
	}
}
