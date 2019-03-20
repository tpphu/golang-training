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
