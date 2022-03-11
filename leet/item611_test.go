package leet

import "testing"

func Test_triangleNumber(t *testing.T) {
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
				nums: []int{2, 2, 3, 4},
			},
			want: 3,
		},
		{
			name: "1",
			args: args{
				nums: []int{4, 2, 3, 4},
			},
			want: 4,
		}, {
			name: "1",
			args: args{
				nums: []int{1, 1, 3, 4},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleNumber(tt.args.nums); got != tt.want {
				t.Errorf("triangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
