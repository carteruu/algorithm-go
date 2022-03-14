package leetcode

import "testing"

func Test_knightProbability(t *testing.T) {
	type args struct {
		n      int
		k      int
		row    int
		column int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				3,
				2,
				0,
				0,
			},
			want: 0.06250,
		}, {
			name: "1",
			args: args{
				8,
				30,
				6,
				4,
			},
			want: 0.00019052566298333648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knightProbability(tt.args.n, tt.args.k, tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("knightProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
