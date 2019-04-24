package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	refactor1Version()
}

func simpleVersion() {
	fmt.Println("This is a simple crawler in Go")
	url := "https://www.thesaigontimes.vn/274113/Bao-giay-van-thu-vi!.html"

	res, _ := http.Get(url)

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	title := doc.Find("title").Text()
	author := doc.Find("#ctl00_cphContent_Lbl_Author").Text()

	fmt.Println("title:", title)
	fmt.Println("author:", author)
}

func refactor1Version() {
	url := "https://www.thesaigontimes.vn/274113/Bao-giay-van-thu-vi!.html"
	res, _ := http.Get(url)
	parser := findParserByUrl(url)
	data := parser.Parse(res)
	fmt.Println("title:", data.Title)
	fmt.Println("author:", data.Author)
}
