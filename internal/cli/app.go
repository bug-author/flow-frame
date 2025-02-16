package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func CreateCLIApp() *cli.App {
	app := &cli.App{
		Name:  "flow-frame",
		Usage: "TickTick",
		Action: func(cCtx *cli.Context) error {
			userCommand := cCtx.Args().Get(0)
			fmt.Println("userCommand:", userCommand)
			return nil
		},
	}

	return app
}
