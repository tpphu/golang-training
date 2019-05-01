package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	c := make(chan bool, 2)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	go func() {
		m["2"] = "b" // Second conflicting access.
		c <- true
	}()
	<-c
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// go build -o race3 race3.go && for i in {1..100}; do ./race3; done;
// go build -o race3 race3.go && for i in {1..100}; do ./race3 >> race3.txt; done;
