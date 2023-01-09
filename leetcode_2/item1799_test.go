package leetcode_2

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2},
			},
			want: 1,
		}, {
			name: "",
			args: args{
				nums: []int{3, 4, 6, 8},
			},
			want: 11,
		}, {
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
			},
			want: 14,
		}, {
			name: "",
			args: args{
				nums: []int{327161, 405335, 360021, 441366, 834880, 197174, 270974, 114045, 265771, 943529},
			},
			want: 219,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.nums); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
