package main

import (
	"net"

	context "golang.org/x/net/context"

	pb "../proto"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

// Viet cai note service de implement cai service da define
type noteService struct{}

func (self *noteService) Create(ctx context.Context, req *pb.NoteReq) (*pb.Note, error) {
	return &pb.Note{
		Id:          123,
		Title:       req.Title,
		Completed:   req.Completed,
		Description: "A description",
	}, nil
}

func (self *noteService) Find(ctx context.Context, req *pb.NoteFindReq) (*pb.Note, error) {
	return &pb.Note{
		Id:    123,
		Title: "Todo 123",
	}, nil
}

func (self *noteService) Delete(ctx context.Context, req *pb.NoteDelReq) (*pb.NoteDelRes, error) {
	if req.Id == 123 {
		return &pb.NoteDelRes{
			Success: true,
		}, nil
	}
	return &pb.NoteDelRes{
		Success:      false,
		ErrorMessage: "Not Found",
	}, nil
}

func main() {
	// 1. Listen/Open a TPC connect at port
	lis, _ := net.Listen("tcp", port)
	// 2. Tao server tu GRP
	grpcServer := grpc.NewServer()
	// 3. Map service to server
	pb.RegisterNoteServiceServer(grpcServer, &noteService{})
	// 4. Binding port
	grpcServer.Serve(lis)
}
