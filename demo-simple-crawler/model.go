package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UrlState int

const (
	UrlStateIdle    UrlState = iota + 1 //1
	UrlStateRunning                     //2
)

type UrlStatus int

const (
	UrlStatusReady          UrlStatus = iota + 1 //1
	UrlStatusSuccess                             //2
	UrlStatusStopped                             //3
	UrlStatusError                               //4
	UrlStatusNotFoundParser                      //5
)

type Article struct {
	gorm.Model
	UrlID       uint `gorm:"not null;unique"`
	Title       string
	PublishedAt time.Time
	Content     string `gorm:"type:varchar(4000)"`
	Author      string
}

type Url struct {
	gorm.Model
	Url              string
	State            UrlState
	Status           UrlStatus
	DownloadHttpCode int
}

type SimpleData struct {
	gorm.Model
	Title       string
	Author      string
	PublishDate string
}
