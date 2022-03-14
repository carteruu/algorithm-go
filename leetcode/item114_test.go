package leetcode

import (
	"algorithm/pkg"
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *pkg.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *pkg.TreeNode
	}{
		// TODO: Add test cases.
		//{
		//	name: "0",
		//	args: args{
		//		root: &pkg.TreeNode{},
		//	},
		//	want: &pkg.TreeNode{},
		//},
		{
			name: "0",
			args: args{
				root: &pkg.TreeNode{1,
					&pkg.TreeNode{2, &pkg.TreeNode{3, nil, nil}, &pkg.TreeNode{4, nil, nil}},
					&pkg.TreeNode{5, nil, &pkg.TreeNode{6, nil, nil}}},
			},
			want: &pkg.TreeNode{1, nil, &pkg.TreeNode{2, nil, &pkg.TreeNode{3, nil, &pkg.TreeNode{4, nil, &pkg.TreeNode{5, nil, &pkg.TreeNode{6, nil, nil}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if flatten(tt.args.root); !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("flatten() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}
