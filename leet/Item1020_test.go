package leet

import "testing"

func Test_numEnclaves(t *testing.T) {
	type args struct {
		grid [][]int
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
				grid: [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
			},
			want: 3,
		}, {
			name: "1",
			args: args{
				grid: [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
			},
			want: 0,
		}, {
			name: "1",
			args: args{
				grid: [][]int{
					{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
					{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
					{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
					{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
					{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
					{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
					{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
					{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
					{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 1, 0, 0, 0, 1},
				},
			},
			want: 3,
		}, {
			name: "1",
			args: args{
				grid: [][]int{
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0},
					{1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0},
					{1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1},
					{1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0},
					{1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1},
					{1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1},
					{1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1},
					{0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0},
					{0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1},
					{0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEnclaves(tt.args.grid); got != tt.want {
				t.Errorf("numEnclaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
