package main

import (
	"net"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	context "golang.org/x/net/context"

	model "../model"
	pb "../proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type noteService struct {
	db *gorm.DB
}

func (self *noteService) Create(ctx context.Context, req *pb.NoteReq) (*pb.Note, error) {
	return &pb.Note{
		Id:    123,
		Title: "Todo 123",
	}, nil
}

func (self *noteService) Find(ctx context.Context, req *pb.NoteFindReq) (*pb.Note, error) {
	m := &model.Note{}
	self.db.Find(m, "id = ?", req.Id)
	note := &pb.Note{
		Id:    int32(m.ID),
		Title: m.Title,
		CreatedAt: &google_protobuf.Timestamp{
			Seconds: m.CreatedAt.Unix(),
		},
	}
	return note, nil
}

func main() {
	// 1. Listen/Open a TPC connect at port
	lis, _ := net.Listen("tcp", port)
	// 2. Tao server tu GRP
	grpcServer := grpc.NewServer()
	// 3. Map service to server
	db, _ := gorm.Open("mysql", "default:secret@/notes?charset=utf8&parseTime=True&loc=Local")
	service := &noteService{
		db: db,
	}
	pb.RegisterNoteServiceServer(grpcServer, service)
	// 4. Binding port
	grpcServer.Serve(lis)
}
