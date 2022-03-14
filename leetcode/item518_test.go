package leetcode

import "testing"

func Test_change(t *testing.T) {
	type args struct {
		amount int
		coins  []int
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
				amount: 5,
				coins:  []int{1, 2, 5},
			},
			want: 4,
		}, {
			name: "0",
			args: args{
				amount: 5,
				coins:  []int{1, 2, 5, 3},
			},
			want: 6,
		}, {
			name: "0",
			args: args{
				amount: 5000,
				coins:  []int{1, 2, 5, 3},
			},
			want: 696738362,
		}, {
			name: "0",
			args: args{
				amount: 5,
				coins:  []int{2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
		})
	}
}
