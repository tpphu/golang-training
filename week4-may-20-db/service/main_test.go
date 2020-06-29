package main

import (
	"context"
	"testing"

	product "../proto"
	repo "./mock"
)

// Code test dc la: testable
// Dependency Inversion Prinple/Giam su phu thuoc
// Code nhung gi ma minh viet ra
// SOLID:
func TestGet(t *testing.T) {

	productRepo := repo.ProductRepoImp{}
	service := productService{
		productRepository: productRepo,
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

func TestGetWithNotFound(t *testing.T) {

	productRepo := repo.ProductRepoImp{}
	service := productService{
		productRepository: productRepo,
	}
	req := product.GetReq{
		Id: 2,
	}
	_, err := service.Get(context.TODO(), &req)
	if err == nil {
		t.Error("Product should not found")
	}

}
