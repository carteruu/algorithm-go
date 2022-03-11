package leet

import (
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 2},
		}, {
			name: "0",
			args: args{
				nums: []int{3, 17},
			},
			want: []int{3},
		}, {
			name: "0",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		}, {
			name: "0",
			args: args{
				nums: []int{1, 2, 4, 8},
			},
			want: []int{1, 2, 4, 8},
		}, {
			name: "0",
			args: args{
				nums: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912, 1073741824},
			},
			want: []int{1, 2, 4, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
