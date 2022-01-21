package app

import "fmt"

func Run() {
	injector, _, err := BuildInjector()
	if err != nil {
		fmt.Println(err)
	}
	injector.Engine.Run()
}
