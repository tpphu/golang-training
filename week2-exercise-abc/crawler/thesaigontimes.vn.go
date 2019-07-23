package crawler

import (
	"time"

	"github.com/PuerkitoBio/goquery"
)

type impSaiGonTimes struct {
	Crawler
}

func CreateSaiGonTimesCrawler() ICrawler {
	selector := Selector{
		Title:         "title",
		PublishedDate: "#ctl00_cphContent_lblCreateDate",
		Author:        "#ctl00_cphContent_Lbl_Author",
		Content:       "#ARTICLEVIEW",
	}
	craler := impSaiGonTimes{}
	craler.selector = selector
	craler.parser = createDefaultParser()
	craler.parser.extractPublishDate = extractPublishDateSaiGonTimes
	return craler
}

func extractPublishDateSaiGonTimes(selector string, doc *goquery.Document) time.Time {
	return time.Now()
}
