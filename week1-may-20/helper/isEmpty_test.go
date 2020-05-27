package helper

import (
	"testing"
)

func Test_isEmpty(t *testing.T) {
	type args struct {
		input interface{}
	}
// 	//input = ["", 0, false, nil] => TRUE
// 	//input = ["a", 1, true] => false
	tests := []struct {
		name string
		args args
		want bool
	}{{
		name: "Test with 1",
		args: args{"a"},
		want: false,
	},{
		name: "Test with 1",
		args: args{1},
		want: false,
	},{
		name: "Test with 1",
		args: args{true},
		want: false,
	},{
		name: "Test with 0",
		args: args{0},
		want: true,
	},{
		name: "Test with 0",
		args: args{""},
		want: true,
	},{
		name: "Test with 0",
		args: args{false},
		want: true,
	},{
		name: "Test with 0",
		args: args{nil},
		want: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEmpty(tt.args.input); got != tt.want {
				t.Errorf("isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_isEmpty_2(t *testing.T) {
	
// 	result := isEmpty(1)
// 	if result != false {
// 		t.Error("Failed. This should be false.")
// 	}
// 	result = isEmpty(0)
// 	if result != true {
// 		t.Error("Failed. This should be true.")
// 	}
// }

// func Test_isEmpty_3(t *testing.T) {
// 	//input = ["", 0, false, nil] => TRUE
// 	//input = ["a", 1, true] => false
// 	result := isEmpty("")
// 	if result != true {
// 		t.Error("Failed. This should be true.")
// 	}
// }