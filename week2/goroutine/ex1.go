package goroutine

import (
	"sync"
)

func Sum(arr []int) int {
	part1 := arr[:len(arr)/2]
	part2 := arr[len(arr)/2:]
	var sum1 int
	var sum2 int
	wg := sync.WaitGroup{}
	// Group 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		sum1 = sum(part1)
	}()
	// Group 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		sum2 = sum(part2)
	}()
	wg.Wait()
	return sum1 + sum2
}

func sum(arr []int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}
