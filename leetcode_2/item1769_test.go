package leetcode_2

import (
	"reflect"
	"testing"
)

func Test_minOperations1769_1(t *testing.T) {
	type args struct {
		boxes string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				boxes: "110",
			},
			want: []int{1, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations1769_1(tt.args.boxes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations1769_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
