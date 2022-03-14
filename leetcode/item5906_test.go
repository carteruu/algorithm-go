package leetcode

import "testing"

func Test_countValidWords(t *testing.T) {
	type args struct {
		sentence string
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
				sentence: "a-b-c",
			},
			want: 0,
		}, {
			name: "0",
			args: args{
				sentence: "!",
			},
			want: 1,
		}, {
			name: "0",
			args: args{
				sentence: "!this  1-s b8d!",
			},
			want: 0,
		}, {
			name: "0",
			args: args{
				sentence: "alice and  bob are playing stone-game10",
			},
			want: 5,
		}, {
			name: "0",
			args: args{
				sentence: "he bought 2 pencils, 3 erasers, and 1  pencil-sharpener.",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValidWords(tt.args.sentence); got != tt.want {
				t.Errorf("countValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
