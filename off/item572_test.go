package off

import "testing"

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				root:    BuildByLevelInterface([]interface{}{3, 4, 5, 1, 2}),
				subRoot: BuildByLevelInterface([]interface{}{4, 1, 2}),
			},
			want: true,
		}, {
			name: "1",
			args: args{
				root:    BuildByLevelInterface([]interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0}),
				subRoot: BuildByLevelInterface([]interface{}{4, 1, 2}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
