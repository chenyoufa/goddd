package main

import (
	"context"
	"fmt"
	"gdemo1/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	h := loggerhook.New(loggergormhook.New(db),
		loggerhook.SetMaxWorkers(c.HookMaxThread),
		loggerhook.SetMaxQueues(c.HookMaxBuffer),
		loggerhook.SetLevels(hookLevels...),
	)
	logger.AddHook(h)
	logger.WithContext(context.Background()).Infoln("")
}
