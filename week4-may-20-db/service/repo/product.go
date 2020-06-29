package repo

import (
	"../model"
	"github.com/jinzhu/gorm"
)

type Product interface {
	Find(id int32) (*model.Product, error)
	Create(*model.Product) error
}

type ProductRepoImp struct {
	DB *gorm.DB
}

func (repo ProductRepoImp) Find(id int32) (*model.Product, error) {
	//@TOTO Can implement cho nay
	// @TOTO Neu minh can Unit Test cho ham Find thi minh chi can kiem tra cai cau select co dung khong
	return nil, nil
}

func (repo ProductRepoImp) Create(product *model.Product) error {
	//@TOTO Can implement cho nay
	//@Todo unit test thi can kiem tra cau sql
	return nil
}
