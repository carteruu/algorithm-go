package leet

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		n int
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
				n: -144,
			},
			want: false,
		}, {
			name: "0",
			args: args{
				n: 14,
			},
			want: false,
		}, {
			name: "0",
			args: args{
				n: 2,
			},
			want: true,
		}, {
			name: "0",
			args: args{
				n: 8,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.n); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
