package crawler

import (
	"testing"
)

func Test_Parse_SaiGonTimes(t *testing.T) {
	url := "https://www.thesaigontimes.vn/291821/dia-phuong-nao-xay-nhieu-khach-san-cao-cap-nhat-.html"
	res := generateMockReponse("./dataset/saigontimes_url_1.html", url)

	var SaiGonTime ICrawler = CreateSaiGonTimesCrawler()
	data := SaiGonTime.Parse(res)
	if data.Title != "Địa phương nào xây nhiều khách sạn cao cấp nhất?" {
		t.Errorf("Title should be expected")
	}
}
