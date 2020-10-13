package main

import "github.com/gin-gonic/gin"

func main() {
	/**
	|-------------------------------------------------------------------------
	| Create defaul t/ new default
	|-----------------------------------------------------------------------*/
	r := gin.Default()
	// Add mot route/handler
	// Context lat nua giai thich
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	counter := 0
	r.GET("/incre", func(c *gin.Context) {
		counter = counter + 1
		c.JSON(200, gin.H{
			"counter": counter,
		})
	})
	r.GET("/incree", func(c *gin.Context) {
		counter = counter + 1
		c.JSON(200, gin.H{
			"counter": counter,
			"key":     "counter2",
		})
	})
	// Run
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
