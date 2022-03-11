package off

import "testing"

func Test_minArray(t *testing.T) {
	type args struct {
		numbers []int
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
				[]int{3, 4, 5, 1, 2},
			},
			want: 1,
		}, {
			name: "1",
			args: args{
				[]int{2, 2, 2, 0, 1},
			},
			want: 0,
		}, {
			name: "1",
			args: args{
				[]int{1, 3},
			},
			want: 1,
		}, {
			name: "1",
			args: args{
				[]int{2, 2, 2, 1, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
