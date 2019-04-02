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
	req := pb.NoteReq{
		Title: "Todo 123",
	}
	res, _ := client.Create(context.TODO(), &req)
	// 4. In ket qua
	fmt.Println("Response:", res)
}
