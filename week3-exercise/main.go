package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Note struct {
	gorm.Model
	Title  string
	Status bool
}

type Pagination struct {
	Page  uint `form:"p"`
	Limit uint `form:"l"`
}

func (self *Pagination) GetPage() uint {
	if self.Page == 0 {
		return 1
	}
	if self.Page > 100 {
		return 100
	}
	return self.Page
}

func (self *Pagination) GetOffset() uint {
	page := self.GetPage()
	limit := self.GetLimit()
	offset := (page - 1) * limit
	return offset
}

func (self *Pagination) GetLimit() uint {
	if self.Limit == 0 || self.Limit > 5 {
		return 5
	}
	return self.Limit
}

func main() {
	db, _ := gorm.Open("mysql", "default:secret@/default?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()
	db.AutoMigrate(&Note{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(201, "Pong\n")
	})

	r.POST("/note", func(c *gin.Context) {
		var note Note
		c.ShouldBind(&note)
		db.Create(&note)
		c.JSON(200, note)
	})
	r.GET("/note/:id", func(c *gin.Context) {
		id := c.Param("id")
		var note Note
		err := db.Where("id = ?", id).First(&note).Error
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		c.JSON(200, note)
	})
	r.GET("/note", func(c *gin.Context) {
		var pagination Pagination
		c.ShouldBindQuery(&pagination)
		offset := pagination.GetOffset()
		limit := pagination.GetLimit()

		var notes []Note
		db.Offset(offset).
			Limit(limit).
			Find(&notes)
		c.JSON(http.StatusOK, notes)
	})
	r.PUT("/note/:id", func(c *gin.Context) {
		// Xu ly du lieu tu client
		id, _ := strconv.Atoi(c.Param("id"))
		var note Note
		db.Where("id = ?", id).First(&note)
		c.ShouldBind(&note)
		note.ID = uint(id)
		// Update vo Db
		err := db.Model(&note).Updates(&note).Error
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		// Tra ket qua ve client
		c.JSON(http.StatusOK, note)
	})
	r.DELETE("/note/:id", func(c *gin.Context) {
		// Xu ly du lieu tu client
		id, _ := strconv.Atoi(c.Param("id"))

		err := db.Where("id = ?", id).Delete(&Note{}).Error
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		// Tra ket qua ve client
		c.JSON(http.StatusOK, true)
	})

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	go func() {
		srv.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)

}
