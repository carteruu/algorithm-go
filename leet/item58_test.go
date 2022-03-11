package leet

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
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
			name: "0",
			args: args{
				s: "Hello World",
			},
			want: 5,
		}, {
			name: "0",
			args: args{
				s: "   fly me   to   the moon  ",
			},
			want: 4,
		}, {
			name: "0",
			args: args{
				s: "  luffy is still joyboy",
			},
			want: 6,
		}, {
			name: "0",
			args: args{
				s: " ",
			},
			want: 0,
		}, {
			name: "0",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
