package leet

import (
	"algorithm/pkg"
	"testing"
)

func Test_countUnivalSubtrees(t *testing.T) {
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
					Val: 5,
					Left: &pkg.TreeNode{
						Val: 1,
						Left: &pkg.TreeNode{
							Val: 5,
						},
						Right: &pkg.TreeNode{
							Val: 5,
						},
					},
					Right: &pkg.TreeNode{
						Val: 5,
						Right: &pkg.TreeNode{
							Val: 5,
						},
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUnivalSubtrees(tt.args.root); got != tt.want {
				t.Errorf("countUnivalSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
