package main

import (
	"./model"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "default:secret@/week3note?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&model.Note{})
	r := gin.Default()
	group := r.Group("/note")
	group.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		note := model.Note{}
		db.Where("id = ?", id).First(&note)
		c.JSON(200, note)
	})
	group.POST("", func(c *gin.Context) {
		note := model.Note{}
		if err := c.ShouldBind(&note); err != nil {
			c.JSON(400, gin.H{
				"message": "error",
			})
		}
		db.Create(&note)
		c.JSON(200, note)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8083") // listen and serve on 0.0.0.0:8080
}
