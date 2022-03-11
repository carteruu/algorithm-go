package leet

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
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
				[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		}, {
			name: "0",
			args: args{
				[]int{-2},
			},
			want: -2,
		}, {
			name: "0",
			args: args{
				[]int{-1, -2},
			},
			want: -1,
		}, {
			name: "0",
			args: args{
				[]int{-1, 3, -2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
