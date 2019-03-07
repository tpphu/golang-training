package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Selector struct {
	title         string
	publishedDate string
	author        string
	content       string
}

type Data struct {
	Title         string
	PublishedDate time.Time
	Author        string
	Content       string
}

type ICrawler interface {
	Parse(res *http.Response) Data
}

type Crawler struct {
	domain   string
	data     Data
	selector Selector
}

func (self *Crawler) Parse(res *http.Response) Data {
	doc, _ := goquery.NewDocumentFromResponse(res)
	// Title
	self.data.Title = strings.TrimSpace(self.extract(self.selector.title, doc))
	// Author
	self.data.Author = strings.TrimSpace(self.extract(self.selector.author, doc))
	// Content
	self.data.Content = strings.TrimSpace(self.extract(self.selector.content, doc))
	// PublishDate
	publishedDate := strings.TrimSpace(self.extract(self.selector.publishedDate, doc))
	fmt.Println("publishedDate:", publishedDate)
	publishedDate = strings.Replace(publishedDate, ",", "", -1)
	r, _ := regexp.Compile("[0-9].+$")
	publishedDate = r.FindString(publishedDate)
	fmt.Println("publishedDate:", publishedDate)
	r, _ = regexp.Compile("[^0-9/:]+")
	publishedDate = string(r.ReplaceAll([]byte(publishedDate), []byte(" ")))
	self.data.PublishedDate, _ = time.Parse("02/1/2006 15:04", publishedDate)
	return self.data
}

func (self *Crawler) extract(selector string, doc *goquery.Document) string {
	if selector == "" {
		return ""
	}
	return doc.Find(selector).Text()
}

func NewCrawler(domain string, selector Selector) ICrawler {
	return &Crawler{
		selector: selector,
		data:     Data{},
	}
}

var KinhTeMoiDoc ICrawler = NewCrawler("www.thesaigontimes.vn", Selector{
	title:         "title",
	publishedDate: "#ctl00_cphContent_lblCreateDate",
	author:        "#ctl00_cphContent_Lbl_Author",
	content:       ".SGTOSummary",
})

var VietNamnetDoc ICrawler = NewCrawler("vietnamnet.vn", Selector{
	title:         ".title.f-22.c-3e",
	publishedDate: ".ArticleDate  .right",
	author:        "#ArticleContent p span.bold",
	content:       "#ArticleContent .ArticleLead",
})

func main() {
	url := "https://www.thesaigontimes.vn/274113/bao-giay-van-thu-vi.html"
	resp, _ := http.Get(url)
	data := KinhTeMoiDoc.Parse(resp)
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
	// go load()
	// go crawl()
	// go save()
	// go update()
}

func load() {
}

func crawl() {

}

func save() {

}

func update() {

}
