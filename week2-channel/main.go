package main

import (
	"fmt"
	"time"
)

// Parralel
// Concurrent
func f1() {
	for i := 1; i < 100; i = i + 2 {
		// fmt.Println("[f1] = ", i)
		if i >= 50 {
			return
		}
	}
}

func f2() {
	for i := 0; i < 100; i = i + 2 {
		// fmt.Println("[f2] = ", i)
		if i >= 50 {
			return
		}
	}
}

// Truong hop 1, su dung wait group
func main() { //go routine main

	// Channel dc hieu nhu mot cai queue
	// Unbuffer Channel = no empty moi dc day vao
	ch := make(chan int)
	go func() {
		// Goi vao queue 1 gia tri de bao la function 1 xong
		defer func() {
			fmt.Println("Goi dc o day 1")
			ch <- 1
			fmt.Println("Khong Goi dc o day 1")
		}()
		// Thuc thi function 1
		f1()
	}()

	go func() {
		// Goi vao queue 1 gia tri de bao la function 2 xong
		defer func() {
			fmt.Println("Goi dc o day 2")
			ch <- 1
			fmt.Println("Khong Goi dc o day 2.1")
			ch <- 1
			fmt.Println("Khong Goi dc o day 2.2")
		}()
		// Thuc thi function 2
		// cancel cai function nay
		f2()
	}()

	<-ch
	fmt.Println("Lay ra mot phan tu")
	// Lay ra phan tu thu 2
	<-ch
	// Neu lay ra dc 2 phan tu thi xem nhu chuong trinh xong
	time.Sleep(10 * time.Second)
	fmt.Println("Cac f1, va f2 da xong. Chuong trinh ket thuc")

}
