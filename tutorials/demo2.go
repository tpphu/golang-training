package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

type Respone struct {
	Result Result
	Status interface{}
}

type Result struct {
	Product  Product `json:"data"`
	Error    interface{}
	MetaData interface{}
}

type Product struct {
	Id      int
	Name    string
	AdminId int
	Price   int `json:"final_price"`
}

func main() {
	http.HandleFunc("/order", order)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		panic(err)
	}
}
func order(w http.ResponseWriter, req *http.Request) {
	s := fmt.Sprintf("Tong so tien cua 3 san pham: %d", total())
	fmt.Fprintf(w, s)
}

// Nhieu may 100/10
// Distributed system
func total() int {
	urls := []string{
		"https://www.sendo.vn/m/wap_v2/full/san-pham/ao-so-mi-jean-nam-dai-tay-cao-cap-hang-vnxk-31331127?platform=web",
		"https://www.sendo.vn/m/wap_v2/full/san-pham/ao-dui-nam-cao-cap-30157047",
		"https://www.sendo.vn/m/wap_v2/full/san-pham/ao-so-mi-nam-hang-hop-100361413"}

	total := 0
	g, ctx := errgroup.WithContext(context.TODO())
	for _, url := range urls {
		url := url
		g.Go(func() error {
			// Fetch the URL.
			product, err := getProduct(url, ctx)
			if err != nil {
				return err
			}
			total = total + product.Price
			return nil
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
	return total
}

func getProduct(url string, ctx context.Context) (*Product, error) {
	httpClient := http.Client{
		// net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
		Timeout: time.Duration(60 * time.Second),
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("NewRequest:", err)
		return nil, err
	}
	// context deadline exceeded
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	//
	req = req.WithContext(ctx)
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("httpClient.Do:", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := Respone{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("json.Unmarshal:", err)
		return nil, err
	}
	return &res.Result.Product, nil
}
