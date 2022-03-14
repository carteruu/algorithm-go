package offer

import "testing"

func Test_longestUnivaluePath(t *testing.T) {
	type args struct {
		root *TreeNode
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
				root: BuildByLevelInterface([]interface{}{5, 4, 5, 1, 1, nil, 5}),
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				root: BuildByLevelInterface([]interface{}{1, 4, 5, 4, 4, 5}),
			},
			want: 2,
		}, {
			name: "2",
			args: args{
				root: BuildByLevelInterface([]interface{}{1, nil, 1, 1, 1, 1, 1, 1}),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestUnivaluePath(tt.args.root); got != tt.want {
				t.Errorf("longestUnivaluePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
