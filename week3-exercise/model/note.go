package model

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Title     string
	Completed bool
}
