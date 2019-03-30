package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

var i int = 0

func pingHandler(c *gin.Context) {
	i += 1
	// Race condition => Hai CPU cung access vo 1 cai bien
	// Atomic => co 100 request vao nhung expected: i = 100, actual: 80
	c.Writer.Header().Set("Gia-Tri-X", strconv.Itoa(i))
	c.String(201, "Pong\n")
}
