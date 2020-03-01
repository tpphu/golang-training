package helper

import "reflect"

func IntersectionInt(allListInput ...interface{}) (ret interface{}, err error) {
	n := len(allListInput)
	ret = allListInput[0]
	for i := 1; i < n; i++ {
		ret, _ = IntersectionOfTwo_2(ret, allListInput[i])
	}
	return ret, err
}

func IntersectionOfTwo_2(lista interface{}, listb interface{}) (ret interface{}, err error) {
	typeOfListA := reflect.TypeOf(lista)
	if typeOfListA.Kind() != reflect.Slice && typeOfListA.Kind() != reflect.Array {
		return ret, InvalidInputError
	}
	typeOfListB := reflect.TypeOf(listb)
	if typeOfListB.Kind() != reflect.Slice && typeOfListB.Kind() != reflect.Array {
		return ret, InvalidInputError
	}

	valueOfListA := reflect.ValueOf(lista)
	valueOfListB := reflect.ValueOf(listb)

	kindOfInputA := typeOfListA.Elem().Kind()
	kindOfInputB := typeOfListB.Elem().Kind()
	if kindOfInputA != kindOfInputB {
		return ret, InvalidInputError
	}

	cap := valueOfListA.Len()
	if cap > valueOfListB.Len() {
		cap = valueOfListB.Len()
	}

	ptr := reflect.New(typeOfListA)
	ptr.Elem().Set(reflect.MakeSlice(typeOfListA, 0, cap))
	offset := 0

	for i := 0; i < valueOfListA.Len(); i++ {
		elem := valueOfListA.Index(i)
		for j := 0; j < valueOfListB.Len(); j++ {
			if elem.Interface() == valueOfListB.Index(j).Interface() {
				ptr.Elem().SetLen(offset + 1)
				ptr.Elem().Index(offset).Set(elem)
				offset++
			}
		}
	}
	return ptr.Elem().Interface(), nil
}

func IntersectionOfTwo(lista interface{}, listb interface{}) (ret interface{}, err error) {
	typeOfListA := reflect.TypeOf(lista)
	if typeOfListA.Kind() != reflect.Slice && typeOfListA.Kind() != reflect.Array {
		return ret, InvalidInputError
	}
	typeOfListB := reflect.TypeOf(listb)
	if typeOfListB.Kind() != reflect.Slice && typeOfListB.Kind() != reflect.Array {
		return ret, InvalidInputError
	}

	valueOfListA := reflect.ValueOf(lista)
	valueOfListB := reflect.ValueOf(listb)

	kindOfInputA := typeOfListA.Elem().Kind()
	kindOfInputB := typeOfListB.Elem().Kind()
	if kindOfInputA != kindOfInputB {
		return ret, InvalidInputError
	}

	cap := valueOfListA.Len()
	if cap > valueOfListB.Len() {
		cap = valueOfListB.Len()
	}

	result := reflect.MakeSlice(typeOfListA, 0, 0)

	for i := 0; i < valueOfListA.Len(); i++ {
		elem := valueOfListA.Index(i)
		for j := 0; j < valueOfListB.Len(); j++ {
			if elem.Interface() == valueOfListB.Index(j).Interface() {
				result = reflect.Append(result, elem)
			}
		}
	}
	return result.Interface(), nil
}
