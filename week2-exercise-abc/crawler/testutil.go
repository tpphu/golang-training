package crawler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func generateMockReponse(file string, url string) *http.Response {
	w := httptest.NewRecorder()
	data, _ := ioutil.ReadFile(file)
	w.Write(data)
	res := w.Result()
	req := httptest.NewRequest("GET", url, nil)
	res.Request = req
	return res
}
