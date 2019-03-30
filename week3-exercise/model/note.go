package model

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Title     string `binding:"required,min=10"`
	Completed bool
}
