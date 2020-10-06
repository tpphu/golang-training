package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("1")
	db, err := sql.Open("mysql", "default:secret@(127.0.0.1:3306)/dogfood")
	fmt.Println("2")
	fmt.Println(db, err)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err=db.Ping()
	fmt.Println("test", err)
}
