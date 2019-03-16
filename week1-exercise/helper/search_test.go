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
