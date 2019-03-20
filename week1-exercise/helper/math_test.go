package helper

import "testing"

func TestMax_1(t *testing.T) {
	v := []int{1, 2, 3}
	expected := 3

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}

func TestMax_2(t *testing.T) {
	v := []int{1, 2, 4}
	expected := 4

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}

func TestMax_3(t *testing.T) {
	v := []int32{1, 2, 4}
	var expected int32 = 4

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}

func TestMap2_1(t *testing.T) {
	v := make([]int, 5)
	v[0] = 1
	v[1] = 5
	v[2] = 3
	v[3] = 7
	v[4] = 8

	result1 := Map2(v, func(x int) int { return x * 3 })
	//_ = result1 // declared and not used
	for i := 0; i < len(v); i++ {
		if result1[i] != v[i]*3 {
			t.Error("khong bang")
		}
	}
}
