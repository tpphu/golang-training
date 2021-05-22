package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	chResult := make(chan int) // un bufferred channel
	// ctx, cancel := context.WithCancel(context.Background(), 1*time.Second)
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	go func() {
		// defer cancel()
		delivery(ctx, chResult)
	}()
	go func() {
		total := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("total :", total)
				return
			case v := <-chResult:
				total += v
			}
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Bye bye nha")
}

func delivery(_ context.Context, chResult chan int) {
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i += 2 {
		wg.Add(1)
		go func(a, b int) {
			defer wg.Done()
			chResult <- sum(a, b)
		}(i, i+1)
	}
	wg.Wait()
}

// func main() {
// 	chResult := make(chan int) // un bufferred channel
// 	wg := sync.WaitGroup{}
// 	for i := 1; i <= 10; i += 2 {
// 		wg.Add(1)
// 		go func(a, b int) {
// 			defer wg.Add(-1)
// 			chResult <- sum(a, b)
// 		}(i, i+1)
// 	}
// 	wg.Wait()
// 	close(chResult)
// 	go func() {
// 		total := 0
// 		for v := range chResult {
// 			total += v
// 			fmt.Println("total:", total)
// 		}
// 	}()
// 	time.Sleep(3 * time.Second)
// }

// sum is illustrate a very long function
func sum(a int, b int) int {
	time.Sleep(2 * time.Second)
	return a + b
}
