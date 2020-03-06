package app

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tpphu/golang-trainning/model"
)

type Application struct {
	db *gorm.DB
}

func NewApp(db *gorm.DB) Application {
	return Application{
		db: db,
	}
}

func (app Application) Load() chan model.Url {
	urlChan := make(chan model.Url, 10)
	go func() {
		for {
			urls := []model.Url{}
			err := app.db.Where("state = ?", model.UrlStateIdle).
				Limit(5).
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

func (app Application) Crawl(chan model.Url) chan model.Article {
	articleChan := make(chan model.Article, 10)
	go func() {
		for {

		}
	}()
	return articleChan
}
