package leet

import "testing"

func Test_findUnsortedSubarray(t *testing.T) {
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
				nums: []int{2, 6, 4, 8, 10, 9, 15},
			},
			want: 5,
		}, {
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 0,
		}, {
			name: "1",
			args: args{
				nums: []int{1, 3, 2, 3, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.args.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
