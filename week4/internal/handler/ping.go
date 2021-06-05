package handler

import (
	"phudt/week4/internal/provider"

	"github.com/gin-gonic/gin"
)

type DefaultHandler struct {
	rp *provider.Provider
}

func NewDefaultHandler(rp *provider.Provider) DefaultHandler {
	return DefaultHandler{
		rp: rp,
	}
}

func (h DefaultHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
