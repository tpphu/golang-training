package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error occurs when load enviroment, %s", err.Error())
	}
	app := &cli.App{
		Name:    "spider",
		Usage:   "Crawler Service",
		Version: "1.0.0",
		Action: func(c *cli.Context) error {
			fmt.Println("Welcome to Crawler Service. To see how to use, type -h for more detail")
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Application's port is listen",
				EnvVars: []string{"PORT"},
				Value:   3000,
			},
			&cli.StringFlag{
				Name:    "database",
				Aliases: []string{"db"},
				Usage:   "Database URI to connect to DB",
				EnvVars: []string{"DB_URI"},
			},
			&cli.BoolFlag{
				Name:    "dblog",
				Usage:   "Application's port is listen",
				EnvVars: []string{"DB_LOG"},
				Value:   false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"run"},
				Usage:   "Start the service",
				Action: func(c *cli.Context) error {
					fmt.Println("Start the service")
					return nil
				},
			},
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "Migrate schema to DB",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "seed",
				Aliases: []string{"s"},
				Usage:   "Seed data to DB",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
