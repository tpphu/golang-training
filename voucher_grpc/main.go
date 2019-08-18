package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"./proto"
	"./storage"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

type voucherServiceImp struct {
	DB *sql.DB
}

const UNKOWN_ERR = 1
const EXIST_ERR = 400

func (s *voucherServiceImp) Register(ctx context.Context, req *proto.VoucherReq) (*proto.VoucherRes, error) {
	fmt.Println("Register | 1")
	voucherStorage := storage.Voucher{
		DB: s.DB,
	}
	res := &proto.VoucherRes{}
	isExist, err := voucherStorage.RegisterIsolation(req)
	if err != nil {
		res.Error = &proto.Error{
			Code:    UNKOWN_ERR,
			Message: err.Error(),
		}
	}
	if isExist == true {
		res.Error = &proto.Error{
			Code:    EXIST_ERR,
			Message: "Code is exist",
		}
	}
	res.Data = &proto.Voucher{}
	res.Data.Id = 1
	res.Data.Code = req.Code
	res.Data.Discount = req.Discount
	res.Data.Start = req.Start
	res.Data.End = req.End
	return res, err
}

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect Db
	db, err := sql.Open("mysql", "default:secret@/voucher")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(30)
	defer db.Close()
	// GRPC
	// 1. Listen/Open a TPC connect at port
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	// 2. Tao server tu GRP
	grpcServer := grpc.NewServer()
	// 3. Map service to server
	voucherService := &voucherServiceImp{
		DB: db,
	}
	proto.RegisterVoucherServiceServer(grpcServer, voucherService)
	// 4. Binding port
	fmt.Println("Start GRPC on " + port)
	grpcServer.Serve(lis)
}
