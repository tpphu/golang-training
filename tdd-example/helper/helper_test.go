package helper

import "testing"

func BenchmarkIsEmpty(b *testing.B) {
	type User struct {
		Name   string
		Age    int
		Active bool
	}
	user := &User{Name: "Phu"}
	for i := 0; i < b.N; i++ {
		IsEmpty(user)
	}
}

func TestIsEmpty(t *testing.T) {
	t.Run("with is empty | simple data case", func(t *testing.T) {
		table := []interface{}{false, "", 0, nil}
		for _, v := range table {
			actual := IsEmpty(v)
			if actual != true {
				t.Fail()
			}
		}
	})
	t.Run("with is not empty | simple data case", func(t *testing.T) {
		table := []interface{}{true, "a", 1}
		for _, v := range table {
			actual := IsEmpty(v)
			if actual != false {
				t.Error("IsEmpty of \"", v, "\" is failed")
			}
		}
	})
	t.Run("with is not empty | slice or array case", func(t *testing.T) {
		value := []int{}
		actual := IsEmpty(value)
		if actual != true {
			t.Error("Result should be true")
		}
	})
	t.Run("with is not empty | struct case", func(t *testing.T) {
		type User struct {
			Name   string
			Age    int
			Active bool
		}
		user := &User{Name: "Phu"}
		actual := IsEmpty(user)
		if actual != false {
			t.Error("Result should be false")
		}
	})
	t.Run("with is empty | struct case", func(t *testing.T) {
		type User struct {
			Name   string
			Age    int
			Active bool
		}
		user := &User{}
		actual := IsEmpty(user)
		if actual != true {
			t.Error("Result should be false")
		}
	})
}
