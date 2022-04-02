package leetcode_1

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{x: 100},
			want: 10,
		}, {
			name: "1",
			args: args{x: 10},
			want: 3,
		}, {
			name: "2",
			args: args{x: 2},
			want: 1,
		}, {
			name: "3",
			args: args{x: 1},
			want: 1,
		}, {
			name: "4",
			args: args{x: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
