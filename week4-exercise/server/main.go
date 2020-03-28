package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/tpphu/week4-exercise/model"
	pb "github.com/tpphu/week4-exercise/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type noteService struct {
	db *gorm.DB
}

func (n noteService) Create(ctx context.Context, req *pb.NoteReq) (*pb.Note, error) {
	// 1. Phai tao ra cai model, fill du lieu tu cai request vao model
	// Tai sao noteModel phai dinh nghia nhu la mot pointer
	noteModel := &model.Note{
		Title:     req.Title,
		Completed: req.Completed,
	}
	// 2. Insert vao db
	n.db.Create(noteModel) // truyen vao phai la pointer
	// 3. Fill out ra cai model da insert vao db de ra cai Res
	noteRes := &pb.Note{
		Id:        int32(noteModel.ID),
		Title:     noteModel.Title,
		Completed: noteModel.Completed,
	}
	// Tra ve
	return noteRes, nil
}
func (n noteService) Find(ctx context.Context, req *pb.NoteFindReq) (*pb.Note, error) {
	noteModel := &model.Note{}
	n.db.Find(noteModel, "id = ?", req.Id) // truyen vao phai la pointer
	if noteModel.ID == 0 {
		return nil, errors.New("id is not exist")
	}
	noteRes := &pb.Note{
		Id:        int32(noteModel.ID),
		Title:     noteModel.Title,
		Completed: noteModel.Completed,
		CreatedAt: &timestamp.Timestamp{Seconds: noteModel.CreatedAt.Unix()},
		UpdatedAt: &timestamp.Timestamp{Seconds: noteModel.UpdatedAt.Unix()},
	}
	return noteRes, nil
}
func (n noteService) Update(ctx context.Context, req *pb.NoteUpdateReq) (*pb.Note, error) {
	noteModel := &model.Note{
		Title:     req.Title,
		Completed: req.Completed,
	}

	// Viet chua dung, vi no update du ca ID
	n.db.Model(model.Note{}).Where("id= ?", req.Id).Updates(noteModel)
	noteRes := &pb.Note{}
	noteModel.Populate(noteRes)
	return noteRes, nil
}

func (n noteService) Delete(ctx context.Context, req *pb.NoteDeleteReq) (*pb.NoteDeleteRes, error) {
	return nil, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := gorm.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		panic(err)
	}
	// 1. Listen/Open a TPC connect at port
	lis, _ := net.Listen("tcp", port)
	// 2. Tao server tu GRP
	grpcServer := grpc.NewServer()
	// 3. Map service to server
	pb.RegisterNoteServiceServer(grpcServer, &noteService{
		db: db,
	})
	// 4. Binding port
	fmt.Println("Start service")
	grpcServer.Serve(lis)
}
