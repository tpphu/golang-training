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

func Contains(list interface{}, v interface{}) bool {
	// 1. Can dung reflect de convert interface truyen vao
	// thanh doi tuong reflect.Value
	listValue := reflect.ValueOf(list)
	value := reflect.ValueOf(v)
	// 2. Can dung reflect de convert interface truyen vao
	// thanh doi tuong reflect.Type
	listType := listValue.Type()
	// 3. Ban dau ket qua dc gan la false
	// Muc dich cho nay la de viet Unit Test
	result := false
	// 4. Voi reflect.Type chung ta dung method Kind de biet
	// kieu list la gi
	switch listType.Kind() {
	case reflect.Slice: //4.1 Slice or Array
		// Vi kieu du lieu la slice nen dung ham .Len o day dc
		for i := 0; i < listValue.Len(); i++ {
			// 4.2 Lay gia tri tai vi tri thu i (reflect.Type)
			// Sau do get ra kieu interface{} de compare voi gia tri truyen vao
			if reflect.DeepEqual(listValue.Index(i).Interface(), v) {
				result = true
				break
			}
		}
		break
	case reflect.String: //4.2 String
		result = strings.Contains(listValue.String(), value.String())
		break
	}
	// 5. Tra ve ket qua
	return result
}
