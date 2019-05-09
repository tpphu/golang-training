package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p1 := Point{10, 11}
	fmt.Println(p1)

	p2 := p1
	p2.X = -9

	fmt.Println(p1)
}
