package main

import (
	"context"
	"fmt"

	product "../proto"
	"google.golang.org/grpc"
)

func main() {
	// Syntax
	server := "localhost:50001"
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := product.NewProductServiceClient(conn)
	TestGet(client)
}

func TestGet(client product.ProductServiceClient) {
	getReq := product.GetReq{
		Id: 1,
	}
	getRes, err := client.Get(context.TODO(), &getReq)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response of Get method is:", getRes)
}

func TestAdd(client product.ProductServiceClient) {
	addReq := product.AddReq{
		Name:  "iphone 6",
		Sku:   "IP6",
		Price: 10000000,
	}
	addRes, err := client.Add(context.TODO(), &addReq)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response of Add method is:", addRes)
}
