package leetcode_2

import "testing"

func Test_getLucky(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "fleyctuuajsr",
				k: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLucky(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("getLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
