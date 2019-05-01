package main

import (
	"fmt"
	"net/http"
)

func refactor1Version() {
	url := "https://www.thesaigontimes.vn/274113/Bao-giay-van-thu-vi!.html"
	res, _ := http.Get(url)
	parser := findParserByUrl(url)
	data := parser.Parse(res)
	fmt.Println("title:", data.Title)
	fmt.Println("author:", data.Author)
}
