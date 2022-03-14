package leetcode

import "testing"

func Test_numSquarefulPerms(t *testing.T) {
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
				nums: []int{1, 17, 8},
			},
			want: 2,
		}, {
			name: "2",
			args: args{
				nums: []int{2, 2, 2},
			},
			want: 1,
		}, {
			name: "3",
			args: args{
				nums: []int{},
			},
			want: 1,
		}, {
			name: "4",
			args: args{
				nums: []int{1, 3},
			},
			want: 2,
		}, {
			name: "5",
			args: args{
				nums: []int{1, 4},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquarefulPerms(tt.args.nums); got != tt.want {
				t.Errorf("numSquarefulPerms() = %v, want %v", got, tt.want)
			}
		})
	}
}
