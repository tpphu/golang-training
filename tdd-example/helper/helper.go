package helper

import "reflect"

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
