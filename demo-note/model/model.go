package model

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Title     string `gorm:"type:varchar(500)"`
	Completed bool
	OwnerId   int
}

type User struct {
	gorm.Model
	Fullname string
}

func GetNoteById(db *gorm.DB, id int) (Note, error) {
	note := Note{}
	err := db.Find(&note, id).Error
	return note, err
}

func GetUserById(db *gorm.DB, id int) (User, error) {
	user := User{}
	err := db.Find(&user, id).Error
	return user, err
}
