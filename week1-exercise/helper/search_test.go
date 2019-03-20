package helper

import "testing"

func TestFind1(t *testing.T) {
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

func TestFind_2(t *testing.T) {
	arr := []int{1}
	actual := Find2(arr, 1)
	expected := 1
	if actual != expected {
		t.Error("Expected 1")
	}
}

func TestFind_2_WithLen(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	v := Find2(arr, 4)
	expected := 4
	if v != expected {
		t.Error("Expected 4")
	}
}

func TestFind_2_WithNil(t *testing.T) {
	arr := []int{1}
	v := Find2(arr, 4)
	if v != nil {
		t.Error("Expected nil")
	}
}

func TestFind_2_WithString(t *testing.T) {
	arr := []string{"a"}
	v := Find2(arr, "a")
	expected := "a"
	if v != expected {
		t.Error("Expected a")
	}
}

func TestFind_2_WithFloat(t *testing.T) {
	arr := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	v := Find2(arr, 1.1)
	var expected float64
	expected = 1.1
	if v != expected {
		t.Error("Expected 1.1")
	}
}

func TestFind_2_WithFunc(t *testing.T) {
	arr := []int{1, 2, 3}
	actual := Find2(arr, func(elem int) bool {
		return elem%2 == 0
	})
	expected := 2
	if actual != expected {
		t.Error("Expected 2")
	}
}
