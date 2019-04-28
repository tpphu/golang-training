package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DBConnectString1 = "default:secret@/crawler_demo?charset=utf8&parseTime=True&loc=Local"

func simpleVersion() {
	fmt.Println("This is a simple crawler in Go")
	urls := []string{
		"https://www.thesaigontimes.vn/274113/Bao-giay-van-thu-vi!.html",
		"https://www.thesaigontimes.vn/287611/nhan-luc-co-chat-luong-khong-phai-de-xuat-khau-.html",
		"https://www.thesaigontimes.vn/287695/viet-nam-co-6-loai-trai-cay-duoc-phep-xuat-khau-vao-my.html",
		"https://www.thesaigontimes.vn/288018/ngan-hang-vao-mua-chia-co-tuc-giay-.html",
	}

	db, err := gorm.Open("mysql", DBConnectString1)
	if err != nil {
		fmt.Println(err)
		panic("err")
	}
	db.AutoMigrate(&SimpleData{})

	for i := 0; i < len(urls); i++ {
		url := urls[i]
		go crawl1(db, url)
	}

	time.Sleep(time.Second * 30)
}

func crawl1(db *gorm.DB, url string) {

	fmt.Println("--------------")
	res, _ := http.Get(url)

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	title := doc.Find("#ctl00_cphContent_lblTitleHtml").Text()
	author := doc.Find("#ctl00_cphContent_Lbl_Author").Text()
	publishDate := doc.Find("#ctl00_cphContent_lblCreateDate").Text()

	article := SimpleData{
		Title:       title,
		Author:      author,
		PublishDate: publishDate,
	}
	db.Create(&article)
	// fmt.Println("title:", title)
	// fmt.Println("author:", author)
	// fmt.Println("publishDate:", publishDate)

}
