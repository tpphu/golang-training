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

var offset int = 0
var DBConnectString = "default:secret@/crawler?charset=utf8&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open("mysql", DBConnectString)
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	db.AutoMigrate(&model.Url{}, &model.Article{})

	done := make(chan bool)

	// Step 1. Load du lieu tu db ra
	urlCrawlChan := load(db)
	// Step 2. Crawl
	articleChan, _ := crawl(db, urlCrawlChan)
	// Step 3. Insert DB | Batch
	save(db, articleChan)
	//update (db, urlUpdatechan)

	fmt.Println("Dong nay se dc in ra truoc tien")
	article := <-articleChan
	fmt.Println("dong nay se in ra thu 3 urlCrawlChan:", article.Title)

	<-done
}

func load(db *gorm.DB) <-chan model.Url {
	urlCrawlChan := make(chan model.Url, 1)
	go func() {
		for {
			fmt.Println("Dong nay se dc in ra thu 2 va bat ki")
			urls := []model.Url{}
			// 1. Lay da cac url can crawl
			// 1.a Khong dc lay nhung cai ma minh da lay ra roi
			// => Khong lam theo offset va limit
			// 1.b Khong dc lay nhung cai may khac no da lay
			// 4. a) Query bao nhieu lan
			// DbReq++
			err := db.Where("state = ? AND status = ?",
				model.UrlStateIdle,
				model.UrlStatusReady).
				Offset(offset).
				Limit(10).
				Find(&urls).
				Error
			if err != nil {
				// 4. b) Cho nay ban can biet no co bao nhieu lan error
				// DbErr++
				time.Sleep(time.Second * 3)
				continue
			}
			// 2. Lay du lieu crawl va push vao channel
			for _, url := range urls {
				urlCrawlChan <- url
			}
			// 3. (issues 1.b) minh mark no la state = RUNING (2) de nhung thang khac
			// no khong lay
			offset += 10
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
			//1. Can chay download url thanh cac goroutines
			// de tang toc viec crawl
			go func(url model.Url) {
				var err error
				// 1. Download cai noi dung ve
				resp, err := http.Get(url.Url)
				if err != nil {
					fmt.Println("Download | error:", err)
					return
				}
				if resp == nil {
					fmt.Println("Download | error:")
					return
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
			// 2. Can limit so luong request ra ngoai
			// if watcher.NumHTTPReq - watcher.NumHTTPReq > 100
			// tam thoi ngung download nua
		}
	}()
	return articleChan, urlUpdateChan
}

func save(db *gorm.DB, articleChan <-chan model.Article) {
	go func() {
		articles := []model.Article{}
		for {
			select {
			// 1. Logic nay dam bao nhieu co nhieu hon n articles thi phai insert
			case article := <-articleChan:
				articles = append(articles, article)
				if len(articles) >= 5 {
					insertArticles(db, articles)
					articles = []model.Article{}
				}
			// 2. Hoac neu sau thoi gian t giay, du chi co <n articles se van phai insert
			case <-time.After(3 * time.Second):
				if len(articles) > 0 {
					insertArticles(db, articles)
					articles = []model.Article{}
				}
			}
		}
	}()
}

func insertArticles(db *gorm.DB, articles []model.Article) {
	err := helper.BatchInsert(db, articles)
	if err != nil {
		fmt.Println("insertArticles | err:", err)
	}
}
