package helper

import (
	"reflect"
)

func IndexOf(arr interface{}, find interface{}) int {
	value := reflect.ValueOf(arr).Index(0).Type()
	switch value.Kind() {
	case reflect.String:
		for i, v := range arr.([]string) {
			if v == find {
				return i
			}
		}
	case reflect.Int32:
		for i, v := range arr.([]int32) {
			if v == find {
				return i
			}
		}
	default:
		panic("Khong biet xu ly sao")
	}
	return -1
}

// IndexOf vs indexOf
func IndexOfInt(arr []int, find int) int {
	for i, v := range arr {
		if v == find {
			return i
		}
	}
	return -1
}

// IndexOfString to search the position of a string in a slice of string.
func IndexOfString(arr []string, find string) int {
	for i, v := range arr {
		if v == find {
			return i
		}
	}
	return -1
}
