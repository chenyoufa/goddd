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
		if _, ok := v.(string); ok {

		}
	}

}
