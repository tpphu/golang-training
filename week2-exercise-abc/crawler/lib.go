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

func (c Crawler) Parse(res *http.Response) Data {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	data := Data{}
	if err != nil {
		return data
	}
	data.Title = c.parser.extractTitle(c.selector.Title, doc)
	data.Content = c.parser.extractContent(c.selector.Content, doc)
	data.Author = c.parser.extractTitle(c.selector.Author, doc)
	data.PublishedDate = c.parser.extractPublishDate(c.selector.PublishedDate, doc)
	return data
}
