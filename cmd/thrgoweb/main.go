package main

import (
	"context"
	"thrgoweb/pkg/logger"
)

func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")
	cxt := logger.NewUserNameContext(context.Background(), "aaa")
	// config.MustLoad("configs/config.toml")
	logger.WithContext(ctx).Infof("dddddd")
	logger.WithContext(cxt).Infof("dddddd")
	// appa.Run()
}
