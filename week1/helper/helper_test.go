package helper

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "Happy case",
		args: args{arr: []int{1, 2, 3, 4}},
		want: 10,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.arr...); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Sum_2(t *testing.T) {
	actual := Sum(1, 2, 3, 4, 5)
	expect := 15
	if actual != expect {
		t.Errorf("Actual %v is not equal to expect %v", actual, expect)
	}
}

func Test_Add(t *testing.T) {
	actual := Add(1, 2)
	expect := 3
	if actual != expect {
		t.Errorf("Actual %v is not equal to expect %v", actual, expect)
	}
}
