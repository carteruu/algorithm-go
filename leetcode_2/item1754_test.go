package leetcode_2

import "testing"

func Test_largestMerge(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				word1: "cabaa",
				word2: "bcaaa",
			},
			want: "cbcabaaaaa",
		}, {
			name: "",
			args: args{
				word1: "guguuuuuuuuuuuuuuguguuuuguug",
				word2: "gguggggggguuggguugggggg",
			},
			want: "guguuuuuuuuuuuuuuguguuuuguugguggggggguuggguuggggggg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestMerge(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("largestMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
