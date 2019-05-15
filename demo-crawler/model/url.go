package model

import (
	"github.com/jinzhu/gorm"
)

type UrlState int
type UrlStatus int

const (
	UrlStateIdle    UrlState = iota + 1 //1
	UrlStateRunning                     //2
)

const (
	UrlStatusReady          UrlStatus = iota + 1 //1
	UrlStatusSuccess                             //2
	UrlStatusStopped                             //3
	UrlStatusError                               //4
	UrlStatusNotFoundParser                      //5
)

type Url struct {
	gorm.Model
	Url              string
	WebsiteID        int
	Status           UrlStatus
	State            UrlState
	HttpDownloadCode int
}
