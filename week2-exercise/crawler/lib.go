package crawler

import (
	"errors"
	"net/http"
	"net/url"
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

var (
	ErrorParseNotFound = errors.New("Parser Not Found")
)

func FindParserByUrl(href string) (ICrawler, error) {
	u, err := url.Parse(href)
	if err != nil {
		return nil, err
	}
	switch u.Host {
	case "www.thesaigontimes.vn":
		return CreateSaiGonTimeCrawler(), nil
	case "vietnamnet.vn":
		return CreateVietNamNetCrawler(), nil
	}
	return nil, ErrorParseNotFound
}
