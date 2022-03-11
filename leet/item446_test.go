package leet

import "testing"

func Test_numberOfArithmeticSlices111(t *testing.T) {
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
				[]int{2, 4, 6, 8, 10},
			},
			want: 7,
		}, {
			name: "1",
			args: args{
				[]int{7, 7, 7, 7, 7},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices111(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices111() = %v, want %v", got, tt.want)
			}
		})
	}
}
