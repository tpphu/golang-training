package cmd

import (
	"github.com/jinzhu/gorm"
	"github.com/tpphu/week2-exercise/service"
	"github.com/urfave/cli/v2"
)

// 1. Load du lieu url tu trong db ra
// 2. Lay cai url crawl du lieu
// 3. Parse cai moi dung html crawl dc ve thanh cau truc
// title. content, published date, author (Article)
// 4. Insert du lieu vao val trong db
var Start = cli.Command{
	Name:  "start",
	Usage: "Start the application",
	Action: func(c *cli.Context) error {
		db := c.App.Metadata["db"].(*gorm.DB)
		svc := service.NewCrawlerService(db)
		// Doc lap chi load urls tu trong db
		urls := svc.Load()
		// Neu ma co urls tu trong db => no se crawl
		articles := svc.Crawl(urls)
		// Neu crawl dc ve content => no insert vao trong db
		svc.InsertArcitlesToDB(articles)
		return nil
	},
}
