package helper

import "reflect"

func Find(arr interface{}, predicate interface{}) interface{} {

	iteratorValue := reflect.ValueOf(arr)
	funcValue := reflect.ValueOf(predicate)

	for i := 0; i < iteratorValue.Len(); i++ {
		item := iteratorValue.Index(i)
		in := []reflect.Value{item}
		out := funcValue.Call(in)
		result := out[0].Bool()

		if result == true {
			return item.Interface()
		}
	}
	return nil
}
