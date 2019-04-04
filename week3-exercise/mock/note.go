package mock

import (
	"errors"

	"../helper"
	"../model"
	"github.com/stretchr/testify/mock"
)

type NoteRepoImpl struct {
	mock.Mock
}

func (self *NoteRepoImpl) Create(note model.Note) (*model.Note, error) {
	if len(note.Title) > 255 {
		return nil, errors.New(`Error 1406: Data too long for column 'title' at row 1`)
	}
	args := self.Called(note)
	return args.Get(0).(*model.Note), args.Error(1)
}

func (self *NoteRepoImpl) Find(id int) (*model.Note, error) {
	args := self.Called(id)
	return args.Get(0).(*model.Note), args.Error(1)
}

func (self *NoteRepoImpl) List(pagination helper.Pagination) ([]model.Note, error) {
	args := self.Called(pagination)
	return args.Get(0).([]model.Note), args.Error(1)
}

func (self *NoteRepoImpl) Update(id int, note model.Note) error {
	if len(note.Title) > 255 {
		return errors.New(`Error 1406: Data too long for column 'title' at row 1`)
	}
	args := self.Called(id, note)
	return args.Error(0)
}

func (self *NoteRepoImpl) Delete(id int) error {
	args := self.Called(id)
	return args.Error(0)
}
