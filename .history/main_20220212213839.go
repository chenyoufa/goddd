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
)

func main() {

	db, err := mysql.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	h := loggerhook.New(loggergormhook.New(db),
		loggerhook.SetLevels(logrus.AllLevels...))
	logger.AddHook(h)
	logger.WithContext(context.Background()).Infoln("")
}
