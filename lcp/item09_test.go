package lcp

import "testing"

func Test_minJump(t *testing.T) {
	type args struct {
		jump []int
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
				jump: []int{2, 5, 1, 1, 1, 1},
			},
			want: 3,
		}, {
			name: "0",
			args: args{
				jump: []int{1, 2, 6, 1, 1, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minJump(tt.args.jump); got != tt.want {
				t.Errorf("minJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
