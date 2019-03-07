package helper

import "reflect"

// Last gets the last element of array.
func Last(arr interface{}) interface{} {
	value := reflect.ValueOf(arr)

	return value.Index(value.Len() - 1).Interface()
}
