package leet

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2, 4},
			},
			want: 2.5,
		}, {
			name: "1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{},
			},
			want: 2.0,
		}, {
			name: "2",
			args: args{
				nums1: []int{1},
				nums2: []int{2},
			},
			want: 1.5,
		}, {
			name: "3",
			args: args{
				nums1: []int{},
				nums2: []int{},
			},
			want: 0,
		}, {
			name: "4",
			args: args{
				nums1: []int{0, 0, 0},
				nums2: []int{2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
