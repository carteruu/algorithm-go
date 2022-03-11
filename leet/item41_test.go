package leet

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				nums: []int{11, 1, 6, 11, 5, 5, -6, 9, -3, 9, 5, 4, 2, -8, 16, -6, -4, 2, 3},
			},
			want: 7,
		},
		{
			name: "0",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			want: 2,
		}, {
			name: "0",
			args: args{
				nums: []int{4, 3, 2, 1},
			},
			want: 5,
		}, {
			name: "0",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 5,
		}, {
			name: "0",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: 3,
		}, {
			name: "0",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
