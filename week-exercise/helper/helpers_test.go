package helper

import "testing"

func TestIsEmpty(t *testing.T) {
	table := []interface{}{false, "", 0, nil}
	for _, v := range table {
		expected := IsEmpty(v)
		if expected != true {
			t.Error("IsEmpty of \"", v, "\" is failed.")
		}
	}
}
