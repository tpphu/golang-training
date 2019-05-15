package crawler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func generateResponse(dataPathFile, url string) *http.Response {
	// 1. Create Reponse
	w := httptest.NewRecorder()
	data, _ := ioutil.ReadFile(dataPathFile)
	w.Write(data)
	resp := w.Result()
	// 2. Append Request
	req := httptest.NewRequest("GET", url, nil)
	// 2.1 Cho nay that la chuoi
	resp.Request = req

	return resp
}
