package helper

import (
	"reflect"
)

// IsEmpty returns if the object is considered as empty or not.
func IsEmpty(obj interface{}) bool {
	if obj == nil || obj == "" || obj == false || obj == 0 {
		return true
	}
	valueOfObject := reflect.ValueOf(obj)
	if valueOfObject.Kind() == reflect.Slice || valueOfObject.Kind() == reflect.Array {
		if valueOfObject.Len() == 0 {
			return true
		}
	}
	if valueOfObject.Kind() == reflect.Ptr {
		valueOfObject = valueOfObject.Elem()
	}
	if valueOfObject.Kind() == reflect.Struct {
		for i := 0; i < valueOfObject.NumField(); i++ {
			field := valueOfObject.Field(i)
			value := field.Interface()
			if !(value == nil || value == "" || value == false || value == 0) {
				return false
			}
		}
		return true
	}
	return false
}

func ContainsInt(list []int, value int) bool {
	result := false
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			result = true
			// break
		}
	}
	return result
}

func ContainsString(list []string, value string) bool {
	result := false
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			result = true
			// break
		}
	}
	return result
}

// func StringIsEmpty(value string) bool {
// 	if value == "" {
// 		return true
// 	}
// 	return false
// }

// func Int32IsEmpty(value int32) bool {
// 	if value == 0 {
// 		return true
// 	}
// 	return false
// }

// func BooleanIsEmpty(value bool) bool {
// 	if value == false {
// 		return true
// 	}
// 	return false
// }
