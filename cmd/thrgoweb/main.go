package main

import (
	"context"
	appa "thrgoweb/internal/app"
	"thrgoweb/internal/app/config"
	"thrgoweb/pkg/logger"
)

func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")
	config.MustLoad("configs/config.toml")
	logger.WithContext(ctx).Infof("dddddd")
	appa.Run()
}
