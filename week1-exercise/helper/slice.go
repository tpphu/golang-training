package helper

import "reflect"

// Last gets the last element of array.
func Last(arr interface{}) interface{} {
	value := reflect.ValueOf(arr)

	return value.Index(value.Len() - 1).Interface()
}

func Filter(collection interface{}, predicate interface{}) interface{} {
	// valueOfPredicate := reflect.ValueOf(predicate)
	// typeOfPredicate := reflect.TypeOf(predicate)
	// if typeOfPredicate == reflect.Func {

	// 	valueOfPredicate.Call
	// }
	return nil
}
