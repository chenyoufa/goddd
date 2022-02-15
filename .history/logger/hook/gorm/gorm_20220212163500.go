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
	ID        uint      ``
	Level     string    ``
	TraceID   string    ``
	UserID    uint64    ``
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