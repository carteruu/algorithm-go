package leetcode

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		}, {
			name: "2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		}, {
			name: "3",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 1}},
			},
			want: false,
		}, {
			name: "4",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{1, 2}, {0, 2}},
			},
			want: true,
		}, {
			name: "5",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 2}, {2, 3}, {3, 1}},
			},
			want: false,
		}, {
			name: "6",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{0, 1}, {0, 2}, {1, 2}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
