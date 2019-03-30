package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

var counter int = 0

func pingHandler(c *gin.Context) {
	counter += 1
	// Race condition => Hai CPU cung access vo 1 cai bien
	// Atomic => co 100 request vao nhung expected: counter = 100, actual: 80
	c.Writer.Header().Set("X-Counter", strconv.Itoa(counter))
	c.String(201, "Pong\n")
}
