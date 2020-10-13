package goroutine

import (
	"fmt"
	"sync"
)

func Sum(arr []int) int { // CPU1
	part1 := arr[:len(arr)/2]
	part2 := arr[len(arr)/2:]
	var sum1 *int
	var sum2 *int
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go sum(part1, sum1, wg, "sum1") // CPU2
	go sum(part2, sum2, wg, "sum2") // CPU3
	// CPU1 wait
	wg.Wait()
	if sum1 != nil && sum2 != nil {
		return *sum1 + *sum2
	}
	return 0
}

func sum(arr []int, result *int, wg *sync.WaitGroup, key string) int {
	fmt.Println("key:", key)
	ret := 0
	result = &ret
	for _, v := range arr {
		ret += v
	}
	fmt.Printf("result | %s: %d\n", key, *result)
	defer wg.Done()
	return ret
}
