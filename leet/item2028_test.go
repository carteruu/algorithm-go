package leet

import (
	"reflect"
	"testing"
)

func Test_missingRolls(t *testing.T) {
	type args struct {
		rolls []int
		mean  int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0",
			args: args{
				rolls: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				mean:  1,
				n:     1,
			},
			want: []int{1},
		}, {
			name: "0",
			args: args{
				rolls: []int{1},
				mean:  3,
				n:     1,
			},
			want: []int{5},
		}, {
			name: "0",
			args: args{
				rolls: []int{4, 1, 4, 3, 3, 5, 3, 5, 4, 2, 3, 1, 2, 6, 2, 4, 2, 5, 5, 5, 2, 2, 4, 6, 6, 5, 3, 1, 5, 1, 1, 1, 3, 1, 2, 1, 3, 2, 2, 5, 2, 1, 6, 5, 2, 2, 5, 2, 5, 3, 1, 2, 6, 3, 6, 1, 1, 6, 1, 3, 5, 5, 3, 5, 5, 3, 1, 6, 5, 1, 6, 4, 1, 3, 3, 6, 3, 3, 3, 6, 1, 3, 2, 4, 4, 5, 6, 6, 3, 4, 3, 2, 5, 6, 2, 6, 5, 1, 4, 4, 5, 2, 2, 6, 4, 2, 3, 5, 3, 3, 4, 6, 3, 6, 2, 3, 1, 2, 2, 5, 2, 3, 5, 5, 3, 4, 1, 4, 4, 3, 4, 5, 4, 3, 4, 1, 3, 2, 5, 1, 5, 4, 3, 6, 5, 2, 1, 1, 2, 2, 6, 5, 6, 4, 1, 6, 5, 5, 6, 3, 5, 3, 2, 6, 2, 3, 5, 6, 6, 3, 4, 5, 6, 6, 3, 1, 5, 4, 6, 6, 3, 3, 2, 5, 6, 3, 2, 1, 3, 1, 6, 6, 5, 4, 2, 5, 5, 1, 4, 5, 3, 5, 1, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 4, 2, 3, 2, 4, 1, 1, 1, 6, 1, 3, 5, 3, 4, 4, 3, 6, 1, 6, 2, 1, 4, 1, 2, 6, 4, 3, 4, 2, 5, 2, 6, 1, 4, 1, 1, 5, 5, 4, 2, 3, 4, 5, 1, 5, 1, 4, 6, 6, 3, 6, 6, 1, 6, 2, 5, 3, 6, 2, 5, 3, 3, 3, 6, 4, 5, 3, 3, 3, 6, 3, 2, 6, 3, 2, 3, 3, 5, 5, 5, 3, 6, 5, 1, 6, 4, 5, 1, 4, 1, 5, 1, 3, 3, 1, 5, 5, 6, 4, 6, 2, 5, 2, 1, 3, 5, 6, 5, 4, 5, 2, 1, 6, 6, 6, 5, 5, 5, 6, 1, 4, 6, 4, 4, 5, 5, 3},
				mean:  5,
				n:     10240,
			},
			want: []int{},
		}, {
			name: "0",
			args: args{
				rolls: []int{6, 3, 4, 3, 5, 3},
				mean:  1,
				n:     6,
			},
			want: []int{},
		}, {
			name: "0",
			args: args{
				rolls: []int{3, 2, 4, 3},
				mean:  4,
				n:     2,
			},
			want: []int{6, 6},
		}, {
			name: "0",
			args: args{
				rolls: []int{1, 2, 3, 4},
				mean:  6,
				n:     4,
			},
			want: []int{},
		}, {
			name: "0",
			args: args{
				rolls: []int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5},
				mean:  4,
				n:     40,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingRolls(tt.args.rolls, tt.args.mean, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingRolls() = %v, want %v", got, tt.want)
			}
		})
	}
}
