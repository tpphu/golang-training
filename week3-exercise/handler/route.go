package handler

import (
	"fmt"

	"../repo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRoutes(engine *gin.Engine, db *gorm.DB) {
	engine.GET("/ping", pingHandler)
	groupRouter := engine.Group("/note")

	// 1. Authentication // Identity
	// 2. Lam logger/tracking
	// 3. Recovery
	// 4. Add nhieu cai middleware va no chay tuan tu
	groupRouter.Use(simpleMiddleware)
	{
		groupRouter.GET("/:id", func(c *gin.Context) {
			noteRepository := &repo.NoteRepoImpl{
				DB: db,
			}
			result, err := NoteGet(c, noteRepository)
			simpleReturnHandler(c, err, result)
		})
		groupRouter.POST("", func(c *gin.Context) {
			// 1. Repo
			repo := &repo.NoteRepoImpl{
				DB: db,
			}
			// 2. Create note
			result, err := NoteCreate(c, repo)
			// 3. Handle result & err
			simpleReturnHandler(c, err, result)
		})
		groupRouter.PUT("/:id", func(c *gin.Context) {
			repo := &repo.NoteRepoImpl{
				DB: db,
			}
			err := NoteUpdate(c, repo)
			simpleReturnHandler(c, err, nil)
		})
		groupRouter.DELETE("/:id", func(c *gin.Context) {
			repo := &repo.NoteRepoImpl{
				DB: db,
			}
			err := NoteDelete(c, repo)
			simpleReturnHandler(c, err, nil)
		})
	}
}

func simpleReturnHandler(c *gin.Context, err error, result interface{}) {
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, result)
}

func simpleMiddleware(c *gin.Context) {
	// 1. Logic quan trong nhat la, xu ly ngung cai cai request
	// 2. Tao ra cac du lieu de set vao context cho cai handler dung
	// 3. Quyet dinh cho phep di tiep den cai middleware tiep theo hoac handler
	// if c.GetHeader("token") != "202cb962ac59075b964b07152d234b70" {
	// 	c.AbortWithStatus(400)
	// 	return
	// }
	fmt.Println("Print here for every request")
	c.Next()
}
