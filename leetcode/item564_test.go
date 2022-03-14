package leetcode

import "testing"

func Test_nearestPalindromic(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{n: "100"},
			want: "99",
		}, {
			name: "1",
			args: args{n: "9"},
			want: "8",
		}, {
			name: "1",
			args: args{n: "99"},
			want: "101",
		}, {
			name: "1",
			args: args{n: "111"},
			want: "101",
		}, {
			name: "1",
			args: args{n: "101"},
			want: "99",
		}, {
			name: "1",
			args: args{n: "1223"},
			want: "1221",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestPalindromic(tt.args.n); got != tt.want {
				t.Errorf("nearestPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
