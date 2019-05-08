package helper

import (
	"reflect"
	"strings"
)

func Contains(list interface{}, v interface{}) bool {
	valueOfList := reflect.ValueOf(list)
	if valueOfList.Kind() == reflect.Array || valueOfList.Kind() == reflect.Slice {
		type1 := reflect.TypeOf(list).Elem().Kind()
		type2 := reflect.TypeOf(v).Kind()
		if type1 != type2 {
			// panic("Not same type of list and value lookup")
			return false
		}
		for i := 0; i < valueOfList.Len(); i++ {
			if valueOfList.Index(i).Interface() == v {
				return true
			}
		}
	}
	if valueOfList.Kind() == reflect.String {
		strValue := valueOfList.String()
		if strings.Index(strValue, v.(string)) >= 0 {
			return true
		}
	}
	return false
}
