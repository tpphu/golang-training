package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ArticleStatus int

const (
	ArticleStatusSuccess ArticleStatus = iota + 1
	ArticleStatusParseError
)

type Article struct {
	gorm.Model
	UrlID       uint `gorm:"not null;unique"`
	Title       string
	PublishedAt time.Time
	Content     string `gorm:"type:varchar(4000)"`
	Author      string
	Status      ArticleStatus
}
