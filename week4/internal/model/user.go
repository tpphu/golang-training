package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string
	Fullname     string
	Password     string
	Coin         float64
	DepartmentId uint
	Department   Department `gorm:"foreignKey:DepartmentId"`
}

type Department struct {
	gorm.Model
	Name string
}
