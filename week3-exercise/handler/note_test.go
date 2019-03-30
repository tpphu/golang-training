package handler

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	m "../mock"
	"../model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Khong nen bind o day vi no se bi loi EOF
// vi khi vao ben trong goi lai ham bin
// c.BindJSON(&note) // Khong nen goi

func TestNoteCreateWithInValidValidator(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// 1. Gia lap context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	data := `{"title":"todo"}`
	c.Request = httptest.NewRequest("POST", "/note", strings.NewReader(data))
	c.Request.Header.Set("Content-Type", "application/json")
	noteRepo := new(m.NoteRepoImpl)
	_, err := NoteCreate(c, noteRepo)
	if err == nil {
		t.Error("Error should not be nil because this will not valid min rule validator")
	}
}

func TestNoteCreateWithValidValidator(t *testing.T) {
	// 1. Gia lap context
	gin.SetMode(gin.ReleaseMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	data := `{"title":"todo 123"}`
	c.Request = httptest.NewRequest("POST", "/note", strings.NewReader(data))
	c.Request.Header.Set("Content-Type", "application/json")
	// 2. Logic nay boc lo code thiet ke khong dung
	// 2.1 Do phai dung json.Unmarshal de thanh gia tri truyen vao ham On("Create")
	noteRepo := new(m.NoteRepoImpl)
	note := model.Note{}
	json.Unmarshal([]byte(data), &note)
	actual := model.Note{
		Title: "todo 123",
		Model: gorm.Model{
			ID:        123,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	noteRepo.On("Create", note).Return(actual, nil)
	expected, err := NoteCreate(c, noteRepo)
	if err != nil {
		t.Error("Error should be nil")
	}
	if expected.ID != 123 {
		t.Error("note.ID should be 123")
	}
}
