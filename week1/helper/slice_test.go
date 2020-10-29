package helper

import (
	"fmt"
	"os"
	"testing"
)

/**
-------------------------------------------------------------------------
❯ go test ./week1/helper/... -run ^TestSlice1$  -v
=== RUN   TestSlice
Cap:  0
Len:  0
-----------------
add at element [0]: 0xc0000a4198
Cap:  1
Len:  1
-----------------
add at element [0]: 0xc0000a41b0
Cap:  2
Len:  2
-----------------
add at element [0]: 0xc0000ce060
Cap:  4
Len:  3
-----------------
add at element [0]: 0xc0000ce060
Cap:  4
Len:  4
-----------------
add at element [0]: 0xc0000b21c0
Cap:  8
Len:  5
-----------------
add at element [0]: 0xc0000b21c0
Cap:  8
Len:  6
-----------------
add at element [0]: 0xc0000b21c0
Cap:  8
Len:  7
-----------------
add at element [0]: 0xc0000b21c0
Cap:  8
Len:  8
--- PASS: TestSlice (0.00s)
PASS
ok      github.com/tpphu/golang-training/week1/helper   0.005s
-----------------------------------------------------------------------*/
func TestSlice1(t *testing.T) {
	arr := []int{}
	print(arr)
	for i := 0; i < 8; i++ {
		fmt.Println("-----------------")
		arr = append(arr, i)
		fmt.Printf("add at element [0]: %p\n", &arr[0])
		print(arr)
	}
}

/**
|-------------------------------------------------------------------------

❯ go test ./week1/helper/... -run ^TestSlice2$  -v
=== RUN   TestSlice2
Cap:  8
Len:  0
-----------------
add at element [0]: 0xc00001a440
Cap:  8
Len:  1
-----------------
add at element [0]: 0xc00001a440
Cap:  8
Len:  2
-----------------
add at element [0]: 0xc00001a440
Cap:  8
Len:  3
-----------------
add at element [0]: 0xc00001a440
Cap:  8
Len:  4
-----------------
add at element [0]: 0xc00001a440
Cap:  8
Len:  5
-----------------
add at element [0]: 0xc00001a440
Cap:  8
Len:  6
-----------------
add at element [0]: 0xc00001a440
Cap:  8
Len:  7
-----------------
add at element [0]: 0xc00001a440
Cap:  8
Len:  8
--- PASS: TestSlice2 (0.00s)
PASS
ok      github.com/tpphu/golang-training/week1/helper   0.005s

|-----------------------------------------------------------------------*/
func TestSlice2(t *testing.T) {
	arr := make([]int, 0, 8)
	print(arr)
	cwd, _ := os.Getwd()
	fmt.Println("cwd", cwd)
	for i := 0; i < 8; i++ {
		fmt.Println("-----------------")
		arr = append(arr, i)
		fmt.Printf("add at element [0]: %p\n", &arr[0])
		print(arr)
	}
}

func print(arr []int) {
	fmt.Println("Cap: ", cap(arr))
	fmt.Println("Len: ", len(arr))
}
