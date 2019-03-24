package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./crawler"
)

var SaiGonTime crawler.ICrawler = crawler.CreateSaiGonTimeCrawler()
var VietNamNet crawler.ICrawler = crawler.CreateVietNamNetCrawler()

func main() {
	// url := "https://www.thesaigontimes.vn/274113/bao-giay-van-thu-vi.html"
	// resp, _ := http.Get(url)
	// data := SaiGonTime.Parse(resp)
	// b, _ := json.Marshal(data)
	// fmt.Println(string(b))

	url := "https://vietnamnet.vn/vn/cong-nghe/ung-dung/cach-su-dung-google-maps-de-giam-sat-vi-tri-cua-tre-nho-514378.html"
	resp, _ := http.Get(url)
	data := VietNamNet.Parse(resp)
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
