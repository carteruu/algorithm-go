package leetcode

import "testing"

func Test_canCross(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				stones: []int{0, 1, 3, 5, 6, 8, 12, 17},
			},
			want: true,
		}, {
			name: "0",
			args: args{
				stones: []int{0, 1, 2, 3, 4, 8, 9, 11},
			},
			want: false,
		}, {
			name: "0",
			args: args{
				stones: []int{0, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCross(tt.args.stones); got != tt.want {
				t.Errorf("canCross() = %v, want %v", got, tt.want)
			}
		})
	}
}
