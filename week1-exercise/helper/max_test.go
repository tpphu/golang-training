package helper

import "testing"

func TestMaxInt(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "Test with normal array",
		args: args{[]int{1, 2, 3, 4}},
		want: 4,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.list); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt_1(t *testing.T) {

	list := []int{1, 2, 3, 4}

	actual := MaxInt(list)
	expect := 4
	if actual != expect {
		t.Errorf("MaxInt() = %v, want %v", actual, expect)
	}
}
