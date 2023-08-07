package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func RunnerFactory(key string) Runner {
	switch key {
	case "silent":
		return &SilentRunner{}
	default:
		return &SecureRunner{}
	}
}

func exec(fileURL string, silent bool) {

	if len(fileURL) < 1 {
		log.Fatal(fmt.Errorf("no file path provided"))
	}

	runner := RunnerFactory("secure")
	if silent {
		runner = RunnerFactory("silent")
	}

	runner.run(fileURL)
}
func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "silent",
				Aliases: []string{"s"},
				Value:   false,
				Usage:   "Execute silently",
			},
		},
		Action: func(cCtx *cli.Context) error {

			if cCtx.NArg() > 0 {
				fileURL := cCtx.Args().Get(0)
				exec(fileURL, cCtx.Bool("silent"))
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
