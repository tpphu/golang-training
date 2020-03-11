package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.GET("/ping3", func(c *gin.Context) {
		time.Sleep(time.Second * 1)
		c.JSON(200, gin.H{
			"message3": "pong3",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
