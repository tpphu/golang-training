package helper

import "reflect"

func Max(list interface{}) interface{} {
	listValue := reflect.ValueOf(list)
	listType := listValue.Type()
	listKind := listType.Kind()

	if listKind != reflect.Array && listKind != reflect.Slice {
		panic("list should array or slice")
	}
	// Cho nay lay Element va get Type
	elementType := listType.Elem()
	elementKind := elementType.Kind()

	max := listValue.Index(0)
	for i := 1; i < listValue.Len(); i++ {
		switch elementKind {
		case reflect.Int:
			if max.Int() < listValue.Index(i).Int() {
				max = listValue.Index(i)
			}
		case reflect.Int32:
			if max.Int() < listValue.Index(i).Int() {
				max = listValue.Index(i)
			}
		}

	}
	return max.Interface()
}

type convert1 func(int) int

// func Map2(list []int, fn convert1) []int {
// 	result := make([]int, len(list))
// 	for i := 0; i < len(list); i++ {
// 		result[i] = fn(list[i])
// 	}
// 	return result
// }
// list := []int{1, 2, 3}
func Map2(list interface{}, fn interface{}) interface{} {
	// 1. Slice or Array
	// 1.1 Lay value of interface
	listValue := reflect.ValueOf(list)
	// 1.2 Lay cai
	listType := listValue.Type()
	listKind := listType.Kind()

	if listKind != reflect.Slice && listKind != reflect.Array {
		panic("list should be slice or array")
	}

	// 2
	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()
	outType := fnType.Out(0)

	// 3. Tao ra ra retResult
	// 3.1 Tao ra cai sliceType
	outSliceType := reflect.SliceOf(outType)
	retResult := reflect.MakeSlice(outSliceType, 0, 0)

	// 4. Loop qua list
	for i := 0; i < listValue.Len(); i++ {
		item := listValue.Index(i)
		// 4.1 Tao input cho func
		in := []reflect.Value{item}
		// 4.2 Call function
		result := fnValue.Call(in)[0]
		// 4.3 Append vao ket qua tu buoc 3 tao ra
		retResult = reflect.Append(retResult, result)
	}
	// 5. Tra ve du lieu
	return retResult.Interface()
}

// list := []interface{}{1, false, "phu"}
// func Map3(list []interface{}, fn interface{}) interface{} {

// 	return nil
// }
