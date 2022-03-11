package leet

import "testing"

func Test_countMaxOrSubsets(t *testing.T) {
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
			args: args{nums: []int{}},
			want: 0,
		}, {
			name: "0",
			args: args{nums: []int{2, 2}},
			want: 3,
		}, {
			name: "0",
			args: args{nums: []int{2, 2, 2}},
			want: 7,
		}, {
			name: "0",
			args: args{nums: []int{1, 2}},
			want: 1,
		}, {
			name: "0",
			args: args{nums: []int{1, 3}},
			want: 2,
		}, {
			name: "0",
			args: args{nums: []int{3, 2, 1, 5}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMaxOrSubsets(tt.args.nums); got != tt.want {
				t.Errorf("countMaxOrSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
