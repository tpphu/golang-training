package handler

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http/httptest"
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

func buildMockContext(data string) *gin.Context {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/note", strings.NewReader(data))
	ctx.Request.Header.Set("Content-Type", "application/json")
	return ctx
}
func Test_NoteCreate_TitleIsEmpty(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1. Chuan bi input dau vao cho ham CreateNote
	data := `{"title": "","completed": false}`
	ctx := buildMockContext(data)
	noteRepo := new(mock.NoteRepoImpl)
	// 2. Goi function can test
	_, err := NoteCreate(ctx, noteRepo)
	// 3. Kiem tra ket qua la dung nhu mong doi
	if err == nil {
		t.Error("Error should not be nil")
	}
}

func Test_NoteCreate_TitleHasMinLength(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1. Chuan bi input dau vao cho ham CreateNote
	data := `{"title": "abc","completed": false}`
	ctx := buildMockContext(data)
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
	largeStr := fakeString(2000)
	data := `{"title": "` + largeStr + `","completed": false}`
	ctx := buildMockContext(data)
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
	ctx := buildMockContext(data)
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
	// 1. String len phai lon hon >1000 thi kiem tra error message phai dung mong doi.
	largeStr := fakeString(2000)
	data := `{"title": "` + largeStr + `","completed": false}`
	note := model.Note{}
	json.Unmarshal([]byte(data), &note)
	ctx := buildMockContext(data)
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

func Test_NoteCreate_TitleMaxLengthCorrectDBError(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1. Tinh huong test pass validation (<1000)
	// 2. nhung lon hon gia tri cho phep trong DB la 255
	largeStr := fakeString(500)
	data := `{"title": "` + largeStr + `","completed": false}`
	ctx := buildMockContext(data)
	noteRepo := new(mock.NoteRepoImpl)
	// 2  Mock function
	actual, err := NoteCreate(ctx, noteRepo)
	// 3. Kiem tra ket qua la dung nhu mong doi
	expectedErr := errors.New(`Error 1406: Data too long for column 'title' at row 1`)
	if err.Error() != expectedErr.Error() {
		t.Error("Expected error should be DB error")
	}
	if actual != nil {
		t.Error("Actual should be nil")
	}
}
