package logger

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

//设置日志等级
//TraceID/UserID 等关键字段的输出
//支持日志钩子写入到`Gorm`
//设置日志输出类型 json/txt

// Logger Logrus
// type Logger = logrus.Logger

// Entry logrus.Entry alias
type Entry = logrus.Entry

// Hook logrus.Hook alias
type Hook = logrus.Hook

type Level = logrus.Level

// Define logger level
const (
	PanicLevel Level = iota //恐慌
	FatalLevel              //致命
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// SetLevel Set logger level
func SetLevel(level Level) {
	logrus.SetLevel(level)
}

// SetFormatter Set logger output format (json/text)
func SetFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(new(logrus.JSONFormatter))
	default:
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}
}

// AddHook Add logger hook
func AddHook(hook Hook) {
	logrus.AddHook(hook)
}

// Define key
const (
	TraceIDKey  = "trace_id"  //跟踪
	UserIDKey   = "user_id"   //用户id
	UserNameKey = "user_name" //用户名
	TagKey      = "tag"       //目标
	StackKey    = "stack"     //堆栈
)

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

func FormTraceIDContext(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewUserIDContext(ctx context.Context, userID uint64) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func FormUserIDContext(ctx context.Context) uint64 {
	v := ctx.Value(userIDKey{})
	if v != nil {
		if s, ok := v.(uint64); ok {
			return s
		}
	}
	return 0
}
func NewUserNameContext(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameKey{}, userName)
}

func FormUserNameContext(ctx context.Context) string {
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

func WithContext(ctx context.Context) *Entry {
	fields := logrus.Fields{}
	if v := FormTraceIDContext(ctx); v != "" {
		fields[TraceIDKey] = v
	}
	if v := FormUserIDContext(ctx); v != 0 {
		fields[UserIDKey] = v
	}
	if v := FormUserNameContext(ctx); v != "" {
		fields[UserNameKey] = v
	}
	if v := FromTagContext(ctx); v != "" {
		fields[TagKey] = v
	}
	if v := FromStackContext(ctx); v != nil {
		fields[StackKey] = fmt.Sprintf("%+v", v)
	}
	return logrus.WithContext(ctx).WithFields(fields)
}

// var (
// 	Tracef         = logrus.Tracef
// 	Debugf         = logrus.Debugf
// 	Infof          = logrus.Infof
// 	Warnf          = logrus.Warnf
// 	Errorf         = logrus.Errorf
// 	Fatalf         = logrus.Fatal
// 	Panicf         = logrus.Panicf
// 	Printf         = logrus.Printf
// 	SetOutput      = logrus.SetReportCaller
// 	StandardLogger = logrus.StandardLogger
// 	ParseLevel     = logrus.ParseLevel
// )
