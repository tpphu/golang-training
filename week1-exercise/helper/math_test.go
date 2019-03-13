package helper

import "testing"

func TestMax1(t *testing.T) {
	v := []int{1, 2, 3}
	expected := 3

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}

func TestMax2(t *testing.T) {
	v := []int{1, 2, 4}
	expected := 4

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}
