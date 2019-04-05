package model

import "github.com/jinzhu/gorm"

type Setting struct {
	gorm.Model
	Key         string `gorm:"unique;not null"`
	ValueInt    uint32
	ValueString string
}
