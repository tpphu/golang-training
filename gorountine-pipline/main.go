package main

import (
	"fmt"
	"runtime"
	"time"
)

// https://gobyexample.com/channel-buffering
// Nhan vao mot danh sach cac so nguyen
// Tra ra mot channel
func gen(num int) <-chan int {
	// Run Step 1
	// Unbufferred Channel
	out := make(chan int, 1) //capacity
	// Run Step 2
	go func() {
		for i := 1; i < num; i++ {
			// Chi day dc vao channel khi ma len < capacity
			out <- i
			// In ra o day co nghia la da push dc
			fmt.Printf("\n[GEN] channel: %d", i)
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

func print(in <-chan int) {
	for v := range in {
		time.Sleep(1 * time.Second)
		fmt.Printf("\n[OUT]: %d", v)
		fmt.Println("\n-------------")
	}
}

func main() {
	fmt.Println("Num CPUs:", runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	// Set up the pipeline.
	// Step 1. Tao ra cai input channel
	in := gen(10)
	// Step 2. Dung cai in => tao ra cai out channel
	out := sq(in)
	// Step 3. Dung cai out thanh in cua step print
	go print(out)

	time.Sleep(30 * time.Second)

}
