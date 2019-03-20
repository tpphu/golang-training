package helper

import (
	"reflect"
	"testing"
)

func TestMapInt(t *testing.T) {
	v := []int{1, 2, 3}
	expected := []int{2, 4, 6}
	actual := Map(v, func(elem int) int {
		return elem * 2
	})
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Value should be expected!")
	}
}

func TestMapString(t *testing.T) {
	v := []string{"a", "b", "c"}
	expected := []string{"aa", "bb", "cc"}
	actual := Map(v, func(elem string) string {
		return elem + elem
	})
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Value should be expected!")
	}
}

func TestMapString_2(t *testing.T) {
	v := []string{"a", "bb", "ccc"}
	expected := []int{1, 2, 3}
	actual := Map(v, func(elem string) int {
		return len(elem)
	})
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Value should be expected!")
	}
}
