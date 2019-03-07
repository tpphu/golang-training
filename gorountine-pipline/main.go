package main

import (
	"fmt"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
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
		for v := range out {
			time.Sleep(1 * time.Second)
			fmt.Println(v)
		}
	}()
	// fmt.Println(<-out) // 4
	// fmt.Println(<-out) // 9
	// fmt.Println(<-out) // 16
	time.Sleep(5 * time.Second)

}
