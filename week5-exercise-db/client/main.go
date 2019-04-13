package main

import (
	"fmt"

	"context"

	proto "../proto"
	"github.com/micro/go-micro"
)

func main() {
	// 1. Create service
	service := micro.NewService(
		micro.Name("note-service"),
	)
	service.Init()
	// 2. Create client
	client := proto.NewNoteService("note-service", service.Client())
	res, err := client.Create(context.TODO(), &proto.NoteCreateReq{Title: "Week 5"})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print response
	fmt.Println(res)
}
