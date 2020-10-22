package main

import (
	"fmt"
)

func main() {
	n := 10
	p := &n
	fmt.Println("[MAIN.1] Address of &p: ", &p)
	fmt.Println("[MAIN.1] Address of p refer to:", p)
	fmt.Println("[MAIN.1] Value of p refer to : ", *p)
	f(p)
	fmt.Println("[MAIN.2] Address of &p: ", &p)
	fmt.Println("[MAIN.2] Address of p refer to:", p)
	fmt.Println("[MAIN.2] Value of p refer to : ", *p)
}

func f(p *int) {
	fmt.Println("[f.1] Address of &p: ", &p)
	fmt.Println("[f.1] Address of p refer to:", p)
	fmt.Println("[f.1] Value of p refer to : ", *p)
	m := 9
	p = &m
	fmt.Println("[f.2] Address of &p: ", &p)
	fmt.Println("[f.2] Address of p refer to:", p)
	fmt.Println("[f.2] Value of p refer to : ", *p)
}
