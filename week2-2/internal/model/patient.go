package model

import (
	"fmt"
	"time"
)

// Cach khai bao enum

type Gender int

const (
	Undefined Gender = iota
	Male
	Female
)

type Person struct {
	Id       int64
	Fullname string
	Address  string
	Birthday time.Time
	// Go tag -> rat tuyet voi.
	Gender Gender `gorm:"-" json:"-"`
	// age
	age int
	// custom field
	X string `gorm:"-" json:"-"`
}

type BenhAn struct {
	TenBenh      string `gorm:"-" json:"-"`
	NgayNhapVien time.Time
	NgayXuatVien time.Time
	X            string `gorm:"-" json:"-"`
}

// Cach export field cho Gorm hieu
type Patient struct {
	Person
	BenhAn
}

type Doctor struct {
	Person
	NgayVaoCaTruc time.Time
	NgayRaCaTruc  time.Time
}

// Them mot function cho struct
func (Patient) TableName() string {
	return "patient"
}

// Function the in cai struct theo string
func (p Person) String() string {
	return fmt.Sprintf("%d - %s", p.Id, p.Fullname)
}

// The modify struct Patient
func (p *Person) GetAge() int {
	// return time.Now().Year() - p.Birthday.Year()
	if p.age != -1 {
		fmt.Println("Minh muon lan sau la vao day")
		return p.age
	}
	fmt.Println("Chi muon vao day 1 lan")
	time.Sleep(3 * time.Second)
	p.age = time.Now().Year() - p.Birthday.Year()
	return p.age
}

// New construct
func NewPatient(fullname string, address string) Patient {
	// fmt.Println("Male:", Male)
	// fmt.Println("Female:", Female)
	p := Patient{
		Person: Person{
			Fullname: fullname,
			Address:  address,
			Birthday: time.Date(1997, time.July, 4, 0, 0, 0, 0, time.UTC),
			Gender:   Male,
			age:      -1,
			X:        "BeTram",
		},
		BenhAn: BenhAn{
			NgayNhapVien: time.Now(),
			NgayXuatVien: time.Now(),
			X:            "BeTram",
		},
	}
	fmt.Println("p.NgayNhapVien:", p.NgayNhapVien)
	fmt.Println("p.BenhAn.X:", p.BenhAn.X)
	return p
}
