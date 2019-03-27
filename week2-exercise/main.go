package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"./crawler"
	"./helper"
	"./model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var SaiGonTime crawler.ICrawler = crawler.CreateSaiGonTimeCrawler()
// var VietNamNet crawler.ICrawler = crawler.CreateVietNamNetCrawler()
var DBConnectString = "default:secret@/crawler?charset=utf8&parseTime=True&loc=Local"
var totalInstance int = 1
var nthInstance int = 0

func main() {
	// Connect vo DB MySQL
	db, err := gorm.Open("mysql", DBConnectString)
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	db.AutoMigrate(&model.Url{}, &model.Article{})
	// go run main.go 2 0
	// go run main.go 2 1
	args := os.Args[1:]
	totalInstance, _ = strconv.Atoi(args[0])
	nthInstance, _ = strconv.Atoi(args[1])
	// Tao mot cai watcher
	watcher := &helper.Watcher{}
	// Chuong trinh chinh cua minh o day
	// Step 1: Load urls need to crawl from [BD]
	urlCrawlChan := load(db, watcher) // I/O : DB
	// Step 2: download and parse/extract article from url: I/O: HTTP
	articleChan, urlUpdateChan := crawl(db, watcher, urlCrawlChan)
	// Step 3: Save lai nhung article vao db
	save(db, watcher, articleChan)
	// Step 4: La danh dau nhung cai url minh da crawl xong
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
			// 1. Lay da cac url can crawl
			err := db.Where("(id % ? = ? AND id > ?) AND state = ? AND status = ?",
				totalInstance,
				nthInstance,
				watcher.DBLoadUrlLastId,
				model.UrlStateIdle,
				model.UrlStatusReady).
				Limit(10).
				Find(&urls).
				Error

			watcher.DBLoadUrlRes++
			if err != nil {
				watcher.DBLoadUrlErr++
				fmt.Println(err)
			}
			watcher.DBLoadUrlTotal += len(urls)
			// 2. Lay du lieu crawl va push vao channel
			for _, url := range urls {
				urlCrawlChan <- url
				if url.ID > watcher.DBLoadUrlLastId {
					// Co the ket hop voi UpdatedAt de ra mot solution
					// Ve viec crawl lai nhung url bi loi
					// Co the ket hop voi them mot thong so ve so lan bi loi
					watcher.DBLoadUrlLastId = url.ID
				}
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
			go func(url model.Url) {
				var err error
				watcher.NumHTTPReq++
				// 1. Download cai noi dung ve
				resp, err := http.Get(url.Url)
				watcher.NumHTTPRes++
				if err != nil {
					watcher.NumHTTPErr++
					fmt.Println("Download | error:", err)
				}
				// 2. Tim cai parse tuong ung cho cai url
				parser, err := crawler.FindParserByUrl(url.Url)
				if err != nil {
					fmt.Println("Find parser | error:", err)
				}
				// 3. Parse cai noi dung
				data := parser.Parse(resp)
				// 4. Dem cai noi dung va gan vao article
				article := model.Article{
					UrlID: url.ID,
				}
				helper.FillDataToArticle(&article, data) // huong function
				// article.Fill(data) //@deprecated //oop
				// 5. Day du lieu vao channel
				articleChan <- article
				urlUpdateChan <- url
			}(url)
		}
	}()
	return articleChan, urlUpdateChan
}

func save(db *gorm.DB, watcher *helper.Watcher, articleChan <-chan model.Article) {
	go func() {
		for article := range articleChan {
			watcher.DBInsArticleReq++
			// Insert du lieu vao db
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
			url.State = model.UrlStateRunning //Idle
			url.Status = model.UrlStatusSuccess
			db.Save(&url)
		}
	}()
}
