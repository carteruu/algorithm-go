package leetcode

import (
	"algorithm/pkg"
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *pkg.TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		// TODO: Add test cases.
		{
			name:    "0",
			args:    args{root: &pkg.TreeNode{1, nil, nil}},
			wantAns: []int{1},
		}, {
			name:    "0",
			args:    args{root: &pkg.TreeNode{1, &pkg.TreeNode{2, nil, &pkg.TreeNode{3, nil, nil}}, &pkg.TreeNode{4, nil, nil}}},
			wantAns: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := preorderTraversal(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("preorderTraversal() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
