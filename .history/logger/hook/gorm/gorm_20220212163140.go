package gorm

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HookGorm struct {
	gorm.DB
}

func (g *HookGorm) exec(l *logrus.Entry) error {
	return nil
}
func (g *HookGorm) close() {
	g.close()
}
