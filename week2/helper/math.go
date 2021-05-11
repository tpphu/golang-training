package helper

type Persion interface {
	GetName() string
	GetAge() int
}

func Max(n ...interface{}) interface{} {
	for i := 0; i < len(n); i++ {
		v := n[i]
	}
	return nil
}

type Fn func(int) int

func demo() {
	ch := make(chan Person, 10)
	ch <- 1
	n := <-ch // Co thang day vao, hoac la channel bi closed
}

func demo2() {
	ch := make(chan Fn, 10)
	ch <- 1
	n := <-ch
}

// func demo() {
// 	ch := make(chan URL, 10)
// 	ch := make(chan Data, 10)
// 	ch <- 1
// 	n := <-ch
// 	if n == 1 {

// 	}
// }

// func Max2(n1 interface{}, n2 interface{}) interface{} {
// 	return nil
// }

// func MaxInt(n1 int, n2 int) int {
// 	if n1 > n2 {
// 		return n1
// 	}
// 	return n2
// }

// func MaxInt64(n1 int64, n2 int64) int64 {
// 	if n1 > n2 {
// 		return n1
// 	}
// 	return n2
// }
