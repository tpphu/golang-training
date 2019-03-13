package helper

import "testing"

// func TestIsEmpty(t *testing.T) {
// 	table := []interface{}{false, "", 0, nil}
// 	for _, v := range table {
// 		expected := IsEmpty(v)
// 		if expected != true {
// 			t.Error("IsEmpty of \"", v, "\" is failed.")
// 		}
// 	}
// }

func TestIsEmpty(t *testing.T) {
	var v int32 = 0
	expected := false
	result := IsEmpty(v)
	if result != expected {
		t.Error("Result should be false")
	}
}

func TestContainsInt1(t *testing.T) {
	v := []int{1, 2, 3}
	result := ContainsInt(v, 4)
	expected := false
	if result != expected {
		t.Error("Result should be false")
	}
}

func TestContainsInt2(t *testing.T) {
	v := []int{1, 2, 3}
	result := ContainsInt(v, 3)
	expected := true
	if result != expected {
		t.Error("Result should be true")
	}
}

func TestContainsString1(t *testing.T) {
	v := []string{"Phu", "Tien", "KoalaBaby"}
	result := ContainsString(v, "Tien")
	expected := true
	if result != expected {
		t.Error("Result should be true")
	}
}
