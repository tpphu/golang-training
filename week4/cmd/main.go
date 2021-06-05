package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "An application ....",
		Usage: "To demo ...",
		Commands: []*cli.Command{&cli.Command{
			Name:  "serve",
			Usage: "Start http server",
			Flags: []cli.Flag{&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "HTTP port to listen",
				EnvVars: []string{
					"HTTP_PORT",
				},
				DefaultText: "8080",
			}},
			Action: ServeAction,
		}, &cli.Command{
			Name:  "subscribe",
			Usage: "Start a job",
			Action: func(*cli.Context) error {
				fmt.Println("start a job")
				return nil
			},
		}},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
