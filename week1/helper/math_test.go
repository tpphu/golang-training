package helper

import (
	"fmt"
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

func TestPointer(t *testing.T) {
	// var n int = 10
	// var p *int = &n
	n := 10 // Khai bao va gan gia tri
	p := &n

	fmt.Println("&n", &n, "n", n)
	fmt.Println("&p", &p, "p", p)
	fmt.Println("*p", *p)
	n = 11 // Change n
	fmt.Println("*p", *p)
	m := 999
	p = &m
	fmt.Println("*p", *p)
	*p = 1000
	fmt.Println("m", m) // 1000
	fmt.Println("n", n) // 11
}

// Struct & Function
