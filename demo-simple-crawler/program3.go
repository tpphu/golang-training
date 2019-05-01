package main

import "github.com/jinzhu/gorm"

type App struct {
	DB *gorm.DB
}

func (self *App) loadUrlsFromDb() {

}

func (self *App) crawlAndExtractArticle() {

}

func (self *App) updateUrlsToDb() {

}

func (self *App) writeArticleToDb() {

}

func program3() {
	app := App{}
	go app.loadUrlsFromDb()
	go app.crawlAndExtractArticle()
	go app.updateUrlsToDb()
	go app.writeArticleToDb()
}
