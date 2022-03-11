package leet

import (
	"algorithm/pkg"
	"testing"
)

func Test_widthOfBinaryTree(t *testing.T) {
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
				root: pkg.BuildByLevelInterface([]interface{}{1, 3, 2, 5, nil, nil, 9, 6, nil, nil, 7}),
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := widthOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
