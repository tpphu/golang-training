package crawler

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Selector struct {
	Title         string
	PublishedDate string
	Author        string
	Content       string
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
	selector Selector
	parser   Parser
}

func (self *Crawler) Parse(res *http.Response) Data {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		panic(err)
	}
	data := Data{}
	data.Title = self.parser.extractTitle(self.selector.Title, doc)
	data.Author = self.parser.extractAuthor(self.selector.Author, doc)
	data.Content = self.parser.extractContent(self.selector.Content, doc)
	data.PublishedDate = self.parser.extractPublishDate(self.selector.PublishedDate, doc)
	return data
}
