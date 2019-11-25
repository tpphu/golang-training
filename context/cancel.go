package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithCancel(context.Background())
	timeout := time.After(5 * time.Second)
	aChan := make(chan int, 3)
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	sum := func(n int) {
		for i := 0; i < 10; i++ {
			time.Sleep(250 * time.Millisecond)
			fmt.Println("n: ", n, " - ", i)
		}
		<-aChan
	}
	i := 0
	for true {
		select {
		// Cho nay cuc ki khong dung
		// Code se khong bao gio reach dc cho nay
		case <-timeout:
			fmt.Println("overslept")
			goto END_MAIN
		// Context Done co the bi goi lai
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			goto END_MAIN
		// Khi code cho nay se khong the dung lai
		case aChan <- time.Now().Second():
			i++
			go sum(i)
		}
	}
END_MAIN:
}
