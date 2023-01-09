package leetcode_2

import "testing"

func Test_minMoves(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 0, 0, 1, 0, 1},
				k:    2,
			},
			want: 1,
		}, {
			name: "",
			args: args{
				nums: []int{1, 0, 0, 0, 0, 0, 1, 1},
				k:    3,
			},
			want: 5,
		}, {
			name: "",
			args: args{
				nums: []int{0, 1, 1, 0, 0, 1, 0, 0, 0},
				k:    3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
