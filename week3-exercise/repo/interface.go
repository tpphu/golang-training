package repo

import "../model"

type NoteRepo interface {
	Find(int) (*model.Note, error)
	List() ([]model.Note, error)
	Update(*model.Note) error
	Delete(int) error
	Create(model.Note) error
}
