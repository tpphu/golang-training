package helper

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMax_1(t *testing.T) {

	list := []int{1, 2, 3, 4}

	actual, err := Max(list)
	if err != nil {
		t.Error("This case should not error")
	}
	expect := 4
	if actual != expect {
		t.Errorf("Max() = %v, want %v", actual, expect)
	}
}

func TestMax_2(t *testing.T) {

	list := []float32{1.0, 2.1, 3.3, 4.4}

	actual, err := Max(list)
	if err != nil {
		t.Error("This case should not error")
	}
	// Se check lai va tra loi default float no luon la 64 hay la tuy thuoc
	// may 64bit hay 32bit
	expect := 4
	kind := reflect.ValueOf(expect).Kind()
	if kind == reflect.Int {
		fmt.Println("The default float is Int")
	} else {
		fmt.Println("The default float is not Int")
	}
	if actual != expect {
		t.Errorf("Max() = %v, want %v", actual, expect)
	}
}

func TestMax_3(t *testing.T) {
	list := ""
	_, err := Max(list)
	if err != InvalidInputError {
		t.Errorf("This test case expected an error: %v", err)
	}
}

func TestMax_4(t *testing.T) {

	list := []interface{}{1.0, 2, "abc", true}

	_, err := Max(list)
	if err == nil {
		t.Error("This case should be error")
	}
	// Se check lai va tra loi default float no luon la 64 hay la tuy thuoc
	// may 64bit hay 32bit
	// expect := 4
	// if actual != expect {
	// 	t.Errorf("Max() = %v, want %v", actual, expect)
	// }
}
