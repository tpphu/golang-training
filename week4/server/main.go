package main

import (
	"context"
	"errors"
	"net"

	"github.com/tpphu/golang-training/week4/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type PersonService struct {
	DB *gorm.DB
}

func (s *PersonService) CretePerson(ctx context.Context, req *proto.CreatePersonRequest) (*proto.Person, error) {
	if req.Name != "tpphu" {
		return nil, errors.New("Invalid Name")
	}
	return &proto.Person{
		Id:       1,
		Name:     req.Name,
		Fullname: req.Fullname,
	}, nil
}

func (s *PersonService) GetPerson(ctx context.Context, req *proto.GetPersonRequest) (*proto.Person, error) {
	if req.Id != 1 {
		return nil, status.Error(codes.NotFound, "Not found user id 1")
	}
	return &proto.Person{
		Id:       1,
		Name:     "tpphu",
		Fullname: "Tran Phong Phu",
	}, nil
}
func main() {
	// TPC listener
	port := "0.0.0.0:50001"
	lis, _ := net.Listen("tcp", port)
	// New server
	grpcServer := grpc.NewServer()
	// service
	service := PersonService{}
	proto.RegisterPersonServiceServer(grpcServer, &service)
	// Start and attach port
	grpcServer.Serve(lis)
}
