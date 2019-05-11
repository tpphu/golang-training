package helper

import (
	"reflect"
)

// if len(arr) > 0 {
//	Max(arr)
// }
func Max(array interface{}) interface{} {
	valueOfArray := reflect.ValueOf(array)
	typeOfArray := reflect.TypeOf(array)
	if typeOfArray.Kind() != reflect.Array && typeOfArray.Kind() != reflect.Slice {
		panic("input should be a slice or array")
	}
	if valueOfArray.Len() == 0 {
		panic("input array not not empty")
	}

	max := valueOfArray.Index(0)
	for i := 1; i < valueOfArray.Len(); i++ {
		// Doi sang dang switch case
		if typeOfArray.Elem().Kind() == reflect.Int {
			if max.Int() < valueOfArray.Index(i).Int() {
				max = valueOfArray.Index(i)
			}
		} else if typeOfArray.Elem().Kind() == reflect.Uint64 ||
			typeOfArray.Elem().Kind() == reflect.Uint32 ||
			typeOfArray.Elem().Kind() == reflect.Uint16 ||
			typeOfArray.Elem().Kind() == reflect.Uint8 {
			if max.Uint() < valueOfArray.Index(i).Uint() {
				max = valueOfArray.Index(i)
			}
		} else if typeOfArray.Elem().Kind() == reflect.Float32 {
			if max.Float() < valueOfArray.Index(i).Float() {
				max = valueOfArray.Index(i)
			}
		}
	}
	return max.Interface()
}

// func MaxInt(array []int) (max int) {
// 	if len(array) == 0 {
// 		panic("input array is not valid")
// 	}
// 	max = array[0]
// 	for i := 1; i < len(array); i++ {
// 		if max < array[i] {
// 			max = array[i]
// 		}
// 	}
// 	return max
// }

// func MaxFloat32(array []float32) (max float32) {
// 	if len(array) == 0 {
// 		panic("input array is not valid")
// 	}
// 	max = array[0]
// 	for i := 1; i < len(array); i++ {
// 		if max < array[i] {
// 			max = array[i]
// 		}
// 	}
// 	return max
// }
