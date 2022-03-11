package off

import (
	"testing"
)

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				nums: []int{10, 2},
			},
			want: "102",
		}, {
			name: "1",
			args: args{
				nums: []int{3, 30, 34, 5, 9},
			},
			want: "3033459",
		}, {
			name: "1",
			args: args{
				nums: []int{3, 30, 34, 95, 90},
			},
			want: "303349095",
		}, {
			name: "1",
			args: args{
				nums: []int{3, 30, 34, 95, 90, 9, 333, 321},
			},
			want: "303349095",
		}, {
			name: "1",
			args: args{
				nums: []int{121, 12},
			},
			want: "12112",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
