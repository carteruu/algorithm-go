package item2027

import "testing"

func Test_minimumMoves(t *testing.T) {
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
			args: args{s: "XXOX"},
			want: 2,
		}, {
			name: "0",
			args: args{s: "XXX"},
			want: 1,
		}, {
			name: "0",
			args: args{s: "OOOO"},
			want: 0,
		}, {
			name: "0",
			args: args{s: "OXOX"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMoves(tt.args.s); got != tt.want {
				t.Errorf("minimumMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
