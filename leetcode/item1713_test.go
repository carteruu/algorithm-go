package leetcode

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		target []int
		arr    []int
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
				target: []int{5, 1, 3},
				arr:    []int{9, 4, 2, 3, 4},
			},
			want: 2,
		}, {
			name: "2",
			args: args{
				target: []int{17, 18, 14, 13, 6, 9, 1, 3, 2, 20},
				arr:    []int{18, 15, 14, 6, 6, 13, 15, 20, 2, 6},
			},
			want: 6,
		}, {
			name: "3",
			args: args{
				target: []int{1, 3, 8},
				arr:    []int{2, 6},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
