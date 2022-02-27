package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	ctx := context.Background()
	app := cli.NewApp()
	app.Name = "cli test"
	app.Commands = []*cli.Command{
		newWebCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("fail err")
	}

}
func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "web run",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "api",
				Usage: "api run",
			},
			&cli.StringFlag{
				Name:  "cms",
				Usage: "cms run",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println(c.String("api"))
		}
	}
}
