package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AppErrorCode int

const (
	// Used when have db error
	ErrDatabase AppErrorCode = 1
)

type Blog struct {
	ID      int
	Title   string `json:"title" binding:"min=10,max=255"`
	Content string `json:"content" binding:"min=1,max=4000"`
}

func (Blog) TableName() string {
	return "blog"
}

func main() {
	/**
	|-------------------------------------------------------------------------
	| Connect to db
	|-----------------------------------------------------------------------*/
	dsn := "default:secret@tcp(127.0.0.1:3306)/dogfood?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	/**
	|-------------------------------------------------------------------------
	| Init GIN
	|-----------------------------------------------------------------------*/
	r := gin.Default()

	r.GET("/health/live", func(c *gin.Context) {
		c.String(200, "ok")
	})

	blogRouter := r.Group("/blog")
	p := bluemonday.UGCPolicy()

	/**
	|-------------------------------------------------------------------------
	|
	|-----------------------------------------------------------------------*/
	blogRouter.Use(func(c *gin.Context) {
		lang, _ := c.Cookie("lang")
		if lang == "" {
			lang = "en"
		}
		c.Set("lang", lang)
	})

	blogRouter.Use(func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	})
	blogRouter.POST("/", func(c *gin.Context) {
		lang, _ := c.Get("lang")
		blog := Blog{}
		// if err := c.BindQuery(&blog); err != nil {
		// 	// todo
		// 	fmt.Println(err)
		// }
		c.Header("X-PHU-DT", lang.(string))
		if err := c.ShouldBindJSON(&blog); err != nil {
			// todo
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    ErrDatabase,
				"message": "Invalid format data",
				"detail":  err.Error(),
			})
			return
		}
		/**
		|-------------------------------------------------------------------------
		| Cho nay lat nua can ban
		|-----------------------------------------------------------------------*/
		blog.Title = p.Sanitize(blog.Title)
		blog.Content = p.Sanitize(blog.Content)
		if err := db.Create(&blog).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    ErrDatabase,
				"message": "Can not create blog",
				"detail":  err.Error(),
			})
			return
		}
		c.JSON(200, blog)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
