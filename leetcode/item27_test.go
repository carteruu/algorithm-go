package leetcode

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
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
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		}, {
			name: "0",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		}, {
			name: "0",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
				val:  2,
			},
			want: 0,
		}, {
			name: "0",
			args: args{
				nums: []int{1, 1, 1},
				val:  3,
			},
			want: 3,
		}, {
			name: "0",
			args: args{
				nums: []int{},
				val:  2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
