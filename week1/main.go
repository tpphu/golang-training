package main

import (
	"encoding/json"
	"fmt"
	"os"
	"phudt/week1/helper"
	"phudt/week1/model"
	"strconv"

	yaml "gopkg.in/yaml.v3"
)

type Person struct {
	Id          int
	Name        string
	YearOfBirth int
}

func main() {
	// _, err := os.Open("input.csv")
	file, err := os.Open("data/input.csv")
	// defer file.Close()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("File name", file.Name())
	data := make([]byte, 1000)
	n, err := file.Read(data)
	fmt.Println(n, err)
	fmt.Println("data:", string(data))
	list_of_person := []Person{Person{
		Id:          1,
		Name:        "Phu 1",
		YearOfBirth: 1986,
	}, Person{
		Id:          2,
		Name:        "Phu 2",
		YearOfBirth: 1987,
	}, Person{
		Id:          3,
		Name:        "Phu 3",
		YearOfBirth: 1988,
	}}
	out, _ := yaml.Marshal(list_of_person)
	fmt.Println(string(out))
}

func main2() {
	var n int = 10
	fmt.Printf("(co dinh) value of n: %v \n", n)
	fmt.Printf("(tuy may, tuy thoi diem) address of n: %v \n", &n)
	var p *int = &n
	fmt.Printf("(co dinh) value of p: %v \n", p)
	fmt.Printf("(tuy may, tuy thoi diem) address of p: %v \n", &p)
	fmt.Printf("(gia tri ma bien p dang refer toi) *value of p: %v \n", *p)
	n = -9
	fmt.Printf("(gia tri ma bien p dang refer toi) *value of p: %v \n", *p)
	*p = -99
	fmt.Printf("(co dinh) value of n: %v \n", n)
	m := 100
	p = &m
	fmt.Printf("(gia tri ma bien p dang refer toi) *value of p: %v \n", *p)
	n = -99
	fmt.Printf("(gia tri ma bien p dang refer toi) *value of p: %v \n", *p)
}

func week1() {
	fmt.Println("Hello world")
	fmt.Println("Sum of: {1, 2, 3, 4} is ", helper.Sum(1, 2, 3, 4))
	p := model.Person{
		Name:          "Phu",
		Year_of_birth: 1986,
	}
	data, _ := json.Marshal(p)
	fmt.Println("JSON marshal: ", string(data))

	s := "100a000"
	fmt.Println(" len of '%s' = ", len(s))
	var n int
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(" Number of %s is %d\n", s, n)
}
