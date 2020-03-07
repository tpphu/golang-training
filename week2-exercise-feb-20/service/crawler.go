package service

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tpphu/week2-exercise/model"
)

type Crawler struct {
	db *gorm.DB
}

func NewCrawlerService(db *gorm.DB) Crawler {
	svc := Crawler{
		db: db,
	}
	return svc
}

// Ham nay cuc ki kho
// 1. Load ra tung phan nhu the nao
// 2. chay tren nhieu may thi sao
func (svc Crawler) Load() chan model.Url {
	urlsChan := make(chan model.Url, 10)
	go func() {
		for {
			urls := []model.Url{}
			svc.db.Find(&urls).Limit(5)
			for i := 0; i < len(urls); i++ {
				urlsChan <- urls[i]
			}
			time.Sleep(1 * time.Second)
			// Nhung van de khac minh se thao luan tren lop
			// Lieu cu 1s ma put 5 cai urls vao channel thi dung khong
		}
	}()

	return urlsChan
}

func (svc Crawler) Crawl(urls chan model.Url) chan model.Article {
	articlesChan := make(chan model.Article, 10)
	go func() {
		for {
			// Lay ra mot cai
			url := <-urls
			// Crawl
			// Code rat nhieu o day
			// Download html ve parse ra noi
			//
			artile := model.Article{
				URLID:   url.ID,
				Title:   "xxx",
				Content: "yyyy",
			}
			articlesChan <- artile
		}
	}()

	return articlesChan
}

// Insert tung bai thi no ton connection
// Muon lam the nao de gom lai 10 bai insert 1 lan
func (svc Crawler) InsertArcitlesToDB(articles chan model.Article) {
	for {
		article := <-articles
		// Insert vao trong db
		svc.db.Create(&article)
		time.Sleep(1 * time.Second)
	}
}
