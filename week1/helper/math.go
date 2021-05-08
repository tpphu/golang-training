package helper

import (
	"errors"
)

var Pi float32 = 3.14

// func FindMax(arr []int) (int, error) {
// 	if len(arr) == 0 {
// 		return 0, errors.New("input is invalid")
// 	}
// 	max := arr[0]
// 	for i := 1; i < len(arr); i++ {
// 		if max < arr[i] {
// 			max = arr[i]
// 		}
// 	}
// 	return max, nil
// }

func FindMax(arr ...int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("input is invalid")
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max, nil
}
