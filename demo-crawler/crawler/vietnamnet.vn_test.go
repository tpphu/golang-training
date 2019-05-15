package crawler

import (
	"testing"
)

func Test_Parse_VietNamNet(t *testing.T) {

	fileDataPath := "./data/vietnamnet.html"
	url := "https://vietnamnet.vn/vn/cong-nghe/ung-dung/cach-su-dung-google-maps-de-giam-sat-vi-tri-cua-tre-nho-514378.html"
	resp := generateResponse(fileDataPath, url)

	var SaiGonTime ICrawler = CreateSaiGonTimeCrawler()
	data := SaiGonTime.Parse(resp)

	if data.Title != "Cách sử dụng Google Maps để giám sát vị trí của trẻ nhỏ" {
		t.Errorf("Title should be expected, actual is %s", data.Title)
	}
}
