package handler

import "github.com/gin-gonic/gin"

func pingHandler(c *gin.Context) {
	c.String(201, "Pong\n")
}
