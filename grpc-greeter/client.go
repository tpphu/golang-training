package main

import (
	"flag"
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
	// 2. Create client
	greeter := proto.NewGreeterService("greeter", service.Client())
	// 2.1 Get arguments
	var name string
	flag.StringVar(&name, "name", "Phu", "Input name")
	flag.Parse()
	// 2.2 Call the greeter
	// Note: TODO returns a non-nil, empty Context.
	// Code should use context.TODO
	// when it's unclear which Context to use
	// or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: name})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print response
	fmt.Println(rsp.Greeting)
}
