package helper

import (
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		n []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{{
		name: "normal case 1",
		args: args{
			n: []interface{}{1, 2},
		},
		want: 2,
	}, {
		name: "normal case 2",
		args: args{
			n: []interface{}{1, 2, 10},
		},
		want: 10,
	}, {
		name: "normal case 3",
		args: args{
			n: []interface{}{int64(1), int64(2), int64(10)},
		},
		want: int64(10),
	}, {
		name: "special case 1",
		args: args{
			n: []interface{}{1, int64(2), 10.1},
		},
		want: 10.1,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.n...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
