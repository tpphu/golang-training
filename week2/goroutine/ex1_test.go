package goroutine

import "testing"

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := Sum(arr)
	if sum != 45 {
		t.Errorf("This want: %d but got: %d", 45, sum)
	}
}
