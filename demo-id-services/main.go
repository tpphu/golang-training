package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("Hello world!")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	maxValue := 0
	currentValue := 0
	db, _ := gorm.Open("mysql", "root:root@/demo?charset=utf8&parseTime=True&loc=Local")
	r.GET("/get-increment-id", func(c *gin.Context) {
		// 1. Get value hien tai by key name
		// SELECT `value` from settings where `key`='incr' LIMIT 1
		if currentValue >= maxValue {
			tx := db.Begin()
			tx.Raw("SELECT `value` from settings where `key`='incr' LIMIT 1 FOR UPDATE").
				Row().
				Scan(&currentValue)
			// 2. Tang cho no mot don vi
			maxValue = currentValue + 100
			// 3. Cap nhat vao db
			tx.Exec("UPDATE settings SET `value` = ? where `key`='incr' LIMIT 1", maxValue)
			tx.Commit()
		}
		currentValue += 1
		c.JSON(200, gin.H{
			"incre": currentValue,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
