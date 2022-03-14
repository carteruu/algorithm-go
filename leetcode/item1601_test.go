package leetcode

import "testing"

func Test_maximumRequests(t *testing.T) {
	type args struct {
		n        int
		requests [][]int
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
				n:        2,
				requests: [][]int{{1, 1}, {1, 0}, {0, 1}, {0, 0}, {0, 0}, {0, 1}, {0, 1}, {1, 0}, {1, 0}, {1, 1}, {0, 0}, {1, 0}},
			},
			want: 11,
		}, {
			name: "1",
			args: args{
				n:        3,
				requests: [][]int{{0, 0}, {1, 2}, {2, 1}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRequests(tt.args.n, tt.args.requests); got != tt.want {
				t.Errorf("maximumRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
