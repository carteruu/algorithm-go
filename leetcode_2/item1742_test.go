package leetcode_2

import "testing"

func Test_countBalls(t *testing.T) {
	type args struct {
		lowLimit  int
		highLimit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				lowLimit:  1,
				highLimit: 10,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBalls(tt.args.lowLimit, tt.args.highLimit); got != tt.want {
				t.Errorf("countBalls() = %v, want %v", got, tt.want)
			}
		})
	}
}
