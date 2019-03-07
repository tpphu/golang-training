package helper

import "testing"

func TestLast(t *testing.T) {
	arr := []int{1, 2, 3}
	v := Last(arr)
	var expected int
	expected = 3
	if v != expected {
		t.Error("Value should be ", expected)
	}
}
