package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmhttp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.GET("/ping1", func(c *gin.Context) {
		transaction := apm.DefaultTracer.StartTransaction("GET /ping1.1", "request")
		defer transaction.End()
		span := transaction.StartSpan("call api pong2", "http", nil)
		defer span.End()
		// Call ping 2
		client := apmhttp.WrapClient(http.DefaultClient)
		req, _ := http.NewRequest("GET", "http://localhost:8082/ping2", nil)
		ctx := apm.ContextWithTransaction(c, transaction)
		resp, _ := client.Do(req.WithContext(ctx))
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		c.JSON(200, gin.H{
			"message1": "pong1",
			"message2": string(body),
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
