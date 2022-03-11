package leet

import (
	"algorithm/pkg"
	"testing"
)

func Test_getMinimumDifference(t *testing.T) {
	type args struct {
		root *pkg.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				&pkg.TreeNode{1,
					nil,
					&pkg.TreeNode{3,
						&pkg.TreeNode{2,
							nil,
							nil,
						},
						nil,
					},
				},
			},
			1,
		},
		//[1,0,48,null,null,12,49]
		{
			"2",
			args{
				root: &pkg.TreeNode{Val: 1,
					Left: &pkg.TreeNode{Val: 0},
					Right: &pkg.TreeNode{
						Val: 48,
						Left: &pkg.TreeNode{
							Val: 12,
						},
						Right: &pkg.TreeNode{
							Val: 49,
						},
					},
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
