package offer

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
			name: "1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		}, {
			name: "1",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		}, {
			name: "1",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		}, {
			name: "1",
			args: args{
				s: "",
			},
			want: 0,
		}, {
			name: "1",
			args: args{
				s: " ",
			},
			want: 1,
		}, {
			name: "1",
			args: args{
				s: "吴",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
