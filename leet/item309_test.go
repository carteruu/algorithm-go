package leet

import "testing"

func Test_maxProfit309(t *testing.T) {
	type args struct {
		prices []int
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
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit309(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit309() = %v, want %v", got, tt.want)
			}
		})
	}
}
