package main

import (
	"fmt"
	"sync"
)

// Parralel
// Concurrent
func f1() {
	for i := 1; i < 100; i = i + 2 {
		fmt.Println("[f1] = ", i)
	}
}
func f2() {
	for i := 0; i < 100; i = i + 2 {
		fmt.Println("[f2] = ", i)
		if i >= 50 {
			return
		}
	}
}

wg := sync.WaitGroup{}

// Truong hop 1, su dung wait group
func main() { //go routine main
	wg := sync.WaitGroup{}

	wg.Add(1)
	
	go func() {
		defer wg.Add(-1)
		f1()
	}()

	wg.Add(1)
	go func() {
		defer wg.Add(-1)
		f2()
	}()

	// Wait
	wg.Wait() 
	fmt.Println("Cac f1, va f2 da xong. Chuong trinh ket thuc")

}
