package main

import (
	appa "thrgoweb/internal/app"
	"thrgoweb/internal/app/config"
)

func main() {
	config.MustLoad("configs/config.toml")

	appa.Run()
}
