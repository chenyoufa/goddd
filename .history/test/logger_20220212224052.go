package test

import (
	"context"
	"fmt"
	"gdemo1/logger"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	dialector := mysql.Open("fage:Fage501526~@/mytest?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(dialector)
	if err != nil {
		fmt.Println(err)
	}

	var hookLevels []logrus.Level

	hookLevels = append(hookLevels, 0, 1, 2, 3, 4, 5, 6)

	h := loggerhook.New(loggergormhook.New(db),
		loggerhook.SetMaxQueues(1),
		loggerhook.SetMaxWorks(1),
		loggerhook.SetLevels(hookLevels...),
	)
	fmt.Println(h.Levels())
	logger.AddHook(h)

	ctx1 := logger.NewTagContext(context.Background(), "__main__")
	logger.WithContext(ctx1).Info("fff")
	time.Sleep(time.Second * 10)
	fmt.Println("end")
}
