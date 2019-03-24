package crawler

import (
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type impSaiGonTime struct {
	Crawler
}

var extractPublishDateSaiGonTime = func(selector string, doc *goquery.Document) time.Time {
	publishedDateStr := strings.TrimSpace(extract(selector, doc))
	publishedDateStr = strings.Replace(publishedDateStr, ",", "", -1)
	r, _ := regexp.Compile("[0-9].+$")
	publishedDateStr = r.FindString(publishedDateStr)
	r, _ = regexp.Compile("[^0-9/:]+")
	publishedDateStr = string(r.ReplaceAll([]byte(publishedDateStr), []byte(" ")))
	publishedDate, _ := time.Parse("02/1/2006 15:04", publishedDateStr)
	return publishedDate
}

func CreateSaiGonTimeCrawler() ICrawler {
	selector := Selector{
		Title:         "title",
		PublishedDate: "#ctl00_cphContent_lblCreateDate",
		Author:        "#ctl00_cphContent_Lbl_Author",
		Content:       ".SGTOSummary",
	}
	crawler := &impSaiGonTime{}
	crawler.selector = selector
	crawler.parser = createDefaultParser()
	crawler.parser.extractPublishDate = extractPublishDateSaiGonTime
	return crawler
}
