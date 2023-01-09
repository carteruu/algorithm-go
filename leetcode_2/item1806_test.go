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
			want: 1,
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
