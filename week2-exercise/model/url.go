package model

import "github.com/jinzhu/gorm"

type UrlState int

const UrlStateIdle UrlState = 1
const UrlStateRuning UrlState = 2

type Url struct {
	gorm.Model
	URL              string
	Status           bool
	State            UrlState `gorm:"default:1"`
	DownloadHttpCode int
}
