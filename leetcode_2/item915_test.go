package leetcode_2

import "testing"

func Test_partitionDisjoint1(t *testing.T) {
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
				nums: []int{5, 0, 3, 8, 6},
			},
			want: 3,
		}, {
			name: "",
			args: args{
				nums: []int{1, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionDisjoint1(tt.args.nums); got != tt.want {
				t.Errorf("partitionDisjoint1() = %v, want %v", got, tt.want)
			}
		})
	}
}
