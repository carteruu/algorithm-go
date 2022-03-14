package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
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
				intervals: [][]int{{1, 4}, {5, 6}},
			},
			want: [][]int{{1, 4}, {5, 6}},
		}, {
			name: "0",
			args: args{
				intervals: [][]int{{1, 1}},
			},
			want: [][]int{{1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge1(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
