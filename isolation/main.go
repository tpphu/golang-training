package main

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Counters struct {
	ID      uint `gorm:"primaryKey"`
	Counter int
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "default:secret@tcp(127.0.0.1:3306)/dogfood?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB, _ := db.DB()
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Counters{})
	counter := Counters{ID: 1}
	wg := new(sync.WaitGroup)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			// The update command for each statement is atomic
			// So no need to start a transaction
			db.Model(&counter).Where("counter > ?", 0).Update("counter", gorm.Expr("counter-1"))
		}()
	}
	wg.Wait()
}
