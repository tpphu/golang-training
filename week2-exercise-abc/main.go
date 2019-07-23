package main

import (
	"fmt"
	"time"

	"./model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Step 1 - Define Model, Connect Db, Migrate DB
// Step 2 - Function de ma doc tu DB ra cac URL
// Step 3 - Crawl noi dung (html string)
// Step 4 - Parse cai Article
// Step 5 - Insert vo DB

func main() {
	db, err := gorm.Open("mysql", "default:secret@/crawler_abc?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.Url{}, &model.Article{})

	// urlChan := load(db)
	// articleChan := crawl(urlChan)
	fmt.Println("Hello world!")
}

func load(db *gorm.DB) <-chan model.Url {
	urlChan := make(chan model.Url, 10)
	go func() {
		for {
			urls := []model.Url{}
			err := db.Where("state = ?", model.UrlStateIdle).
				Limit(10).
				Find(&urls).Error
			if err != nil {
				time.Sleep(30 * time.Second)
				continue
			}
			for i := 0; i < len(urls); i++ {
				urlChan <- urls[i]
			}
			time.Sleep(1 * time.Second)
		}
	}()
	return urlChan
}

func crawl(urlChan <-chan model.Url) <-chan model.Article {
	articleChan := make(chan model.Article, 10)
	return articleChan
}
