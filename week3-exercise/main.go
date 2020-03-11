package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		delay, _ := strconv.Atoi(c.Query("delay"))
		time.Sleep(time.Second * time.Duration(delay))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	counter := 0
	r.GET("/counter", func(c *gin.Context) {
		counter = counter + 1
		c.JSON(200, gin.H{
			"counter": counter,
		})
	})
	r.POST("/note", func(c *gin.Context) {
		type Note struct {
			gorm.Model
			Title      string
			Completed  bool
			CategoryId int
		}
		form := Note{}
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		catId, _ := strconv.Atoi(c.Query("category_id"))
		fmt.Println("db_uri:", os.Getenv("DB_URI"))
		db, err := gorm.Open("mysql", os.Getenv("DB_URI"))
		db.DropTableIfExists(&form)
		db.AutoMigrate(&form)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		form.CreatedAt = time.Now()
		form.CategoryId = catId

		db.Create(&form)
		c.JSON(200, form)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
