package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for true {
		select {
		// Cho nay cuc ki khong dung
		// Code se khong bao gio reach dc cho nay
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
			time.Sleep(500 * time.Millisecond)
			goto END_MAIN
		// Context Done co the bi goi lai
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			time.Sleep(500 * time.Millisecond)
			goto END_MAIN
		// Khi code cho nay se khong the dung lai
		default:
			for i := 0; i < 10; i++ {
				fmt.Println("i: ", i)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
END_MAIN:
}
