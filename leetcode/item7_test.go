package leetcode

import (
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				x: -2147483412,
			},
			want: -2143847412,
		}, {
			name: "0",
			args: args{
				x: 1563847412,
			},
			want: 0,
		}, {
			name: "0",
			args: args{
				x: 321,
			},
			want: 123,
		}, {
			name: "0",
			args: args{
				x: 3210,
			},
			want: 123,
		}, {
			name: "0",
			args: args{
				x: -3210,
			},
			want: -123,
		}, {
			name: "0",
			args: args{
				x: math.MaxInt32,
			},
			want: 0,
		}, {
			name: "0",
			args: args{
				x: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
