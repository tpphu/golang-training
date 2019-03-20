package helper

import "reflect"

func Map(list interface{}, iterateeFunc interface{}) interface{} {
	listValue := reflect.ValueOf(list)
	listType := listValue.Type()

	iterateeFuncValue := reflect.ValueOf(iterateeFunc)
	typeOfResult := reflect.SliceOf(iterateeFuncValue.Type().Out(0))
	result := reflect.MakeSlice(typeOfResult, 0, 0)

	listKind := listType.Kind()
	if listKind == reflect.Slice || listKind == reflect.Array {
		for i := 0; i < listValue.Len(); i++ {
			elem := listValue.Index(i)
			in := []reflect.Value{elem}
			out := iterateeFuncValue.Call(in)[0]
			result = reflect.Append(result, out)
		}

	}
	return result.Interface()
}
