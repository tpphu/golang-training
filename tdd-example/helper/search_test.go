package helper

import (
	"testing"
)

func TestContains(t *testing.T) {
	// 	_.contains([1, 2, 3], 3);
	// => true
	t.Run("with array | test case is exist", func(t *testing.T) {
		arr := []int{1, 2, 3}
		actual := Contains(arr, 3)
		expected := true
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("with array | test case is not exist", func(t *testing.T) {
		arr := []int{1, 2, 3}
		actual := Contains(arr, 4)
		expected := false
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("with string | should not contain", func(t *testing.T) {
		s := "hello world"
		actual := Contains(s, "hallo")
		expected := false
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("with string | should contain", func(t *testing.T) { //1 What's a case
		s := "hello world"
		actual := Contains(s, "hello") //2. Actually
		expected := true               // 3 Expectation
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("with not same list and value", func(t *testing.T) { //1 What's a case
		arr := []int{1, 2, 3}
		actual := Contains(arr, "hello") //2. Actually
		expected := false                // 3 Expectation
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("with slice in a slice", func(t *testing.T) { //1 What's a case
		arr := []int{1, 2, 3}
		actual := Contains(arr, []int{1, 4}) //2. Actually
		expected := false                    // 3 Expectation
		if actual != expected {
			t.Fail()
		}
	})
	// t.Run("with slice in a slice", func(t *testing.T) { //1 What's a case
	// 	arr := []int{1, 2, 3}
	// 	actual := Contains(arr, []int{1, 2}) //2. Actually
	// 	expected := true                     // 3 Expectation
	// 	if actual != expected {
	// 		t.Fail()
	// 	}
	// })
}
