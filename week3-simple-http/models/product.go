package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	SKU   string
	Price uint
}
