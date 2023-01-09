package leetcode_2

import "testing"

func Test_maxRepeating(t *testing.T) {
	type args struct {
		sequence string
		word     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				sequence: "ababc",
				word:     "ba",
			},
			want: 1,
		}, {
			name: "",
			args: args{
				sequence: "ababc",
				word:     "ab",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepeating(tt.args.sequence, tt.args.word); got != tt.want {
				t.Errorf("maxRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}
