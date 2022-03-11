package leet

import "testing"

func Test_minimumTimeRequired(t *testing.T) {
	type args struct {
		jobs []int
		k    int
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
				jobs: []int{3, 2, 3},
				k:    3,
			},
			want: 3,
		}, {
			name: "1",
			args: args{
				jobs: []int{1, 2, 4, 7, 8},
				k:    2,
			},
			want: 11,
		}, {
			name: "1",
			args: args{
				jobs: []int{250, 250, 256, 251, 254, 254, 251, 255, 250, 252, 254, 255},
				k:    10,
			},
			want: 501,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeRequired(tt.args.jobs, tt.args.k); got != tt.want {
				t.Errorf("minimumTimeRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
