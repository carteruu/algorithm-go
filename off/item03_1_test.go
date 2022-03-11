package off

import "testing"

func Test_findRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5, 3},
			},
			want: 2,
		}, {
			name: "0",
			args: args{
				nums: []int{0, 0},
			},
			want: 0,
		}, {
			name: "0",
			args: args{
				nums: []int{3, 4, 2, 0, 0, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
