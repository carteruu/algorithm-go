package leet

import (
	"math"
	"testing"
)

func Test_nextGreaterElement_1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{n: 230241},
			want: 230412,
		}, {
			name: "0",
			args: args{n: 2302431},
			want: 2303124,
		}, {
			name: "0",
			args: args{n: 12},
			want: 21,
		}, {
			name: "0",
			args: args{n: 21},
			want: -1,
		}, {
			name: "0",
			args: args{n: math.MaxInt32},
			want: -1,
		}, {
			name: "0",
			args: args{n: 2147483634},
			want: 2147483643,
		}, {
			name: "0",
			args: args{n: math.MaxInt32 - 1},
			want: -1,
		}, {
			name: "0",
			args: args{n: 1},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement_1(tt.args.n); got != tt.want {
				t.Errorf("nextGreaterElement_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
