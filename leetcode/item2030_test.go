package leetcode

import "testing"

func Test_smallestSubsequence(t *testing.T) {
	type args struct {
		s          string
		k          int
		letter     byte
		repetition int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				s:          "leet",
				k:          3,
				letter:     'e',
				repetition: 2,
			},
			want: "eet",
		}, {
			name: "0",
			args: args{
				s:          "leetcode",
				k:          4,
				letter:     'e',
				repetition: 2,
			},
			want: "eet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubsequence(tt.args.s, tt.args.k, tt.args.letter, tt.args.repetition); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
