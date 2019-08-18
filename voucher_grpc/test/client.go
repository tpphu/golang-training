package main

import (
	"context"
	"fmt"
	"time"

	"../proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

const (
	address = "localhost:10000"
)

func main() {
	// 1. Connect to server at TCP port
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	// 2. New client
	client := proto.NewVoucherServiceClient(conn)
	// 3. Call Create
	now := time.Now()
	start := &timestamp.Timestamp{
		Seconds: now.Unix(),
	}
	end := &timestamp.Timestamp{
		Seconds: now.Add(time.Hour * 48).Unix(),
	}
	req := proto.VoucherReq{
		Code:     "ABC",
		Discount: 0.05,
		Start:    start,
		End:      end,
	}
	res, err := client.Register(context.TODO(), &req)
	if err != nil {
		panic(err)
	}
	// 4. In ket qua
	fmt.Println("Response:", res)
}
