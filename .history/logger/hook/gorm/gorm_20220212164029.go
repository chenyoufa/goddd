package gorm

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HookGorm struct {
	gorm.DB
}

type LoggerEntry struct {
	ID        uint      `gorm:"primaryKey;"`
	Level     string    `gorm:"size:20;index;"`
	TraceID   string    `gorm:"size:128;index"`
	UserID    uint64    `gorm:"index"`
	UserName  string    ``
	Tag       string    ``
	Message   string    ``
	Data      string    ``
	CreatedAt time.Time ``
}

func (g *HookGorm) exec(l *logrus.Entry) error {
	return nil
}
func (g *HookGorm) close() {
	g.close()
}
