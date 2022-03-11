package inte

import (
	"reflect"
	"testing"
)

func Test_getMaxMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "1",
			args: args{matrix: [][]int{
				{-1, 0},
				{0, -1},
			}},
			want: []int{0, 1, 0, 1},
		}, {name: "1",
			args: args{matrix: [][]int{
				{-4, -5},
			}},
			want: []int{0, 0, 0, 0},
		},
		{name: "2",
			args: args{matrix: [][]int{
				{9, -8, 1, 3, -2},
				{-3, 7, 6, -2, 4},
				{6, -4, -4, 8, -7}},
			},
			want: []int{0, 0, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMaxMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
