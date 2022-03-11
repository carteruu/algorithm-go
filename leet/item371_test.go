package leet

import (
	"fmt"
	"math"
	"testing"
)

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
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
				a: 3,
				b: 6,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				fmt.Println(math.MaxInt32)
				fmt.Println(math.MinInt32)
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
