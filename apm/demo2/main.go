package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

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
	r.GET("/ping2", func(c *gin.Context) {
		tx := apm.TransactionFromContext(c.Request.Context())
		// defer tx.End()
		// fmt.Println("tx.EnsureParent()=", tx.EnsureParent())
		client := apmhttp.WrapClient(http.DefaultClient)
		req, _ := http.NewRequest("GET", "http://localhost:8083/ping3", nil)
		ctx := apm.ContextWithTransaction(c, tx)
		resp, _ := client.Do(req.WithContext(ctx))
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		time.Sleep(time.Second * 1)
		c.JSON(200, gin.H{
			"message2": "pong2",
			"message3": string(body),
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
