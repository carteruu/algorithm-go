package leetcode

import (
	"reflect"
	"testing"
)

func Test_generatePalindromes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				s: "aabb",
			},
			want: []string{"abba", "baab"},
		}, {
			name: "2",
			args: args{
				s: "abc",
			},
			want: []string{},
		}, {
			name: "3",
			args: args{
				s: "a",
			},
			want: []string{"a"},
		}, {
			name: "4",
			args: args{
				s: "abb",
			},
			want: []string{"bab"},
		}, {
			name: "5",
			args: args{
				s: "aaa",
			},
			want: []string{"aaa"},
		}, {
			name: "6",
			args: args{
				s: "aaaabbbb",
			},
			want: []string{"aabbbbaa", "ababbaba", "abbaabba", "baabbaab", "babaabab", "bbaaaabb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePalindromes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePalindromes() = %v, want %v", got, tt.want)
			}
		})
	}
}
