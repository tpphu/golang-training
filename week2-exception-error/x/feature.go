package x

import "fmt"

// Parralel
// Concurrent
func f1() {
	for i := 1; i < 100; i = i + 2 {
		fmt.Println("[f1] = ", i)
	}
}
func f2() {
	for i := 0; i < 100; i = i + 2 {
		fmt.Println("[f2] = ", i)
		if i >= 50 {
			return
		}
	}
}
