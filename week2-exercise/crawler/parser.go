package crawler

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ExtractString = func (selector string, doc *goquery.Document) string
type ExtractTime = func (selector string, doc *goquery.Document) time.Time

type Parser struct {
	extractPublishDate ExtractTime
	extractContent ExtractString
	extractAuthor ExtractString
	extractTitle ExtractString
}

func createDefaultParser() Parser {
	parser := Parser{
		extractPublishDate: extractPublishDate,
		extractContent: extractContent,
		extractAuthor: extractAuthor,
		extractTitle: extractTitle,
	}
	return parser
}

func extractPublishDate(selector string, doc *goquery.Document) time.Time {
	panic("Should implement")
}

func extractContent(selector string, doc *goquery.Document) string {
	value := extract(selector, doc)
	return strings.TrimSpace(value)
}

func extractAuthor(selector string, doc *goquery.Document) string {
	value := extract(selector, doc)
	return strings.TrimSpace(value)
}

func extractTitle(selector string, doc *goquery.Document) string {
	value := extract(selector, doc)
	return strings.TrimSpace(value)
}

func extract(selector string, doc *goquery.Document) string {
	if selector == "" {
		return ""
	}
	return doc.Find(selector).Text()
}
