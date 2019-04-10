package main

import (
	"fmt"

	"context"

	proto "./proto"
	"github.com/micro/go-micro"
)

func main() {
	// 1. Create service
	service := micro.NewService(
		micro.Name("greeter"),
	)
	service.Init()
	// 2. Create client
	greeter := proto.NewGreeterService("greeter", service.Client())
	// 2.1 Get arguments
	// 2.2 Call the greeter
	// Note: TODO returns a non-nil, empty Context.
	// Code should use context.TODO
	// when it's unclear which Context to use
	// or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Dau"})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print response
	fmt.Println(rsp.Greeting)
}
