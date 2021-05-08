package helper

import (
	"testing"
)

// func TestFindMax(t *testing.T) {
// 	type args struct {
// 		arr []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{{
// 		name: "happy case",
// 		args: args{
// 			arr: []int{1, 2, 3, 4},
// 		},
// 		want: 4,
// 	}, {
// 		name: "unhappy case",
// 		args: args{
// 			arr: []int{},
// 		},
// 		want: -1,
// 	}}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := FindMax(tt.args.arr); got != tt.want {
// 				t.Errorf("FindMax() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestFindMax(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{{
		name: "happy case",
		args: args{
			arr: []int{1, 2, 3, 4},
		},
		want:    4,
		wantErr: false,
	}, {
		name: "unhappy case",
		args: args{
			arr: []int{},
		},
		want:    0,
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMax(tt.args.arr...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMax_HappyCase(t *testing.T) {
	// input := []int{1, 2, 3, 4, 5}
	actual, err := FindMax(1, 2, 3, 4, 5, 6, 7)
	expect := 5
	if actual != expect {
		t.Errorf("Expect: %v, Actual: %v", expect, actual)
	}
	if err != nil {
		t.Error("Error should be nil")
	}
}
