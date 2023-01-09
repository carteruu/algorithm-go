package offer_1

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cbtInserter := Constructor(tt.args.root)
			if got := cbtInserter.Get_root(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get_root() = %v, want %v", got, tt.want)
			}
		})
	}
}
