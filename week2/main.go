package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	// Timeout context
	// url := "https://vnexpress.net/pho-thu-tuong-chua-can-gian-cach-xa-hoi-ca-nuoc-4276285.html"
	url := "https://reqres.in/api/users?delay=5"
	data, err := download2(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func getDefaultClient() *http.Client {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	return netClient
}

func download2(url string) (string, error) {
	client := getDefaultClient()
	// Dung cai function cancel
	// De force timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := client.Do(req)
	// Error se vao day
	if err != nil {
		fmt.Println(" Chi toi day thoi")
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body[0:150]), nil
}

func download(url string, ch chan string) {
	client := getDefaultClient()
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println("url:", url)
	// Thi cho nay khong the push them vao dc.
	ch <- string(body[0:150])
	// fmt.Println("url:", url)
}

// func fun1() {
// 	for {
// 		x := getABC()
// 		fmt.Println(x, "1")
// 		fmt.Println(x,"2")
// 		fmt.Println(x,"3")
// 		fmt.Println(x,"4")
// 		fmt.Println(x,"5")
// 		fmt.Println(x,"6")
// 		fmt.Println(x,"7")
// 		fmt.Println(x,"8")
// 		fmt.Println(x,"9")
// 		fmt.Println(x,"10")
// 	}
// }

func goprocs() {
	for i := 0; i < 10000; i++ {
		go func() {
			for j := 0; j < 1000000000; j++ {
				_ = j
			}
		}()
	}
	time.Sleep(5 * time.Minute)
}

func go_and_chan() {
	urls := []string{
		"https://vnexpress.net/pho-thu-tuong-chua-can-gian-cach-xa-hoi-ca-nuoc-4276285.html",
		"https://vnexpress.net/quoc-lo-51-con-duong-am-anh-tai-xe-4275845.html",
		"https://vnexpress.net/them-30-ca-covid-19-4276501.html",
		"https://vnexpress.net/6-toa-metro-so-1-ve-toi-sai-gon-4275344.html",
		"https://vnexpress.net/cao-toc-hon-12-000-ty-dong-vang-bong-xe-4270527.html",
	}
	// Queue
	// Capacity channel la co limit
	// Cappacity = empty => Unbufferred channel => no chi chua toi da 1 phan tu
	// La mot con so != empty => Bufferred channel=> co suc chua cu the > 1
	ch := make(chan string, 3)
	for i := 0; i < len(urls); i++ {
		// 1. Co nghia chay dong thoi
		go download(urls[i], ch)
	}
	// Neu ma khong co logic de pull ra
	content1 := <-ch
	fmt.Println("Content1:", content1)
	content2 := <-ch
	fmt.Println("Content2:", content2)
	// content3 := <-ch
	// fmt.Println("Content3:", content3)
	// content4 := <-ch
	// fmt.Println("Content4:", content4)
	// content5 := <-ch
	// fmt.Println("Content5:", content5)
	time.Sleep(5 * time.Minute)
}
