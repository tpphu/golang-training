package main

import (
	"flag"
	"fmt"

	"context"

	proto "./proto"
	"github.com/micro/go-micro"
)

func main() {
	// New service
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)
	// Create client
	greeter := proto.NewGreeterService("greeter", service.Client())
	// Get arguments
	var name string
	flag.StringVar(&name, "name", "Phu", "Input name")
	flag.Parse()
	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: name})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print response
	fmt.Println(rsp.Greeting)
}
