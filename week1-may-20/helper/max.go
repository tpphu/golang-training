package helper

import (
	"errors"
)

var ErrLenOfListIsEmpty = errors.New("The len of input is empty!")

// func max(list []int) (int, error) {
// 	if len(list) == 0 {
// 		return 0, ErrLenOfListIsEmpty
// 	}
// 	max := list[0]
// 	for i := 1; i < len(list); i++ {
// 		if list[i] > max {
// 			max = list[i]
// 		}
// 	}
// 	return max, nil
// }

func max(list []interface{}) (interface{}, error) {
	if len(list) == 0 {
		return 0, ErrLenOfListIsEmpty
	}
	max := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}
	return max, nil
}