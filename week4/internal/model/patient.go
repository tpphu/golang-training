package model

type Patient struct {
	Id       int64
	Fullname string
	Address  string
	Birthday string
	Gender   int32
	Age      int32
}

func (Patient) TableName() string {
	return "patient"
}
