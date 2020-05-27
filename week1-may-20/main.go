package main

// built-in
import (
	"fmt"
	"time"
)

// goroutine main
func main() {
	// Thoại: out ra cái stdout, stdin, stderr
	// để nó đọc cái nội dung và show ra màn hình
	fmt.Println("Hello") // write ra cái standard output, stdio
	sum := 0
	fmt.Println("address of sum:=", &sum)
	// goroutine 2
	go func() {
		// CPU 2/1
		for i := 0; i < 100000000; i++ {
			sum = sum + i
		}
	}()
	// goroutine 3
	go func() {
		// CPU 2/3
		for j := 0; j < 100000000; j++ {
			sum = sum + j
		}
	}()
	// CPU 1
	time.Sleep(1 * time.Hour) //Stupid
	fmt.Println("Ket thuc")
	fmt.Println("sum = ", sum)
}
