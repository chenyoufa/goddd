package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("fail err")
	}

}
func newWebCmd(ctx context.Context) *cli.Command {

}
