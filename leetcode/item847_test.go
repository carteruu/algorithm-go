package leetcode

import "testing"

func Test_shortestPathLength(t *testing.T) {
	type args struct {
		graph [][]int
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
				graph: [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}},
			},
			want: 4,
		},
		{
			name: "1",
			args: args{
				graph: [][]int{{1, 2, 3}, {0}, {0}, {0}},
			},
			want: 4,
		}, {
			name: "1",
			args: args{
				graph: [][]int{{2, 3, 5, 7}, {2, 3, 7}, {0, 1}, {0, 1}, {7}, {0}, {10}, {9, 10, 0, 1, 4}, {9}, {7, 8}, {7, 6}},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathLength(tt.args.graph); got != tt.want {
				t.Errorf("shortestPathLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
