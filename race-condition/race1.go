package main

import (
	"fmt"
)

func main() {
	fmt.Println("------------------")
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
	}()
	m["2"] = "b" // Second conflicting access.
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// go build -o race1 race1.go && for i in {1..100}; do ./race1; done;
// go build -o race1 race1.go && for i in {1..100}; do ./race1 >> race1.txt; done;
