package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	UrlID       uint
	Title       string
	Content     string
	Author      string
	PublishedAt time.Time
}
