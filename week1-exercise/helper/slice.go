package helper

import (
	"fmt"
	"reflect"
)

// Last gets the last element of array.
func Last(collection interface{}) interface{} {
	typeOfCollection := reflect.TypeOf(collection)
	if typeOfCollection.Kind() != reflect.Array && typeOfCollection.Kind() != reflect.Slice {
		panic("collection should be array or slice")
	}
	valueOfCollection := reflect.ValueOf(collection)

	return valueOfCollection.Index(valueOfCollection.Len() - 1).Interface()
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
	} else if typeOfPredicate.Kind() == reflect.String {
		for i := 0; i < valueOfCollection.Len(); i++ {
			elm := valueOfCollection.Index(i)
			field := elm.FieldByName(valueOfPredicate.String())
			if field.Bool() == true {
				out = reflect.Append(out, elm)
			}
		}
	} else if typeOfPredicate.Kind() == reflect.Array || typeOfPredicate.Kind() == reflect.Slice {
		// fieldName := valueOfPredicate.Index(0).String() //"<interface {} Value>"
		fieldName := valueOfPredicate.Index(0).Interface()
		fieldValue := valueOfPredicate.Index(1).Interface()
		fmt.Println("fieldName:", fieldName)
		for i := 0; i < valueOfCollection.Len(); i++ {
			elm := valueOfCollection.Index(i)
			field := elm.FieldByName(fieldName.(string))
			if field.Interface() == fieldValue {
				out = reflect.Append(out, elm)
			}
		}
	}
	return out.Interface()
}
