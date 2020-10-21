package command

import (
	"fmt"
	"sync"

	"github.com/tpphu/golang-training/week2/config"
	"github.com/tpphu/golang-training/week2/provider"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	// Name string
	Price uint
}

var InfoCommand = cli.Command{
	Name:  "info",
	Usage: "run the service",
	Action: func(c *cli.Context) error {
		fmt.Println("Show info")
		return nil
	},
}

var ServeCommand = cli.Command{
	Name:  "serve",
	Usage: "run the service",
	Action: func(c *cli.Context) error {
		conf := config.NewDefaultConfig()
		rp := provider.MustBuildRP(conf)
		// Create
		wg := sync.WaitGroup{}
		for j := 0; j < 5; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 2000; i++ {
					rp.DB.Create(&Product{Code: fmt.Sprintf("D%d", i), Price: 100})
				}
			}()
		}
		wg.Wait()
		return nil
	},
}

var MigrateCommand = cli.Command{
	Name:  "migrate",
	Usage: "use to migrate database",
	Action: func(c *cli.Context) error {
		conf := config.NewDefaultConfig()
		rp := provider.MustBuildRP(conf)
		// Only for development
		// Migrate the schema
		rp.DB.AutoMigrate(&Product{})
		return nil
	},
}
