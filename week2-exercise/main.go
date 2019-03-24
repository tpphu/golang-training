package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./crawler"
)

var SaiGonTime crawler.ICrawler = crawler.CreateSaiGonTimeCrawler()

func main() {
	url := "https://www.thesaigontimes.vn/274113/bao-giay-van-thu-vi.html"
	resp, _ := http.Get(url)
	data := SaiGonTime.Parse(resp)
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
	// go load()
	// go crawl()
	// go save()
	// go update()
}

func load() {
}

func crawl() {

}

func save() {

}

func update() {

}
