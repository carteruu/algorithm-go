package leetcode

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
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
				s: "(()",
			},
			want: 2,
		}, {
			name: "0",
			args: args{
				s: "(())",
			},
			want: 4,
		}, {
			name: "0",
			args: args{
				s: ")()())",
			},
			want: 4,
		}, {
			name: "0",
			args: args{
				s: "",
			},
			want: 0,
		}, {
			name: "0",
			args: args{
				s: "(()(()(((()(()()))())((((())))))))(()())(()())()",
			},
			want: 48,
		}, {
			name: "0",
			args: args{
				s: ")))())((((())))))))(()())(()())()",
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
