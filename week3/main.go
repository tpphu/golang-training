package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"phudt/week3/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/covid?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// New mot cai Gin Engine Default
	r := gin.Default()
	// Khong code Go nhu the nay.
	counter := 0
	// Add mot cai route "/ping", method: "GET"
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	/**
	|-------------------------------------------------------------------------
	| concurrency: 100 (k6: -u --vus = 100)
	| total: 500k request (k6: -i 500000)
	| Khong thee tra ve con so 500k dc: 49xk?
	| race condition
	|-----------------------------------------------------------------------*/
	var mutex = &sync.Mutex{}
	r.GET("/counter", func(c *gin.Context) {
		// lock => performance kem lai
		// thu 5 tuan toi.
		// sync Go - vi la goroutines (vs. Nodejs se can co async)
		mutex.Lock()
		counter++
		mutex.Unlock()
		//unlock
		c.JSON(200, gin.H{
			"counter": counter,
		})
	})
	/**
	|-------------------------------------------------------------------------
	| Create:
	| Get:
	| Update
	| List
	| Delete: Không có delete dữ liệu. (soft delete) -> hide, active/delete
	|-----------------------------------------------------------------------*/
	r.POST("/patient", func(c *gin.Context) {
		patient := model.Patient{}
		if err := c.ShouldBindJSON(&patient); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		err = db.Create(&patient).Error
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, patient)
	})
	r.GET("/patient/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.String(http.StatusBadRequest, "Invalid id")
			return
		}
		patient := model.Patient{}
		err := db.First(&patient, id).Error
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, patient)
	})
	r.PUT("/patient/:id", func(c *gin.Context) {
		// id := c.Param("id")
		// if id == "" {
		// 	c.String(http.StatusBadRequest, "Invalid id")
		// 	return
		// }
		// patient := model.Patient{}
		// err := db.First(&patient, id).Error
		// if err != nil {
		// 	c.String(http.StatusInternalServerError, err.Error())
		// 	return
		// }
		// c.JSON(http.StatusOK, patient)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(addr)
}
