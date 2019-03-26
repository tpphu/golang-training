package model

import (
	"github.com/jinzhu/gorm"
)

type UrlState int

const (
	UrlStateIdle UrlState = iota + 1
	UrlStateRunning
)

type UrlStatus int

const (
	UrlStatusReady UrlStatus = iota + 1
	UrlStatusSuccess
	UrlStatusStopped
	UrlStatusError
	UrlStatusNotFoundParser
)

type Url struct {
	gorm.Model
	Url              string
	State            UrlState
	Status           UrlStatus
	DownloadHttpCode int
}
