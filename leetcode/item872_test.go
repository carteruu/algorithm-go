package leetcode

import (
	"algorithm/pkg"
	"testing"
)

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *pkg.TreeNode
		root2 *pkg.TreeNode
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
				root1: &pkg.TreeNode{
					Val: 1,
					Left: &pkg.TreeNode{
						Val: 2,
					},
					Right: &pkg.TreeNode{
						Val: 3,
					},
				},
				root2: &pkg.TreeNode{
					Val: 1,
					Left: &pkg.TreeNode{
						Val: 2,
					},
					Right: &pkg.TreeNode{
						Val: 3,
					},
				},
			},
			want: true,
		}, {
			name: "2",
			args: args{
				root1: &pkg.TreeNode{
					Val: 1,
					Left: &pkg.TreeNode{
						Val: 2,
					},
					Right: &pkg.TreeNode{
						Val: 3,
					},
				},
				root2: &pkg.TreeNode{
					Val: 1,
					Left: &pkg.TreeNode{
						Val: 3,
					},
					Right: &pkg.TreeNode{
						Val: 2,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
