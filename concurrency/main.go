package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Note struct {
	gorm.Model // 1. Embedded Struct
	Title      string
	CreatorID  int
}

type User struct {
	gorm.Model
	Name string
}

// 2.1 Muc tieu la de sau nay viet test
func getNote(db *gorm.DB, creatorID int) (*Note, error) {
	note := &Note{}
	err := db.Where("creator_id = ?", creatorID).First(&note).Error
	return note, err
}

// 2.2 Muc tieu la de sau nay viet test
func getCreator(db *gorm.DB, id int) (*User, error) {
	creator := &User{}
	err := db.Where("id = ?", id).First(&creator).Error
	return creator, err
}

func main() {

	db, err := gorm.Open("mysql", "default:secret@/notes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Note{}, &User{})

	creatorID := 1
	// 3. Pattern WaitGroup
	// 3.1 +/- di cai state
	wg := new(sync.WaitGroup)
	note := &Note{}
	creator := &User{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		note, _ = getNote(db, creatorID)
		fmt.Println("????: ", note)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		creator, _ = getCreator(db, creatorID)
	}()

	wg.Wait()

	fmt.Println("note:", note)
	fmt.Println("creator:", creator)

	time.Sleep(time.Second * 30)
}
