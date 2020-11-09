package main

import (
	"context"
	"fmt"

	"github.com/tpphu/golang-training/week4/proto"
	"google.golang.org/grpc"
)

func main() {
	server := "localhost:50001"
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	client := proto.NewPersonServiceClient(conn)
	req := &proto.CreatePersonRequest{
		Name:     "tpphu1",
		Fullname: "Tran Phong Phu",
	}
	res, err := client.CretePerson(context.Background(), req)
	if err != nil {
		fmt.Println("Error is:", err.Error())
	}
	fmt.Println("Response:", res)

	req2 := &proto.GetPersonRequest{Id: 1}
	res2, err2 := client.GetPerson(context.Background(), req2)
	if err2 != nil {
		fmt.Println("Error is 2:", err2.Error())
	}
	fmt.Println("Response 2:", res2)
}
