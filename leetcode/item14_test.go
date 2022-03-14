package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
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
				strs: []string{"a", ""},
			},
			want: "",
		}, {
			name: "0",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		}, {
			name: "0",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		}, {
			name: "0",
			args: args{
				strs: []string{"dog"},
			},
			want: "dog",
		}, {
			name: "0",
			args: args{
				strs: []string{"dog", "dog"},
			},
			want: "dog",
		}, {
			name: "0",
			args: args{
				strs: []string{"do", "dog"},
			},
			want: "do",
		}, {
			name: "0",
			args: args{
				strs: []string{""},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
