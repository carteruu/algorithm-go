package leet

import "testing"

func Test_numberOfGoodSubsets(t *testing.T) {
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
			args: args{nums: []int{18, 28, 2, 17, 29, 30, 15, 9, 12}},
			want: 19,
		}, {
			name: "1",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 6,
		}, {
			name: "1",
			args: args{nums: []int{12, 3}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfGoodSubsets(tt.args.nums); got != tt.want {
				t.Errorf("numberOfGoodSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
