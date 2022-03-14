package leetcode

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		}, {
			name: "2",
			args: args{
				s: "aa",
				p: "aa",
			},
			want: true,
		}, {
			name: "3",
			args: args{
				s: "aa",
				p: "a?",
			},
			want: true,
		}, {
			name: "4",
			args: args{
				s: "",
				p: "",
			},
			want: true,
		}, {
			name: "5",
			args: args{
				s: "aa",
				p: "*",
			},
			want: true,
		}, {
			name: "6",
			args: args{
				s: "cb",
				p: "?a",
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				s: "adceb",
				p: "*a*b",
			},
			want: true,
		},
		{
			name: "8",
			args: args{
				s: "acdcb",
				p: "a*c?b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
