package helper

import (
	"reflect"
	"testing"
)

func TestLast(t *testing.T) {
	arr := []int{1, 2, 3}
	actual := Last(arr)
	var expected int
	expected = 3
	if actual != expected {
		t.Error("Value should be ", expected)
	}
}

func TestFilter(t *testing.T) {
	// _.filter(users, function(o) { return !o.active; });
	t.Run("Filter with Func", func(t *testing.T) {
		arr := []int{1, 2, 3}
		actual := Filter(arr, func(ele int) bool {
			return ele%2 == 0
		})
		expected := []int{2}
		if !reflect.DeepEqual(expected, actual) {
			t.Error("Value should be ", expected)
		}
	})
	// _.filter(users, 'active');
	t.Run("Filter with Field", func(t *testing.T) {
		type User struct {
			Name   string
			Age    int
			Active bool
		}
		users := []User{
			User{"barney", 36, true},
			User{"fred", 40, false},
		}
		actual := Filter(users, "Active")
		expected := []User{
			User{"barney", 36, true},
		}
		if !reflect.DeepEqual(expected, actual) {
			t.Error("Value should be ", expected)
		}
	})
	// _.filter(users, ['active', false]);
	t.Run("Filter with Field", func(t *testing.T) {
		type User struct {
			Name   string
			Age    int
			Active bool
		}
		users := []User{
			User{"barney", 36, true},
			User{"fred", 40, false},
		}
		actual := Filter(users, []interface{}{"Active", false})
		expected := []User{
			User{"fred", 40, false},
		}
		if !reflect.DeepEqual(expected, actual) {
			t.Error("Value should be ", expected)
		}
	})
}
