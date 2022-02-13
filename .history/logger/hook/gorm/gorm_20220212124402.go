package gorm

import "github.com/sirupsen/logrus"

type gorm struct {
}

func (g *gorm) exec(l *logrus.Entry) error {
	return nil
}
func (g *gorm) close() {

}
