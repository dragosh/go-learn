package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"dragosh.kmade.net/go-learn/pkg/greet"
	"github.com/urfave/cli/v2"
)

var (
	Revision = "fafafaf"
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s revision=%s\n", c.App.Version, Revision)
	}
	app := &cli.App{
		Version: "v0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "Language for the greeting",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "hello",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					name := c.Args().Get(0)
					fmt.Println("CLI:", greet.Hello(name))
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
