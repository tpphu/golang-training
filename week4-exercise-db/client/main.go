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
	// req := pb.NoteFindReq{
	// 	Id: 8,
	// }
	// res, _ := client.Find(context.TODO(), &req)
	// // 4. In ket qua
	// fmt.Println("Response:", res)

	NoteUpdate(client)
}

func NoteUpdate(client pb.NoteServiceClient) {
	req := pb.NoteUpdateReq{
		Id:        8,
		Title:     "[Updated] Todo 8",
		Completed: true,
	}
	note, _ := client.Update(context.TODO(), &req)
	fmt.Println("Response:", note)
}
