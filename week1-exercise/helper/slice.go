package helper

import "reflect"

// Last gets the last element of array.
func Last(arr interface{}) interface{} {
	value := reflect.ValueOf(arr)

	return value.Index(value.Len() - 1).Interface()
}

func Filter(collection interface{}, predicate interface{}) interface{} {
	valueOfPredicate := reflect.ValueOf(predicate)
	typeOfPredicate := reflect.TypeOf(predicate)
	valueOfCollection := reflect.ValueOf(collection)
	typeOfCollection := reflect.TypeOf(collection)
	out := reflect.MakeSlice(typeOfCollection, 0, 0)
	if typeOfPredicate.Kind() == reflect.Func {
		for i := 0; i < valueOfCollection.Len(); i++ {
			elm := valueOfCollection.Index(i)
			in := []reflect.Value{elm}
			result := valueOfPredicate.Call(in)[0]
			if result.Bool() == true {
				out = reflect.Append(out, elm)
			}
		}
	}
	return out.Interface()
}
