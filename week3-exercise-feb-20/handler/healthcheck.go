package handler

import "github.com/gin-gonic/gin"

type HealthcheckHandler struct {
	Engine *gin.Engine
}

func (h HealthcheckHandler) inject() {
	h.Engine.GET("/ping", h.ping)
}

func (h HealthcheckHandler) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
