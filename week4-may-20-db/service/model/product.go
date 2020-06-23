package model

import (
	product "../../proto"
)

type Product struct {
	Id    int32 `gorm:"primary_key"`
	Sku   string
	Name  string
	Price float32
	Qty   int32
}

// 1. Neu get/set qua met thi viet function
// 2. Neu ma van met nua thi reflect, de hieu struct va gan vao
// 3. Neu qua lam bieng => thi plugin

func (p *Product) Set(in *product.AddReq) {
	p.Sku = in.Sku
	p.Name = in.Name
	p.Price = in.Price
	p.Qty = in.Qty
}

func (p *Product) Fill(in *product.Product) {
	// Co nhung function de copy de dang hon
	in.Id = p.Id
	in.Sku = p.Sku
	in.Name = p.Name
	in.Price = p.Price
	in.Qty = p.Qty
}
