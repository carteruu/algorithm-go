package leetcode

import (
	"algorithm/pkg"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
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
			name: "1",
			args: args{
				root: &pkg.TreeNode{
					Val: 1,
					Left: &pkg.TreeNode{
						Val: 2,
						Left: &pkg.TreeNode{
							Val: 4,
						},
						Right: &pkg.TreeNode{
							Val: 5,
						},
					},
					Right: &pkg.TreeNode{
						Val: 3,
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
