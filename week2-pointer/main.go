package main

import "fmt"

func main() {
	var n int
	n = 7
	fmt.Println("Address of n:", &n)
	fmt.Println("Value of n:", n)
	//
	var p *int
	fmt.Println("Address of p:", &p)
	fmt.Println("Value of p:", p)
	p = &n
	fmt.Println("Value of p:", p)
	fmt.Println("Refer Value of p:", *p)
	*p = 8
	fmt.Println("Refer Value of p:", *p)
	fmt.Println("Value of n:", n)

	//
	a := n
}
