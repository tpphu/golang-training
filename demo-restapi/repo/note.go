package repo

import (
	"../model"
	"github.com/jinzhu/gorm"
)

// 1. Dinh nghia interface
type NoteRepo interface {
	Create(note model.Note) (model.Note, error)
}

// 2. Define cai struct de ma ready cho viec implement interface
type NoteRepoImpl struct {
	DB *gorm.DB
}

// 3. Phuong thuc create
func (self NoteRepoImpl) Create(input model.Note) (model.Note, error) {
	err := self.DB.Create(&input).Error
	return input, err
}

// 4. NewNoteRepo
