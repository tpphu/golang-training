package helper

import (
	"testing"
)

func TestMax(t *testing.T) {
	t.Run("Test invalid input", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("The code should panic")
			}
		}()
		Max("abc")
	})
	t.Run("Test empty input", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("The code should panic")
			}
		}()
		Max([]int{})
	})
	t.Run("Test with int", func(t *testing.T) {
		arr := []int{4, 2, 8, 6}
		actual := Max(arr)
		expected := 8
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("Test with float32", func(t *testing.T) {
		arr := []float32{4.5, 2.3, 8.1, 6.2}
		actual := Max(arr)
		var expected float32 = 8.1
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("Test with uint8", func(t *testing.T) {
		arr := []uint8{1, 5, 7, 9}
		actual := Max(arr)
		var expected uint8 = 9
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("Test with uint16", func(t *testing.T) {
		arr := []uint16{1, 5, 7, 9}
		actual := Max(arr)
		var expected uint16 = 9
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("Test with uint32", func(t *testing.T) {
		arr := []uint32{1, 5, 7, 9}
		actual := Max(arr)
		var expected uint32 = 9
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("Test with uint64", func(t *testing.T) {
		arr := []uint64{1, 5, 7, 9}
		actual := Max(arr)
		var expected uint64 = 9
		if actual != expected {
			t.Fail()
		}
	})
}

// Test case cho truong hop panic
// func TestMaxWithEmptySlice(t *testing.T) {
// 	// arr := []int{}
// 	// Can test cai logic recover
// }
