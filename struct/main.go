package main

import "fmt"

type A struct {
	X int
	Y int
}

type B struct {
	A
	Z int
}

func update1(a *A) {
	a.X = 1
	a.Y = 2
}

func update2(a A) {
	a.X = 1
	a.Y = 2
}

func update3(someone *int) {
	*someone = 7
}

func case1() {
	b := B{}
	update1(&b.A)
	fmt.Print(b)
}

func case2() {
	b := B{}
	update2(b.A)
	fmt.Print(b)
}

func case3() {
	b := B{}
	update3(&b.Z)
	fmt.Print(b)
}

func main() {
	case1()
	case2()
	case3()
}
