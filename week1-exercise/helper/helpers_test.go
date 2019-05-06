package helper

import "testing"

func TestIsEmptySimple(t *testing.T) {
	table := []interface{}{false, "", 0, nil}
	for _, v := range table {
		actual := IsEmpty(v)
		if actual == false {
			t.Error("IsEmpty of \"", v, "\" is failed.")
		}
	}
}

func TestIsNotEmptySimple(t *testing.T) {
	table := []interface{}{true, "t", 1}
	for _, v := range table {
		actual := IsEmpty(v)
		if actual == true {
			t.Error("IsEmpty of \"", v, "\" is failed.")
		}
	}
}

func TestIsEmptySlice(t *testing.T) {
	value := []int{}
	actual := IsEmpty(value)
	if actual == false {
		t.Error("Result should be true")
	}
}

func TestIsNotEmptySlice(t *testing.T) {
	value := []int{1}
	actual := IsEmpty(value)
	if actual == true {
		t.Error("Result should be false")
	}
}

func TestIsEmptyStruct(t *testing.T) {
	type User struct {
		Name   string
		Age    int
		Active bool
	}
	user := &User{}
	actual := IsEmpty(user)
	if actual == false {
		t.Error("Result should be true")
	}
}

func TestIsNotEmptyStruct(t *testing.T) {
	type User struct {
		Name   string
		Age    int
		Active bool
	}
	user := &User{Name: "Phu"}
	actual := IsEmpty(user)
	if actual == true {
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
