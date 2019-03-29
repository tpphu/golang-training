package handler

import (
	"fmt"

	"../repo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRoutes(engine *gin.Engine, db *gorm.DB) {
	engine.GET("/ping", pingHandler)
	note := engine.Group("/note")
	note.Use(simpleMiddleware)
	{
		note.GET("/:id", func(c *gin.Context) {
			repo := &repo.NoteRepoImpl{
				DB: db,
			}
			result, err := NoteGet(c, repo)
			simpleReturnHandler(c, err, result)
		})
		note.PUT("/:id", func(c *gin.Context) {
			repo := &repo.NoteRepoImpl{
				DB: db,
			}
			err := NoteUpdate(c, repo)
			simpleReturnHandler(c, err, nil)
		})
		note.DELETE("/:id", func(c *gin.Context) {
			repo := &repo.NoteRepoImpl{
				DB: db,
			}
			err := NoteDelete(c, repo)
			simpleReturnHandler(c, err, nil)
		})
		note.POST("", func(c *gin.Context) {
			repo := &repo.NoteRepoImpl{
				DB: db,
			}
			result, err := NoteCreate(c, repo)
			simpleReturnHandler(c, err, result)
		})
	}
}

func simpleReturnHandler(c *gin.Context, err error, result interface{}) {
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, result)
}

func simpleMiddleware(c *gin.Context) {
	// if true {
	// 	c.AbortWithStatus(400)
	// 	return
	// }
	fmt.Println("Print here for every request")
	c.Next()
}
