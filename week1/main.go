package main

import (
	"fmt"
	"phudt/week1/helper"
	"strconv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Chuong trinh ket thuc voi error", err)
		}
	}()
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()

	fmt.Println("Hello world")

	var n int = 10
	m := 10
	s := "abc"
	var isSystem bool
	var number int
	var s1 string
	fmt.Println(n, m, s, isSystem, number, s1)

	aNumber := "10"
	n, err := strconv.Atoi(aNumber)
	if err != nil {
		panic(err)
		// fmt.Println("strconv | Have error: ", err)
	}

	// https://github.com/golangci/golangci-lint
	max, err := helper.FindMax([]int{1, 2, 3, 4}...)
	if err != nil {
		fmt.Println("Have error: ", err)
		return
	}

	fmt.Println("Max = ", max)
	fmt.Println("Pi = ", helper.Pi)
	defer func() {
		fmt.Println("defer 3")
	}()
}
