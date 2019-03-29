package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine) {
	engine.GET("/ping", pingHandler)
	note := engine.Group("/note")
	{
		note.GET("/:id", NoteGet)
		note.PUT("/:id", NoteUpdate)
		note.DELETE("/:id", NoteDelete)
		note.POST("", NoteCreate)
	}
}
