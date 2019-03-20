package helper

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMax_1(t *testing.T) {
	v := []int{1, 2, 3}
	expected := 3

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}

func TestMax_2(t *testing.T) {
	v := []int{1, 2, 4}
	expected := 4

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}

func TestMax_3(t *testing.T) {
	v := []int32{1, 2, 4}
	var expected int32 = 4

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}

// 1. Ham nay dang dung
// 2. Minh can phai refactor cho no tong qua
func TestMap2_1(t *testing.T) {
	v := make([]int, 5)
	v[0] = 1
	v[1] = 5
	v[2] = 3
	v[3] = 7
	v[4] = 8

	expected := []int{3, 15, 9, 21, 24}

	actual := Map2(v, func(x int) int { return x * 3 })
	fmt.Println("actual", actual)
	//_ = result1 // declared and not used
	// for i := 0; i < len(v); i++ {
	// 	if result1[i] != v[i]*3 {
	// 		t.Error("khong bang")
	// 	}
	// }

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Actual should be expected")
	}
}

func TestMap2_2(t *testing.T) {
	v := []int{1, 2, 3}

	expected := []int{2, 4, 6}

	actual := Map2(v, func(x int) int { return x * 2 })

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Actual should be expected")
	}
}

func TestMap2_3_Count_String_Len(t *testing.T) {
	v := []string{"a", "bb", "ccc"}

	expected := []int{1, 2, 3}

	actual := Map2(v, func(elem string) int { return len(elem) })

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Actual should be expected")
	}
}
