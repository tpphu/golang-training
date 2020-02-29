package helper

import (
	"errors"
	"reflect"
)

type HelperErr struct {
	Code    int
	Message string
}

var InvalidInputError error = errors.New("Invalid input")
var InvalidKindError error = errors.New("Input should not be interface")

func Max(list interface{}) (interface{}, error) {
	typeOfList := reflect.TypeOf(list)
	if typeOfList.Kind() != reflect.Array && typeOfList.Kind() != reflect.Slice {
		return nil, InvalidInputError
	}

	// kind of item
	kindOfElem := typeOfList.Elem().Kind()
	if kindOfElem == reflect.Interface {
		return nil, InvalidKindError
	}

	valueOfList := reflect.ValueOf(list)
	max := valueOfList.Index(0) //reflect.Value

	for i := 1; i < valueOfList.Len(); i++ {
		if kindOfElem == reflect.Int {
			if max.Int() < valueOfList.Index(i).Int() {
				max = valueOfList.Index(i)
			}
		} else if kindOfElem == reflect.Float32 {
			if max.Float() < valueOfList.Index(i).Float() {
				max = valueOfList.Index(i)
			}
		}
	}
	return max.Interface(), nil // convert to interface{}
}
