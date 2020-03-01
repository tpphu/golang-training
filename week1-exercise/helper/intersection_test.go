package helper

import (
	"reflect"
	"testing"
)

func TestIntersectionInt(t *testing.T) {
	lista := []int{1, 2, 3}
	listb := []int{2, 3, 4}
	listc := []int{3, 4, 5}
	expect := []int{3}
	actual, _ := IntersectionInt(lista, listb, listc)
	if reflect.DeepEqual(actual, expect) != true {
		t.Errorf("IntersectionInt() = %v, want: %v", actual, expect)
	}
}

func TestIntersectionSimple(t *testing.T) {
	lista := []int{1, 2}
	listb := []int{2, 3}
	expect := []int{2}
	actual, _ := IntersectionInt(lista, listb)
	if reflect.DeepEqual(actual, expect) != true {
		t.Errorf("IntersectionInt() = %v, want: %v", actual, expect)
	}
}

func TestIntersection_2(t *testing.T) {
	lista := []int{1, 2, 3, 4, 5}
	listb := []int{2, 3, 4, 5, 6}
	expect := []int{2, 3, 4, 5}
	actual, _ := IntersectionInt(lista, listb)
	if reflect.DeepEqual(actual, expect) != true {
		t.Errorf("IntersectionInt() = %v, want: %v", actual, expect)
	}
}
