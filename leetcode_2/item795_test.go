package leetcode_2

import "testing"

func Test_numSubarrayBoundedMax(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:  []int{2, 1, 4, 3},
				left:  2,
				right: 3,
			},
			want: 3,
		}, {
			name: "",
			args: args{
				nums:  []int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52},
				left:  32,
				right: 69,
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayBoundedMax(tt.args.nums, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("numSubarrayBoundedMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
