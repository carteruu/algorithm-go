package offer

import "testing"

func Test_numRabbits(t *testing.T) {
	type args struct {
		answers []int
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
				answers: []int{1, 1},
			},
			want: 2,
		}, {
			name: "1",
			args: args{
				answers: []int{1, 0, 1, 0, 0},
			},
			want: 5,
		}, {
			name: "1",
			args: args{
				answers: []int{1, 1, 2},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRabbits(tt.args.answers); got != tt.want {
				t.Errorf("numRabbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
