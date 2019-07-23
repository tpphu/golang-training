package model

import "github.com/jinzhu/gorm"

type UrlState int

const UrlStateIdle UrlState = 1
const UrlStateRuning UrlState = 2

type Url struct {
	gorm.Model
	Url              string
	Status           int
	State            UrlState
	DownloadHttpCode int
}
