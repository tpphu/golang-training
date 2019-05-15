package crawler

import (
	"testing"
)

func Test_Parse_SaiGonTime(t *testing.T) {

	fileDataPath := "./data/thesaigontimes.html"
	url := "https://www.thesaigontimes.vn/274113/bao-giay-van-thu-vi.html"
	resp := generateResponse(fileDataPath, url)

	var SaiGonTime ICrawler = CreateSaiGonTimeCrawler()
	data := SaiGonTime.Parse(resp)

	if data.Title != "Báo giấy vẫn thú vị!" {
		t.Errorf("Title should be expected")
	}
}
