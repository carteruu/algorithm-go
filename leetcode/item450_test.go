package leetcode

import (
	"algorithm/pkg"
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *pkg.TreeNode
		key  int
	}
	root1 := &pkg.TreeNode{
		Val: 5,
		Left: &pkg.TreeNode{
			Val: 3,
			Left: &pkg.TreeNode{
				Val: 2,
			},
			Right: &pkg.TreeNode{
				Val: 4,
			},
		},
		Right: &pkg.TreeNode{
			Val: 6,
			Right: &pkg.TreeNode{
				Val: 7,
			},
		},
	}
	tests := []struct {
		name string
		args args
		want *pkg.TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				root: root1,
				key:  3,
			},
			want: root1,
		}, {
			name: "1",
			args: args{
				root: root1,
				key:  5,
			},
			want: root1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode1(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
