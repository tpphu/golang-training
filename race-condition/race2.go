package main

import (
	"fmt"
)

func main() {
	fmt.Println("------------------")
	m := make(map[string]string)
	c := make(chan bool)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// go build -o race2 race2.go && for i in {1..100}; do ./race2; done;
// go build -o race2 race2.go && for i in {1..100}; do ./race2 >> race2.txt; done;
