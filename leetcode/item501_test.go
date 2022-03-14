package leetcode

import (
	"algorithm/pkg"
	"reflect"
	"testing"
)

func Test_findMode(t *testing.T) {
	type args struct {
		root *pkg.TreeNode
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{&pkg.TreeNode{
				Val: 1,
				Right: &pkg.TreeNode{
					Val: 2,
					Left: &pkg.TreeNode{
						Val: 2,
					},
				},
			}},
			want: []int{2},
		}, {
			name: "2",
			args: args{&pkg.TreeNode{
				Val: 1,
				Right: &pkg.TreeNode{
					Val: 2,
				},
			}},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
