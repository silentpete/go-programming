package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "say a greeting"
	app.Action = func(c *cli.Context) error {
		println("Greetings")
		return nil
	}

	app.Run(os.Args)
}
