package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan int)

	go func() {
		ch <- 1000
		fmt.Println("hello 1")
	}()

	go func() {
		fmt.Println("hello 2")
		time.Sleep(time.Second * 1)
		<-ch
		time.Sleep(time.Second * 1)
		done <- 1
	}()

	<-done
}
