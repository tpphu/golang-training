package main

import (
	"phudt/week4/internal/config"
	"phudt/week4/internal/handler"
	"phudt/week4/internal/model"
	"phudt/week4/internal/provider"
	"phudt/week4/internal/server"

	"github.com/urfave/cli/v2"
)

func ServeAction(ctx *cli.Context) error {
	port := ctx.String("port")
	cfg := config.Config{}
	cfg.MySQL = config.MySQLConfigDefault()
	rp := provider.NewProvider(&cfg)
	rp.DB.AutoMigrate(&model.User{}, &model.Department{})
	s := server.NewServer(port)
	errChan := make(chan error)
	go func() {
		errChan <- s.Start()
	}()
	r := s.Engine
	defaultHandler := handler.NewDefaultHandler(&rp)
	r.GET("/ping", defaultHandler.Ping)
	userHandler := handler.NewUserHandler(&rp)
	r.POST("/v1/user", userHandler.Create)
	r.POST("/v2/user", userHandler.CreateWithDepartment)
	for {
		select {
		case err := <-errChan:
			return err
		}
	}
}
