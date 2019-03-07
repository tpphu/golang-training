package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func getTagContent(doc string, tag string) (string, error) {
	start := strings.Index(doc, "<"+tag+">")
	if start == -1 {
		return "", errors.New("Missing " + tag)
	}
	end := strings.Index(doc, "</"+tag+">")
	if end == -1 {
		return "", errors.New("Missing " + tag)
	}
	title := doc[start+len(tag)+2 : end]
	return strings.TrimSpace(title), nil
}

func getTitle(doc string) (string, error) {
	return getTagContent(doc, "title")
}

func extractDoc(i int, url string) {
	fmt.Println("Index: ", i)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Print("Get url error: ", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	doc := string(body)

	title, _ := getTitle(doc)

	fmt.Println("Title:", title)
}

func main() {
	// var url string
	// url = "https://www.thesaigontimes.vn/274113/bao-giay-van-thu-vi.html"
	// var url string = "https://www.thesaigontimes.vn/274113/bao-giay-van-thu-vi.html"
	arrUrls := []string{
		"https://reqres.in/api/users?delay=3",
		"https://reqres.in/api/users?delay=2",
		"https://www.thesaigontimes.vn/274113/bao-giay-van-thu-vi.html",
	}

	for i := 0; i < len(arrUrls); i++ {
		url := arrUrls[i]
		go extractDoc(i, url)

	}
	time.Sleep(time.Second * 10)
}
