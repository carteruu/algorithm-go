package leetcode

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 2, 0, 634, 0, 0, 0, 3, 3},
			},
			want: []int{1, 2, 634, 3, 3, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
