package leet

import (
	"reflect"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: [][]int{{2, 6}, {1, 7}, {1, 2, 5}, {1, 1, 6}},
		}, {
			name: "0",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     2,
			},
			want: [][]int{{1, 1}, {2}},
		}, {
			name: "0",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: [][]int{{2, 6}, {2, 2, 2, 2}, {1, 7}, {1, 2, 5}, {1, 1, 6}, {1, 1, 2, 2, 2}, {1, 1, 1, 5}, {1, 1, 1, 1, 2, 2}, {1, 1, 1, 1, 1, 1, 2}, {1, 1, 1, 1, 1, 1, 1, 1}},
		}, {
			name: "0",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5, 1, 2, 7, 6, 1, 5},
				target:     30,
			},
			want: [][]int{{2, 6}, {2, 2, 2, 2}, {1, 7}, {1, 2, 5}, {1, 1, 6}, {1, 1, 2, 2, 2}, {1, 1, 1, 5}, {1, 1, 1, 1, 2, 2}, {1, 1, 1, 1, 1, 1, 2}, {1, 1, 1, 1, 1, 1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
