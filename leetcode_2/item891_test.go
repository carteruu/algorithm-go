package leetcode_2

import "testing"

func Test_sumSubseqWidths(t *testing.T) {
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
				nums: []int{2, 1, 3},
			},
			want: 6,
		}, {
			name: "",
			args: args{
				nums: []int{5, 69, 89, 92, 31, 16, 25, 45, 63, 40, 16, 56, 24, 40, 75, 82, 40, 12, 50, 62, 92, 44, 67, 38, 92, 22, 91, 24, 26, 21, 100, 42, 23, 56, 64, 43, 95, 76, 84, 79, 89, 4, 16, 94, 16, 77, 92, 9, 30, 13},
			},
			want: 857876214,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubseqWidths2(tt.args.nums); got != tt.want {
				t.Errorf("sumSubseqWidths() = %v, want %v", got, tt.want)
			}
		})
	}
}
