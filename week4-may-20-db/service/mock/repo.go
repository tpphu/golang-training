package mock

import (
	"errors"

	"../model"
)

type ProductRepoImp struct {
}

func (repo ProductRepoImp) Find(id int32) (*model.Product, error) {
	if id == 2 {
		return nil, errors.New("Not found")
	}
	product := &model.Product{
		Id: id,
	}
	return product, nil
}

func (repo ProductRepoImp) Create(product *model.Product) error {
	//@TOTO Can implement cho nay
	return nil
}
