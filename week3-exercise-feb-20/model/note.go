package model

import (
	"github.com/jinzhu/gorm"
	"github.com/tpphu/week2-exercise/form"
)

type Note struct {
	gorm.Model
	Title     string
	Completed *bool
}

func (n *Note) Fill(input form.Note) {
	n.Title = input.Title
	n.Completed = &input.Completed
}
