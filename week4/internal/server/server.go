package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type server struct {
	port   string
	Engine *gin.Engine
}

func NewServer(port string) server {
	return server{
		port:   port,
		Engine: gin.Default(),
	}
}

func (s server) Start() error {
	return s.Engine.Run(fmt.Sprintf("0.0.0.0:%s", s.port))
}
