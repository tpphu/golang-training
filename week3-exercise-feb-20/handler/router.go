package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRouter(r *gin.Engine, db *gorm.DB) {
	// Register Healthcheck Handler
	healthcheckHandler := &HealthcheckHandler{
		Engine: r,
	}
	healthcheckHandler.inject()
	// Inject Note Handler
	noteHandler := &NoteHandler{
		Engine: r,
		DB:     db,
	}
	noteHandler.inject()
}
