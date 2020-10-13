package goroutine

import (
	"fmt"
	"sync"
)

func Sum(arr []int) int { // CPU1
	part1 := arr[:len(arr)/2]
	part2 := arr[len(arr)/2:]
	var sum1 int
	var sum2 int
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go sum(part1, &sum1, wg, "sum1") // CPU2
	go sum(part2, &sum2, wg, "sum2") // CPU3
	// CPU1 wait
	wg.Wait()
	// if sum1 != nil && sum2 != nil {
	// 	return *sum1 + *sum2
	// }
	// return 0
	return sum1 + sum2
}

func sum(arr []int, result *int, wg *sync.WaitGroup, key string) int {
	defer wg.Done()
	defer func() {
		fmt.Printf("result | %s | 1\n", key)
	}()
	defer func() {
		fmt.Printf("result | %s | 2\n", key)
	}()
	fmt.Println("key:", key)
	ret := 0
	for _, v := range arr {
		ret += v
	}
	if ret > 10 {
		return 0
	}
	*result = ret
	fmt.Printf("result | %s: %d\n", key, *result)
	return ret
}
