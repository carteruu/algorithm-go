package leet

import (
	"algorithm/pkg"
	"reflect"
	"testing"
)

func Test_upsideDownBinaryTree(t *testing.T) {
	type args struct {
		root *pkg.TreeNode
	}
	root1 := &pkg.TreeNode{
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
				root: root1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upsideDownBinaryTree(tt.args.root); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("upsideDownBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
