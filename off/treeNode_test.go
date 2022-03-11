package off

import (
	"reflect"
	"testing"
)

func TestBuildByLevel(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				arr: []int{4, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 8,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildByLevel(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildByLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildByLevelInterface(t *testing.T) {
	type args struct {
		arr []interface{}
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				arr: []interface{}{4, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 8,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildByLevelInterface(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildByLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
