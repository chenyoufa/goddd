package gorm

import (
	"encoding/json"
	"fmt"
	"thrgoweb/pkg/logger"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Logger struct {
	ID        uint      `gorm:"primaryKey;"`
	Level     string    `gorm:"size:20;index"`
	TraceID   string    `gorm:"size:128;index"`
	UserID    uint64    `gorm:"index;"`
	UserName  string    `gorm:"size:64;index"`
	Tag       string    `gorm:"size:128;index"`
	Message   string    `gorm:"size:1024"`
	Data      string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"index"`
}

func New(db *gorm.DB) *Hook {
	db.AutoMigrate(new(Logger))
	return &Hook{
		db: db,
	}
}

type Hook struct {
	db *gorm.DB
}

func (h *Hook) Exec(entry *logrus.Entry) error {
	item := &Logger{
		Level:     entry.Level.String(),
		Message:   entry.Message,
		CreatedAt: entry.Time,
	}
	ctx := entry.Context
	item.TraceID = logger.FormTraceIDContext(ctx)
	item.UserID = logger.FormUserIDContext(ctx)
	item.UserName = logger.FormUserNameContext(ctx)
	item.Tag = logger.FromTagContext(ctx)
	data := entry.Data
	if v := logger.FromStackContext(ctx); v != nil {
		data["stack"] = fmt.Sprintf("%+v", v)
	}
	if len(data) > 0 {
		b, _ := json.Marshal(data)
		item.Data = string(b)
	}
	return h.db.Create(item).Error
}

func (h *Hook) Close() error {
	db, err := h.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
