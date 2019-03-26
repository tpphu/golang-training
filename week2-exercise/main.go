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

	watcher := &helper.Watcher{}
	urlCrawlChan := load(db, watcher)
	articleChan, urlUpdateChan := crawl(db, watcher, urlCrawlChan)
	save(db, watcher, articleChan)
	update(db, watcher, urlUpdateChan)

	for {
		watcher.Out()
		time.Sleep(1 * time.Second)
	}
}

func load(db *gorm.DB, watcher *helper.Watcher) <-chan model.Url {
	urlCrawlChan := make(chan model.Url, 10)
	go func() {
		for {
			urls := []model.Url{}
			watcher.DBLoadUrlReq++
			err := db.Where("state = ? AND status = ?",
				model.UrlStateIdle,
				model.UrlStatusReady).
				Limit(1).
				Find(&urls).
				Error

			watcher.DBLoadUrlRes++
			if err != nil {
				watcher.DBLoadUrlErr++
				fmt.Println(err)
			}
			watcher.DBLoadUrlTotal += len(urls)
			for _, url := range urls {
				urlCrawlChan <- url
			}
			time.Sleep(time.Second * 3)
		}
	}()
	return urlCrawlChan
}

func crawl(db *gorm.DB, watcher *helper.Watcher, urlCrawlChan <-chan model.Url) (<-chan model.Article, <-chan model.Url) {
	articleChan := make(chan model.Article, 10)
	urlUpdateChan := make(chan model.Url, 10)
	go func() {
		for url := range urlCrawlChan {
			var err error
			watcher.NumHTTPReq++
			resp, err := http.Get(url.Url)
			watcher.NumHTTPRes++
			if err != nil {
				watcher.NumHTTPErr++
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

func save(db *gorm.DB, watcher *helper.Watcher, articleChan <-chan model.Article) {
	go func() {
		for article := range articleChan {
			watcher.DBInsArticleReq++
			err := db.Create(&article).Error
			watcher.DBInsArticleRes++
			if err != nil {
				watcher.DBInsArticleErr++
			}
		}
	}()
}

func update(db *gorm.DB, watcher *helper.Watcher, urlUpdateChan <-chan model.Url) {
	go func() {
		for url := range urlUpdateChan {
			url.State = model.UrlStateRunning
			url.Status = model.UrlStatusSuccess
			db.Save(&url)
		}
	}()
}
