package leetcode

import (
	"math"
	"testing"
)

func Test_mySqrt(t *testing.T) {
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
				math.MaxInt32,
			},
			want: 1,
		}, {
			name: "0",
			args: args{
				8,
			},
			want: 2,
		}, {
			name: "0",
			args: args{
				4,
			},
			want: 2,
		}, {
			name: "0",
			args: args{
				1,
			},
			want: 1,
		}, {
			name: "0",
			args: args{
				0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
