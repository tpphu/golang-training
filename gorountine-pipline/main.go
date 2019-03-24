package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/channel-buffering
// Nhan vao mot danh sach cac so nguyen
// Tra ra mot channel
func gen(nums ...int) <-chan int {
	// Run Step 1
	// Bufferred Channel
	out := make(chan int, 4) //capacity
	// Run Step 2
	go func() {
		for _, n := range nums {
			// Chi day dc vao channel khi ma len < capacity
			out <- n
			// In ra o day co nghia la da push dc
			fmt.Printf("\n[GEN] channel: %d", n)
		}
		close(out)
	}()
	// Run Step
	return out
}

func sq(in <-chan int) <-chan int {
	// Unbufferred Channel
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
			fmt.Printf("\n[SQ] channel: %d", n*n)
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	c := gen(2, 3, 4, 5)
	out := sq(c)

	// Consume the output.
	go func() {
		// Vi out la Unbufferred Channel
		// nen khi sleep, thi func sq khong the push them dc value vao sq channel
		// do vay ban se khong thay dong in ra
		for v := range out {
			time.Sleep(1 * time.Second)
			fmt.Printf("\n[OUT]: %d", v)
		}
	}()
	time.Sleep(30 * time.Second)

}
