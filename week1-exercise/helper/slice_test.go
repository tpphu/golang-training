package helper

import (
	"reflect"
	"testing"
)

func TestLast(t *testing.T) {
	arr := []int{1, 2, 3}
	v := Last(arr)
	var expected int
	expected = 3
	if v != expected {
		t.Error("Value should be ", expected)
	}
}

func Test_Filter_ByString(t *testing.T) {
	arr := []int{1, 2, 3}
	actual := Filter(arr, func(ele int) bool {
		return ele%2 == 0
	})
	expected := []int{2}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Value should be ", expected)
	}
}
