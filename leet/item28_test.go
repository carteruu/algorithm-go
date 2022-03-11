package leet

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
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
				haystack: "",
				needle:   "",
			},
			want: 0,
		}, {
			name: "0",
			args: args{
				haystack: "",
				needle:   "aaa",
			},
			want: -1,
		}, {
			name: "0",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		}, {
			name: "0",
			args: args{
				haystack: "aaaaa",
				needle:   "aab",
			},
			want: -1,
		}, {
			name: "0",
			args: args{
				haystack: "baa",
				needle:   "aa",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
