package repo

import (
	"../model"
	"github.com/jinzhu/gorm"
)

type NoteRepoImpl struct {
	db *gorm.DB
}

func (self *NoteRepoImpl) Find(id int) (*model.Note, error) {
	return nil, nil
}
