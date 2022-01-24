package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Logger Logrus
type Logger = logrus.Logger

// Entry logrus.Entry alias
type Entry = logrus.Entry

// Hook logrus.Hook alias
type Hook = logrus.Hook

type Level = logrus.Level

// Define logger level
const (
	PanicLevel Level = iota
	FatalLevel
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
	TraceIDKey  = "trace_id"
	UserIDKey   = "user_id"
	UserNameKey = "user_name"
	TagKey      = "tag"
	StackKey    = "stack"
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
	v := ctx.Value(tranceid)
}
