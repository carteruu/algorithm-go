package leetcode

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
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
			name: "1",
			args: args{
				[]int{1, 2, 3, 4, 5, 6},
			},
			want: 10,
		}, {
			name: "1",
			args: args{
				[]int{1, 2, 3, 4},
			},
			want: 3,
		}, {
			name: "2",
			args: args{
				[]int{1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
