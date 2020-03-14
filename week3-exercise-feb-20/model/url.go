package model

import "github.com/jinzhu/gorm"

type Url struct {
	gorm.Model
	URL    string
	Status bool
}
