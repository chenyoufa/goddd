package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "gin-admin"
	app.Version = "v1"
	app.Usage = "hello world"
	app.Commands = []*cli.Command{}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
func newWebCmd() {
	return &cli.Command{
		Name:  "web",
		Usage: "Run http server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "app configuration",
				Required: true,
			}
		},
	}
}
