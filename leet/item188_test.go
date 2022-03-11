package leet

import "testing"

func Test_maxProfit188(t *testing.T) {
	type args struct {
		k      int
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
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		}, {
			name: "0",
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit188(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit188() = %v, want %v", got, tt.want)
			}
		})
	}
}
