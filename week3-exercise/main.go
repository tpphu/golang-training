package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handler"
	"./model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "default:secret@/notes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&model.Note{})

	r := gin.Default()
	handler.InitRoutes(r, db)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)

}
