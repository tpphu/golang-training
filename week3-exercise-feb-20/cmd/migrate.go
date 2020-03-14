package cmd

import (
	"github.com/jinzhu/gorm"
	"github.com/tpphu/week2-exercise/model"
	"github.com/urfave/cli/v2"
)

var Migrate = cli.Command{
	Name:  "migrate",
	Usage: "Migreate schema to database",
	Action: func(c *cli.Context) error {
		db := c.App.Metadata["db"].(*gorm.DB)
		db.DropTableIfExists(&model.Url{}, &model.Article{}, &model.Note{})
		db.AutoMigrate(&model.Url{}, &model.Article{}, &model.Note{})
		return nil
	},
}
