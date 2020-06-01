package main

import "fmt"

func device(n int, d int) int {
	return n / d
}
func main() {
	f()
}
func d() {
	fmt.Println("Defer 2")
}
func f() {
	// Cuc ki hay ho, sau lam ban se thay rat thich
	// Vi dieu hon nhung ngon ngu khac
	defer func() {
		// Neu co exception
		err := recover()
		if err != nil {
			// catch lam cai gi do
			fmt.Println("Chuong trinh da co loi gi do:", err)
		}
		// finally chung ta lam gi do
		fmt.Println("Chuong trinh ket thuc")
	}()
	defer d()
	defer func() {
		fmt.Println("Defer 3")
	}()
	//
	r := device(10, 0)
	fmt.Println("Result of 10/2=", r)
}
