package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Parser struct {
}

func findParserByUrl(url string) Parser {
	parser := Parser{}
	return parser
}

func (self *Parser) Parse(res *http.Response) Data {
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	title := doc.Find("title").Text()
	author := doc.Find("#ctl00_cphContent_Lbl_Author").Text()

	data := Data{}
	data.Title = title
	data.Author = author
	return data
}
