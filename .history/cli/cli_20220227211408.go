package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Run(os.Args)

}
func newWebCmd(ctx context.Context) *cli.Command {

}
