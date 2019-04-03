package mock

import (
	"../helper"
	"../model"
	"github.com/stretchr/testify/mock"
)

type NoteRepoImpl struct {
	mock.Mock
}

func (self *NoteRepoImpl) Create(note model.Note) (*model.Note, error) {
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
	args := self.Called(id, note)
	return args.Error(0)
}

func (self *NoteRepoImpl) Delete(id int) error {
	args := self.Called(id)
	return args.Error(0)
}
