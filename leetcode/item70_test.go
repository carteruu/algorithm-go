package leetcode

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
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
				n: 1,
			},
			want: 1,
		}, {
			name: "0",
			args: args{
				n: 2,
			},
			want: 2,
		}, {
			name: "0",
			args: args{
				n: 7,
			},
			want: 1,
		}, {
			name: "0",
			args: args{
				n: 11,
			},
			want: 1,
		}, {
			name: "0",
			args: args{
				n: 100,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
