package leetcode

import (
	"fmt"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
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
				dividend: 10,
				divisor:  3,
			},
			want: 3,
		}, {
			name: "1",
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			want: -2,
		}, {
			name: "1",
			args: args{
				dividend: 0,
				divisor:  3,
			},
			want: 0,
		},
		{
			name: "1",
			args: args{
				dividend: -1,
				divisor:  1,
			},
			want: -1,
		}, {
			name: "1",
			args: args{
				dividend: -2147483648,
				divisor:  -1,
			},
			want: 2147483647,
		}, {
			name: "1",
			args: args{
				dividend: -2147483648,
				divisor:  1,
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("%d - %d", 2, 2<<31)
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
