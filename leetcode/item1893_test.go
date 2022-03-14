package leetcode

import (
	"testing"
)

func Test_isCovered(t *testing.T) {
	type args struct {
		ranges [][]int
		left   int
		right  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				ranges: [][]int{
					{1, 10},
					{10, 20},
				},
				left:  21,
				right: 21,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCovered(tt.args.ranges, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("isCovered() = %v, want %v", got, tt.want)
			}
		})
	}
}
