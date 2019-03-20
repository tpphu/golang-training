package helper

import "testing"

func TestFind(t *testing.T) {
	in := []int{1, 4, 3}
	expected := 4
	actual := Find(in, func(elem int) bool {
		return elem%2 == 0
	})

	if actual.(int) != expected {
		t.Error("Actual should be expected!")
	}
}

// func TestContains_1(t *testing.T) {
// 	list := []int{1, 2, 3}
// 	value := 3
// 	expected := true
// 	actual := Contains2(list, value)
// 	if actual != expected {
// 		t.Error("actual should be same exptected")
// 	}
// }

// func TestContains_2(t *testing.T) {
// 	list := "Hello world"
// 	value := "world"
// 	expected := true
// 	actual := Contains2(list, value)
// 	if actual != expected {
// 		t.Error("actual should be same exptected")
// 	}
// }
