package helper

import "testing"

func TestIsEmpty(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case: input value is nil",
			args: args{nil},
			want: true,
		},
		{
			name: "Case: input value is nil",
			args: args{0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.in); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty_1(t *testing.T) {
	var input int = 1
	actual := IsEmpty(input)
	expect := false
	if actual != expect {
		t.Errorf("IsEmpty() = %v, want %v", actual, expect)
	}
}
