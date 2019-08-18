package model

import "time"

type Voucher struct {
	Id       int
	Code     string
	Discount float32
	Start    time.Time
	End      time.Time
}
