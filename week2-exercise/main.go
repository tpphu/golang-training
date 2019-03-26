package main

import (
	"fmt"
	"net/http"
	"time"

	"./crawler"
	"./helper"
	"./model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var SaiGonTime crawler.ICrawler = crawler.CreateSaiGonTimeCrawler()
var VietNamNet crawler.ICrawler = crawler.CreateVietNamNetCrawler()
var DBConnectString = "default:secret@/crawler?charset=utf8&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open("mysql", DBConnectString)
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	db.AutoMigrate(&model.Url{}, &model.Article{})

	urlCrawlChan := load(db)
	articleChan, urlUpdateChan := crawl(db, urlCrawlChan)
	save(db, articleChan)
	update(db, urlUpdateChan)

	for {
		time.Sleep(3 * time.Second)
	}
}

func load(db *gorm.DB) <-chan model.Url {
	urlCrawlChan := make(chan model.Url, 10)
	go func() {
		for {
			urls := []model.Url{}
			err := db.Where("state = ? AND status = ?",
				model.UrlStateIdle,
				model.UrlStatusReady).
				Find(&urls).
				Error
			if err != nil {
				fmt.Println(err)
			}
			for _, url := range urls {
				urlCrawlChan <- url
			}
			time.Sleep(time.Second * 3)
		}
	}()
	return urlCrawlChan
}

func crawl(db *gorm.DB, urlCrawlChan <-chan model.Url) (<-chan model.Article, <-chan model.Url) {
	articleChan := make(chan model.Article, 10)
	urlUpdateChan := make(chan model.Url, 10)
	go func() {
		for url := range urlCrawlChan {
			var err error
			resp, err := http.Get(url.Url)
			if err != nil {
				fmt.Println("Download | error:", err)
			}
			parser, err := crawler.FindParserByUrl(url.Url)
			if err != nil {
				fmt.Println("Find parser | error:", err)
			}
			data := parser.Parse(resp)
			article := model.Article{
				UrlID: url.ID,
			}
			helper.FillDataToArticle(&article, data)
			articleChan <- article
			urlUpdateChan <- url
		}
	}()
	return articleChan, urlUpdateChan
}

func save(db *gorm.DB, articleChan <-chan model.Article) {
	go func() {
		for article := range articleChan {
			db.Create(&article)
		}
	}()
}

func update(db *gorm.DB, urlUpdateChan <-chan model.Url) {
	go func() {
		for url := range urlUpdateChan {
			url.State = model.UrlStateRunning
			url.Status = model.UrlStatusSuccess
			db.Save(&url)
		}
	}()
}
