package leetcode_2

import "testing"

func Test_arrayStringsAreEqual(t *testing.T) {
	type args struct {
		word1 []string
		word2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				word1: []string{"ab", "c"},
				word2: []string{"a", "bc"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayStringsAreEqual(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("arrayStringsAreEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
