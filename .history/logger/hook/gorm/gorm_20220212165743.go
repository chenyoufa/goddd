package gorm

import (
	"gdemo1/logger"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HookGorm struct {
	db *gorm.DB
}

type LoggerEntry struct {
	ID        uint      `gorm:"primaryKey;"`
	Level     string    `gorm:"size:20;index;"`
	TraceID   string    `gorm:"size:128;index"`
	UserID    uint64    `gorm:"index"`
	UserName  string    `gorm:"size:64;index"`
	Tag       string    `gorm:"size:128;index"`
	Message   string    `gorm:"size:1024"`
	Data      string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"index"`
}

func (g *HookGorm) exec(entry *logrus.Entry) error {
	item := &LoggerEntry{
		Level:     entry.Level.String(),
		Message:   entry.Message,
		CreatedAt: entry.Time,
	}
	ctx := entry.Context
	item.TraceID = logger.FromTraceIDContext(ctx)
	item.UserID = uint64(logger.FromUserIDContext(ctx))
	item.UserName = logger.FromTagContext(ctx)
	item.Tag = logger.FromTagContext(ctx)
	data := entry.Data
	return nil
}
func (g *HookGorm) close() {
	g.close()
}
func New(db *gorm.DB) *HookGorm {

}
