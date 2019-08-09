package main

import (
	"./model"
	"./storage"
	"github.com/gin-gonic/gin"
)

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "default:secret@/voucher")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		voucher := model.Voucher{}
		voucherStorage := storage.Voucher{
			DB: db,
		}
		if err := c.ShouldBindJSON(&voucher); err != nil {
			c.JSON(400, err)
			return
		}
		isExist, err := voucherStorage.IsExit(voucher)
		if err != nil {
			c.JSON(400, err)
			return
		}
		if isExist {
			c.JSON(400, gin.H{
				"error": "Exist",
			})
			return
		}
		voucherStorage.Register(&voucher)
		c.JSON(200, voucher)
	})
	r.GET("/verify", func(c *gin.Context) {
		voucher := model.Voucher{}
		c.JSON(200, voucher)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
