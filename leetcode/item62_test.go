package leetcode

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		}, {
			name: "0",
			args: args{
				m: 3,
				n: 1,
			},
			want: 1,
		}, {
			name: "0",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		}, {
			name: "0",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		}, {
			name: "0",
			args: args{
				m: 1,
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
