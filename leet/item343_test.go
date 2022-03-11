package leet

import "testing"

func Test_integerBreak(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				n: 2,
			},
			want: 1,
		}, {
			name: "1",
			args: args{
				n: 10,
			},
			want: 36,
		}, {
			name: "3",
			args: args{
				n: 3,
			},
			want: 2,
		}, {
			name: "1",
			args: args{
				n: 4,
			},
			want: 4,
		}, {
			name: "1",
			args: args{
				n: 58,
			},
			want: 1549681956,
		}, {
			name: "1",
			args: args{
				n: 5,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerBreak(tt.args.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
