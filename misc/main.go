package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	cli "github.com/urfave/cli/v2"
)

func main() {
	// Tien xu ly file .env thanh bien moi truong ENV
	// Hoac la minh load tu file .evn
	// Hoac la minh load tu file YAML
	app := &cli.App{
		Action: func(c *cli.Context) error {

			host := c.String("host")
			user := c.String("user")
			passwd := c.String("passwd")
			port := c.String("port")
			dbname := c.String("dbname")

			fmt.Println(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, host, port, dbname))

			dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, host, port, dbname)
			_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

			return err
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "host",
				Aliases:  []string{"H"},
				Value:    "127.0.0.1",
				Usage:    "ip of database",
				EnvVars:  []string{"MYSQL_HOST"},
				FilePath: `.env`,
			},
			&cli.StringFlag{
				Name:    "user",
				Aliases: []string{"u"},
				Value:   "root",
				Usage:   "user of database",
				EnvVars: []string{"MYSQL_USER"},
			},
			&cli.StringFlag{
				Name:    "passwd",
				Aliases: []string{"P"},
				Value:   "root",
				Usage:   "password of database",
				EnvVars: []string{"MYSQL_PASSWD"},
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   "3306",
				Usage:   "port of database",
				EnvVars: []string{"MYSQL_PORT"},
			},
			&cli.StringFlag{
				Name:    "dbname",
				Aliases: []string{"db"},
				Value:   "user",
				Usage:   "db name of database",
				EnvVars: []string{"MYSQL_DB_NAME"},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
