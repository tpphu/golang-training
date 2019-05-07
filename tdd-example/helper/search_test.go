package helper

import "testing"

func TestContains(t *testing.T) {
	t.Run("with string | should not contain", func(t *testing.T) {
		s := "hello world"
		actual := Contains(s, "hallo")
		if actual != false {
			t.Fail()
		}
	})
	t.Run("with string | should contain", func(t *testing.T) {
		s := "hello world"
		actual := Contains(s, "hello")
		if actual != true {
			t.Fail()
		}
	})
}
