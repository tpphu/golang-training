package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null",binding:"required"`
	Email    string `gorm:"unique;not null",binding:"required"`
	Password string `binding:"required"`
	Fullname string
	Bod      *time.Time
}

type UserLoginForm struct {
	Login    string `binding:"required"`
	Password string `binding:"required"`
}

type UserLoginReponse struct {
	ID       uint
	Fullname string
	Token    string
}
