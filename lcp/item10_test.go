package lcp

import "testing"

func Test_minimalExecTime(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				root: &TreeNode{
					Val: 47,
					Left: &TreeNode{
						Val: 74,
					},
					Right: &TreeNode{
						Val: 31,
					},
				},
			},
			want: 121,
		}, {
			name: "2",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
							Left: &TreeNode{
								Val: 5,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
						Right: &TreeNode{
							Val: 0,
							Left: &TreeNode{
								Val: 5,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
							Left: &TreeNode{
								Val: 5,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
						Right: &TreeNode{
							Val: 0,
						},
					},
				},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimalExecTime(tt.args.root); got != tt.want {
				t.Errorf("minimalExecTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
