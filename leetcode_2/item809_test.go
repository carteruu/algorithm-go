package leetcode_2

import "testing"

func Test_expressiveWords(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "heeellooo",
			args: args{
				s:     "heeellooo",
				words: []string{"hello", "hi", "helo"},
			},
			want: 1,
		}, {
			name: "aaa",
			args: args{
				s:     "aaa",
				words: []string{"aaaa"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expressiveWords(tt.args.s, tt.args.words); got != tt.want {
				t.Errorf("expressiveWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
