package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handler"
	"./model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// 0. Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// 1. Lien quan toi database
	db, err := gorm.Open("mysql", "default:secret@/notes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&model.Note{}, &model.User{}, &model.Setting{})

	// 2. Write access log ra file & de giu lai cai Println -> Stdout
	fileWriter, err := os.Create("access.log")
	if err != nil {
		panic(err)
	}
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = fileWriter

	// 3. Tao ra router
	r := gin.Default()
	handler.InitRoutes(r, db) // Move cai code minh lam qua cho khac
	// 4. Start chuong trinh
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8081"
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	// 5. Handle stop chuong trinh
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)

}
