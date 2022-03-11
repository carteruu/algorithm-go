package leet

import "testing"

func Test_areNumbersAscending(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				s: "1 box has 3 blue 4 red 6 green and 12 yellow marbles",
			},
			want: true,
		}, {
			name: "0",
			args: args{
				s: "hello world 5 x 5",
			},
			want: false,
		}, {
			name: "0",
			args: args{
				s: "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s",
			},
			want: false,
		}, {
			name: "0",
			args: args{
				s: "4 5 11 26",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areNumbersAscending(tt.args.s); got != tt.want {
				t.Errorf("areNumbersAscending() = %v, want %v", got, tt.want)
			}
		})
	}
}
