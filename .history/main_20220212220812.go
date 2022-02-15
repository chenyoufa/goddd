package main

import (
	"context"
	"fmt"
	"gdemo1/logger"
	loggerhook "gdemo1/logger/hook"
	loggergormhook "gdemo1/logger/hook/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dialector := mysql.Open("fage:Fage501526~@/mytest?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(dialector)
	if err != nil {
		fmt.Println(err)
	}
	logger.NewTagContext(context.Background(), "__main__")

	var hookLevels []logrus.Level
	h := loggerhook.New(loggergormhook.New(db),
		loggerhook.SetMaxQueues(1),
		loggerhook.SetMaxWorks(1),
		loggerhook.SetLevels(hookLevels...),
	)
	logger.AddHook(h)

	logger.WithContext(context.Background()).Infoln("fff")
}