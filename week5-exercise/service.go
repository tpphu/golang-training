package main

import (
	"fmt"
	"time"

	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	proto "github.com/tpphu/week5-exercise/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Xin chao " + req.Name
	fmt.Println(time.Now(), "Hello from server")
	return nil
}

func main() {
	// 1. Create service
	service := grpc.NewService(
		micro.Name("greeter"), // Quan trong la co mot cai ten
	)
	// 1.1 Registry to Consul
	service.Init()

	// 2. Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 3.  Run the server
	fmt.Println("Start server")
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
