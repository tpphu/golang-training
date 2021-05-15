package main

import (
	"fmt"
	"time"
)

// Research de to chuc struct sao cho tiet kiem memory
// Struct memory management go
type Person struct {
	Id       int
	Fullname string
	Birthday time.Time
	age      int
}

type HoSoBenhAn struct {
	Name         string
	NgayNhapVien time.Time
	NgayXuatVien time.Time
}

func (h HoSoBenhAn) GetName() string {
	return h.Name
}

type Patient struct {
	Person
	HoSoBenhAn
}

type BenhVien struct {
	Name    string
	Address string
}
type Doctor struct {
	Person
	BenhVien
}

func NewPatient(id int, fullname string, yearOfBirth int) Patient {
	p := Patient{Person{
		Id:       10,
		Fullname: "Duc dep trai",
		Birthday: time.Date(yearOfBirth, time.January, 1, 0, 0, 0, 0, time.UTC),
		age:      -1,
	}, HoSoBenhAn{}}
	return p
}

func (p *Patient) GetAge() int {
	if p.age != -1 {
		return p.age
	}
	now := time.Now()
	p.age = now.Year() - p.Birthday.Year()
	return p.age
}

func (p Person) String() string {
	// Stream And Buffer
	// return "Khong cho phep ban in ra"
	return fmt.Sprintf(`Id = %d,
	Fullname= "%s",
	YearOfBirth = %d`, p.Id, p.Fullname, p.Birthday.Year())
}

func (h Person) GetName() string {
	return h.Fullname
}

// func (h Patient) GetName() string {
// 	return h.HoSoBenhAn.GetName()
// }

func main() {
	// Khong can constructor la nhu vay
	// var p Patient
	// p.Id = 10
	// p.Fullname = "Duc dep trai"
	// p := Patient{
	// 	Id:       10,
	// 	Fullname: "Duc dep trai",
	// 	Birthday: time.Date(1992, time.July, 11, 0, 0, 0, 0, time.UTC),
	// }
	p := NewPatient(10, "Duc dep trai", 1992)
	fmt.Println(p)
	fmt.Println("Duc's age:", p.GetAge())

	fmt.Println("Duc's age:", p.GetAge())
	fmt.Println("ID DUc:", p.Id)

	// fmt.Println("GetName", p.GetName())

	d := Doctor{}
	d.Fullname = "Phu"
	fmt.Println(d)
}
