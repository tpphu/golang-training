package handler

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	mock "../mock"
	"../model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func fakeString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func buildMockContext(method string, path string, data string) *gin.Context {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(method, path, strings.NewReader(data))
	ctx.Request.Header.Set("Content-Type", "application/json")
	return ctx
}
func Test_NoteCreate_TitleIsEmpty(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1. Chuan bi input dau vao cho ham CreateNote
	data := `{"title": "","completed": false}`
	ctx := buildMockContext("POST", "/note", data)
	noteRepo := new(mock.NoteRepoImpl)
	// 2. Goi function can test
	_, err := NoteCreate(ctx, noteRepo)
	// 3. Kiem tra ket qua la dung nhu mong doi
	if err == nil {
		t.Fail()
	}
}
func Test_NoteCreate_TitleHasMinLength(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1. Chuan bi input dau vao cho ham CreateNote
	data := `{"title": "abc","completed": false}`
	ctx := buildMockContext("POST", "/note", data)
	noteRepo := new(mock.NoteRepoImpl)
	// 2. Goi function can test
	_, err := NoteCreate(ctx, noteRepo)
	// 3. Kiem tra ket qua la dung nhu mong doi
	if err == nil {
		t.Error("Error should not be nil")
	}
}

func Test_NoteCreate_TitleHasMaxLength(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1. Chuan bi input dau vao cho ham CreateNote
	title := fakeString(300)
	data := `{"title": "` + title + `","completed": false}`
	ctx := buildMockContext("POST", "/note", data)
	noteRepo := new(mock.NoteRepoImpl)
	// 2. Goi function can test
	_, err := NoteCreate(ctx, noteRepo)
	// 3. Kiem tra ket qua la dung nhu mong doi
	if err == nil {
		t.Error("Error should not be nil")
	}
}

func Test_NoteCreate_TitleIsValid(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1. Chuan bi input dau vao cho ham CreateNote
	data := `{"title": "Should do homework","completed": false}`
	note := model.Note{}
	json.Unmarshal([]byte(data), &note)
	ctx := buildMockContext("POST", "/note", data)
	noteRepo := new(mock.NoteRepoImpl)
	// 2. Goi function can test
	// 2.1 Design cai expectation
	expected := model.Note{
		Title:     "Should do homework",
		Completed: false,
		Model: gorm.Model{
			ID:        28,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	// 2.2 Cach khai bao so 2, neu cach 1 qua phuc tap
	// expected.ID = 28
	// expected.Title = "Should do homework"
	// expected.DeletedAt = nil
	// expected.Completed = false
	// expected.CreatedAt = time.Now()
	// expected.UpdatedAt = time.Now()

	// 2.3 Voi mock test minh phat bieu rang, voi ham Create truyen vao
	// cai note thi tra ve cai expected
	noteRepo.On("Create", note).Return(&expected, nil)
	// 2.4 Phat bieu quan trong
	// Giai su cai ham trong DB tra ve ket qua dung
	// Thi ham NoteCreate minh can test con tra ve ket qua dung khong?
	actual, err := NoteCreate(ctx, noteRepo)
	// 3. Kiem tra ket qua la dung nhu mong doi
	if err != nil {
		t.Error("Error should not be nil")
	}
	if actual.ID != expected.ID {
		t.Error("Actual note should be same expected note")
	}
}

func Test_NoteCreate_TitleMaxLengthWithCorrectError(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1. String len phai lon hon >255 thi kiem tra error message phai dung mong doi.
	title := fakeString(300)
	data := `{"title": "` + title + `","completed": false}`
	note := model.Note{}
	json.Unmarshal([]byte(data), &note)
	ctx := buildMockContext("POST", "/note", data)
	noteRepo := new(mock.NoteRepoImpl)
	// 2. Goi function can test
	// 2.1 Design cai expectation = nil
	// vi cai length > 100 quy dinh trong db
	var expected *model.Note
	// 2.2  Mock function
	expectedErr := errors.New(`Key: 'Note.Title' Error:Field validation for 'Title' failed on the 'max' tag`)
	noteRepo.On("Create", note).Return(expected, expectedErr)
	// 2.4 Phat bieu quan trong
	// Giai su cai ham trong DB tra ve ket qua dung
	// Thi ham NoteCreate minh can test con tra ve ket qua dung khong?
	actual, err := NoteCreate(ctx, noteRepo)
	// 3. Kiem tra ket qua la dung nhu mong doi
	if err.Error() != expectedErr.Error() {
		t.Error("Expected error should be max tag")
	}
	if actual != nil {
		t.Error("Actual should be nil")
	}
}

func Test_NoteUpdate_TitleMaxLengthIsHitLimit(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1.1 "[Editted] " len=10 + 245 = 255 still ok
	title := fakeString(245)
	id := 1
	data := `{"title": "` + title + `","completed": false}`
	ctx := buildMockContext("PUT", "/note/"+strconv.Itoa(id), data)
	ctx.Params = append(ctx.Params, gin.Param{"id", strconv.Itoa(id)})
	// 1.2 Mock db
	noteRepo := new(mock.NoteRepoImpl)
	note := model.Note{}
	json.Unmarshal([]byte(data), &note)
	note.Title = "[Editted] " + note.Title
	noteRepo.On("Update", id, note).Return(nil)
	// 2  Mock function
	err := NoteUpdate(ctx, noteRepo)
	// 3. Kiem tra ket qua la dung nhu mong doi
	if err != nil {
		t.Error("This should not be error")
	}
}

// 1. Doi voi test case nay khong nen test voi case NoteCreate
// 2. Su dung tinh huong cua NoteUpdate de test giup cac ban de hieu hon
func Test_NoteUpdate_TitleMaxLengthCorrectDBError(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1. Tinh huong test pass validation (<255)
	// 2. Nhung sau do cai ham xu ly lam cho cai Title > 255
	title := fakeString(255)
	id := 1
	data := `{"title": "` + title + `","completed": false}`
	ctx := buildMockContext("PUT", "/note/"+strconv.Itoa(id), data)
	ctx.Params = append(ctx.Params, gin.Param{"id", strconv.Itoa(id)})
	noteRepo := new(mock.NoteRepoImpl)
	// 2  Mock function
	err := NoteUpdate(ctx, noteRepo)
	// 3. Kiem tra ket qua la dung nhu mong doi
	expectedErr := errors.New(`Error 1406: Data too long for column 'title' at row 1`)
	if err == nil || err.Error() != expectedErr.Error() {
		t.Error("Expected error should be DB error")
	}
}
