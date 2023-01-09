package leetcode_2

import "testing"

func Test_reinitializePermutation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 2,
			},
			want: 1,
		}, {
			name: "",
			args: args{
				n: 4,
			},
			want: 2,
		}, {
			name: "",
			args: args{
				n: 4,
			},
			want: 2,
		}, {
			name: "",
			args: args{
				n: 6,
			},
			want: 2,
		}, {
			name: "",
			args: args{
				n: 8,
			},
			want: 2,
		}, {
			name: "",
			args: args{
				n: 10,
			},
			want: 2,
		}, {
			name: "",
			args: args{
				n: 1000,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reinitializePermutation(tt.args.n); got != tt.want {
				t.Errorf("reinitializePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
