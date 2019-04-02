package main

import (
	"context"
	"fmt"

	pb "../proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// 1. Connect to server at TCP port
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	// 2. New client
	client := pb.NewNoteServiceClient(conn)
	// 3. Call Create
	req := pb.NoteFindReq{
		Id: 13,
	}
	res, _ := client.Find(context.TODO(), &req)
	// 4. In ket qua
	fmt.Println("Response:", res)
}
