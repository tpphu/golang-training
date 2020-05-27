package helper

import (
	"testing"
)

func Test_max(t *testing.T) {
	type args struct {
		list []int
	}
	type want struct {
		result int
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "Test with normal case",
		args: args{list: []int{1, 3, 4, 5}},
		want: want{
			result: 5,
			err: nil,
		},
	},{
		name: "Test with not normal case",
		args: args{list: []int{}},
		want: want{
			result: 0, 
			err: ErrLenOfListIsEmpty,
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := max(tt.args.list);
			if got != tt.want.result {
				t.Errorf("max() = %v, want %v", got, tt.want.result)
			}
			// if tt.want.err == nil && err == nil {
			// 	return
			// }
			// if tt.want.err != nil && err == nil {
			// 	t.Error("2 error khac nhau => loi")
			// }
			// if tt.want.err == nil && err != nil {
			// 	t.Error("2 error khac nhau => loi")
			// }
			if err != tt.want.err {
				t.Errorf("max() = %v, want %v", err, tt.want.err)
			}
		})
	}
}
