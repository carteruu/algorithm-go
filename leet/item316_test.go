package leet

import "testing"

func Test_removeDuplicateLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{
				s: "dfhfdsfdsfsdgrhgf",
			},
			want: "dfhsgr",
		}, {
			name: "0",
			args: args{
				s: "jnhjrgbgdfhfdsfdsfsdgrhgf",
			},
			want: "jnbdfhsgr",
		}, {
			name: "0",
			args: args{
				s: "jnhjrgbgdfhgrhgf",
			},
			want: "jnbdfgrh",
		}, {
			name: "0",
			args: args{
				s: "cbacdcbc",
			},
			want: "acdb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicateLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
