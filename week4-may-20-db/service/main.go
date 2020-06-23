package main

import (
	"context"
	"fmt"
	"net"

	product "../proto"
	model "./model"
	_ "github.com/go-sql-driver/mysql" // De minh import thu vien cho thang gorm xai
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

// Nhiem vu that su cua chung
type productService struct {
	DB *gorm.DB
}

func (s *productService) Add(ctx context.Context, req *product.AddReq) (res *product.AddRes, err error) {
	// Fill vao gorm de insert
	p := model.Product{}
	p.Set(req)
	s.DB.Create(&p)
	// Fill nguoc lai cai Product ma ban khai bao trong proto
	productRet := &product.Product{}
	p.Fill(productRet)
	res = &product.AddRes{}
	res.Product = productRet
	return res, nil
}
func (s *productService) Update(context.Context, *product.UpdateReq) (*product.UpdateRes, error) {
	return nil, nil
}
func (s *productService) Delete(context.Context, *product.DeleteReq) (*product.DeleteRes, error) {
	return nil, nil
}
func (s *productService) Get(ctx context.Context, req *product.GetReq) (res *product.GetRes, err error) {
	pDB := model.Product{}
	s.DB.First(&pDB, req.Id)
	//
	productRet := &product.Product{}
	pDB.Fill(productRet)
	//
	res = &product.GetRes{}
	res.Product = productRet
	return res, nil
}

func main() {
	var db *gorm.DB
	var err error
	db, err = gorm.Open("mysql", "root:root@(127.0.0.1)/gomay20?charset=utf8&parseTime=True&loc=Local")
	// if error
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(3) // Neu ma chuong trinh khong su dung connect thi nen dong lai connection do, va chi giu lai 10 cai
	db.DB().SetMaxOpenConns(5) // Neu co nhu cau mo connection nhieu thi chi mo toi da la 20
	// db.DropTableIfExists(&model.Product{})
	db.AutoMigrate(&model.Product{})
	// Nen de duoi sau error
	defer db.Close()
	// Syntax
	port := "0.0.0.0:50001"
	lis, _ := net.Listen("tcp", port) // TPC listener
	grpcServer := grpc.NewServer()
	// Syntax
	service := productService{
		DB: db,
	}
	product.RegisterProductServiceServer(grpcServer, &service)
	// Syntax

	fmt.Println("Start service at :50001")
	grpcServer.Serve(lis)

}
