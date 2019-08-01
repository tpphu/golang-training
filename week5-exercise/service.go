package main

import (
	"fmt"
	"time"

	"context"

	proto "./proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
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
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
