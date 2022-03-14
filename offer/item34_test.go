package offer

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				root:   BuildByLevelInterface([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}),
				target: 22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
