package main

import (
	"fmt"
	cli "github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
	"phudt/week2-2/internal/job"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Start application",
				Action: func(c *cli.Context) error {
					fmt.Println("This is main app")
					return nil
				},
			},
			{
				Name:  "migrate",
				Usage: "Migrate application",
				Action: func(c *cli.Context) error {
					fmt.Println("This is migrate app")
					return nil
				},
			},
			{
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "n",
						Usage:       "Mot gia tri nao do ",
						Value:       0,
						DefaultText: "a number",
					},
					&cli.BoolFlag{
						Name:        "x",
						Usage:       "On/Off mot gia tri nao do",
						Value:       false,
						DefaultText: "a bool value",
					},
				},
				Name:   "job",
				Usage:  "Job application",
				Action: job.Job,
			},
		}}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
