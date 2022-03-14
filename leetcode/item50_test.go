package leetcode

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
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
				x: 2,
				n: 10,
			},
			want: 1024.0,
		}, {
			name: "1",
			args: args{
				x: 2,
				n: 0,
			},
			want: 1,
		}, {
			name: "1",
			args: args{
				x: 2,
				n: -2,
			},
			want: 0.25,
		}, {
			name: "1",
			args: args{
				x: 0.00001,
				n: 2147483647,
			},
			want: 0,
		}, {
			name: "1",
			args: args{
				x: 2.00000,
				n: -2147483648,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
