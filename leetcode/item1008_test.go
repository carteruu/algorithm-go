package leetcode

import (
	"algorithm/pkg"
	"reflect"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	type args struct {
		preorder []int
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
				preorder: []int{8, 5, 1, 7, 10, 12},
			},
			want: pkg.BuildByLevelInterface([]interface{}{8, 5, 10, 1, 7, nil, 12}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstFromPreorder(tt.args.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
