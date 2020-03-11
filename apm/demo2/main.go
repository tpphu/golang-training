package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.GET("/ping2", func(c *gin.Context) {
		tx := apm.TransactionFromContext(c)
		defer tx.End()
		fmt.Println("tx.EnsureParent()=", tx.EnsureParent())
		// transaction := apm.DefaultTracer.StartTransaction("GET /ping2.1", "request")
		// defer transaction.End()
		time.Sleep(time.Second * 1)
		c.JSON(200, gin.H{
			"message": "pong2",
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
