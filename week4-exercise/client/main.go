package main

import (
	"context"
	"fmt"

	pb "github.com/tpphu/week4-exercise/proto"
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
	// 3. Todo
	res, err := updateANote(client)
	// 4
	if err != nil {
		fmt.Println("err:", err)
	}
	// 4. In ket qua
	fmt.Println("Response:", res)
}

func updateANote(client pb.NoteServiceClient) (*pb.Note, error) {
	req := pb.NoteUpdateReq{
		Id:        8,
		Title:     "Todo 8.3", // Chi cap nhat dc field nay.
		Completed: "0",
	}
	return client.Update(context.TODO(), &req)
}

func findANote(client pb.NoteServiceClient) (*pb.Note, error) {
	req := pb.NoteFindReq{
		Id: 8,
	}
	return client.Find(context.TODO(), &req)
}

func createANote(client pb.NoteServiceClient) (*pb.Note, error) {
	req := pb.NoteReq{
		Title:     "Todo 8",
		Completed: true,
	}
	return client.Create(context.TODO(), &req)
}
