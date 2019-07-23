package crawler

import (
	"time"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type ExtractString = func (selector string, doc *goquery.Document) string
type ExtractTime = func (selector string, doc *goquery.Document) time.Time
type Parser struct {
	extractTitle ExtractString
	extractContent ExtractString
	extractAuthor ExtractString
	extractPublishDate ExtractTime
}

func extractSimpleString(selector string, doc *goquery.Document) string {
	text := doc.Find(selector).Text()
	return strings.TrimSpace(text)
}

func extractNotYetImplement(selector string, doc *goquery.Document) time.Time{
	panic("Should implement at instance struct")
}

func createDefaultParser() Parser{
	return Parser{
		extractTitle: extractSimpleString,
		extractContent: extractSimpleString,
		extractAuthor: extractSimpleString,
		extractPublishDate: extractNotYetImplement,
	}
}

