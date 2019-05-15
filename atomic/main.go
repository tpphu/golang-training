package main

import "fmt"
import "runtime"
import "sync/atomic"

func main() {
	var a, b int32 = 0, 0

	go func() {
		atomic.StoreInt32(&a, 789)
		atomic.AddInt32(&b, 123)
		runtime.Gosched()
	}()

	for {
		if atomic.LoadInt32(&b) == 123 {
			fmt.Println(atomic.LoadInt32(&a)) // always print 789?
			return
		}
	}

	select {}
}
