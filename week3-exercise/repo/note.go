package repo

import (
	"../helper"
	"../model"
	"github.com/jinzhu/gorm"
)

type NoteRepo interface {
	Find(int) (*model.Note, error)
	List(helper.Pagination) ([]model.Note, error)
	Update(int, model.Note) error
	Delete(int) error
	Create(model.Note) (*model.Note, error)
}

type NoteRepoImpl struct {
	DB *gorm.DB
}

// 1. That su la co mot func phu thuoc vao db
func (self *NoteRepoImpl) Create(note model.Note) (*model.Note, error) {
	err := self.DB.Create(&note).Error
	return &note, err
}

func (self *NoteRepoImpl) Find(id int) (*model.Note, error) {
	note := &model.Note{}
	err := self.DB.Where("id = ?", id).First(note).Error
	return note, err
}

func (self *NoteRepoImpl) List(pagination helper.Pagination) ([]model.Note, error) {
	notes := []model.Note{}
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	err := self.DB.Offset(offset).
		Limit(limit).
		Find(&notes).
		Error
	return notes, err
}

func (self *NoteRepoImpl) Update(id int, note model.Note) error {
	err := self.DB.Where("id = ?", id).Update(&note).Error
	return err
}

func (self *NoteRepoImpl) Delete(id int) error {
	err := self.DB.Where("id = ?", id).Delete(&model.Note{}).Error
	return err
}
