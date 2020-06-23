package main

import (
	"context"
	"testing"

	product "../proto"

	_ "github.com/go-sql-driver/mysql" // De minh import thu vien cho thang gorm xai
	"github.com/jinzhu/gorm"
)

func TestGet(t *testing.T) {
	// Mock DB
	// Quan trong
	// Cac ban tim hieu
	db, _ := gorm.Open("mysql", "root:root@(127.0.0.1)/gomay20?charset=utf8&parseTime=True&loc=Local")
	service := productService{
		DB: db,
	}
	req := product.GetReq{
		Id: 1,
	}
	res, err := service.Get(context.TODO(), &req)
	if err != nil {
		t.Error(err)
	}
	if res.Product == nil {
		t.Error("Product should not be nil")
	}
	if res.Product.Id != 1 {
		t.Error("Product Id should be 1")
	}
}
