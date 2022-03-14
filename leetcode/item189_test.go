package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotate11(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		}, {
			name: "0",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			want: []int{3, 99, -1, -100},
		}, {
			name: "0",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    0,
			},
			want: []int{-1, -100, 3, 99},
		}, {
			name: "0",
			args: args{
				nums: []int{},
				k:    1,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if rotate189(tt.args.nums, tt.args.k); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
