package leet

import "testing"

func Test_search11(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				nums:   []int{3, 1},
				target: 1,
			},
			want: 1,
		}, {
			name: "0",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		}, {
			name: "0",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		}, {
			name: "0",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		}, {
			name: "0",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		}, {
			name: "0",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: -1,
		}, {
			name: "0",
			args: args{
				nums:   []int{1, 3, 5},
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search11(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
