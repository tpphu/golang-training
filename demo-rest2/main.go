package main

import (
	"fmt"

	"./model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	note, user := getNoteAndUser()
	fmt.Println("note:", note)
	fmt.Println("user:", user)
}

func getNoteAndUser() (model.Note, model.User) {
	db, err := gorm.Open("mysql", "root:root@/demo_note?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&model.Note{}, &model.User{})
	ch := make(chan bool, 2)

	note := model.Note{}
	user := model.User{}
	go func() {
		note, _ = model.GetNoteById(db, 1)
		ch <- true
	}()

	go func() {
		defer func() {
			ch <- true
		}()
		user, err = model.GetUserById(db, 2)
		if err != nil {
			return
		}

	}()
	<-ch
	<-ch

}
