package leetcode

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
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
			args: args{nums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8}},
			want: 2,
		}, {
			name: "1",
			args: args{nums: []int{1, 1, 2, 2, 4, 4, 5, 5, 9}},
			want: 9,
		}, {
			name: "1",
			args: args{nums: []int{1, 1, 2}},
			want: 2,
		}, {
			name: "1",
			args: args{nums: []int{1, 2, 2}},
			want: 1,
		}, {
			name: "1",
			args: args{nums: []int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
