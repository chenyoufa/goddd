package main

import (
	"context"
	"gdemo1/logger"
	loggerhook "gdemo1/logger/hook"
	loggergormhook "gdemo1/logger/hook/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dialector := mysql.Open("fage:Fage501526~@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(dialector)
	if err != nil {

	}
	var hookLevels []logrus.Level
	h := loggerhook.New(loggergormhook.New(db),
		loggerhook.SetMaxQueues(1),
		loggerhook.SetMaxWorks(1),
		loggerhook.SetLevels(hookLevels...),
	)
	logger.AddHook(h)

	logger.NewTagContext(context.Background(), "__main__")
	logger.WithContext(context.Background()).Infoln("fff")
}
