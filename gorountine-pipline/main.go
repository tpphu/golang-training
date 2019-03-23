package main

import (
	"fmt"
	"time"
)

// Nhan vao mot danh sach cac so nguyen
// Tra ra mot channel
func gen(nums ...int) <-chan int {
	// Run Step 1
	out := make(chan int, 1) //capacity
	// Run Step 2
	go func() {
		for _, n := range nums {
			// Chi day dc vao channel khi ma len < capacity
			fmt.Print("\nKhi nao in ra?\n")
			out <- n
		}
		close(out)
	}()
	// Run Step
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		//fmt.Println("in:", in)
		// for n := range in {
		// 	time.Sleep(time.Second * 1)
		// 	out <- n * n
		// }
		n := <-in
		out <- n * n
		//close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	c := gen(2, 3, 4, 5)
	out := sq(c)

	// Consume the output.
	go func() {
		for v := range out {
			//time.Sleep(1 * time.Second)
			fmt.Println(v)
		}
	}()
	// fmt.Println(<-out) // 4
	// fmt.Println(<-out) // 9
	// fmt.Println(<-out) // 16
	time.Sleep(30 * time.Second)

}
