package main

import "fmt"

func main() {
	var x interface{} = true
	x = "3"
	fmt.Println(x)
	m := min(7, 8)
	fmt.Println("min", m)
}

func min(n1 int, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}
