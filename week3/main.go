package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)
	dsn := "default:secret@tcp(127.0.0.1:3306)/dogfood?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
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

	blogRouter.GET("/", func(c *gin.Context) {
		blogs := []Blog{}
		strSize := c.DefaultQuery("size", "2")
		strPage := c.DefaultQuery("page", "1")
		q := c.DefaultQuery("q", "")
		user_ids_str := c.DefaultQuery("user_ids", "")
		user_ids := []string{}
		if user_ids_str != "" {
			user_ids = strings.Split(user_ids_str, ",")
		}
		size, _ := strconv.ParseInt(strSize, 10, 32)
		page, _ := strconv.ParseInt(strPage, 10, 32)
		builder := db.Limit(int(size)).Offset(int(page)).Order("id desc")
		if q != "" {
			/**
			|-------------------------------------------------------------------------
			| Khong tot, minh nen tao _search (index field), fulltext search
			| Analyzer => Filter nhung ky tu khong can thiet html, loai dup tu moi insert
			|-----------------------------------------------------------------------*/
			builder.Where(`title LIKE ? OR content LIKE ?`, "%"+q+"%", "%"+q+"%")
		}
		if len(user_ids) > 0 {
			builder.Where(`user_id IN (?)`, user_ids)
		}
		if err := builder.Find(&blogs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    ErrDatabase,
				"message": "Can not list blogs",
				"detail":  err.Error(),
			})
			return
		}
		c.JSON(200, blogs)
		return
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
