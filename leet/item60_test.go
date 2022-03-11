package leet

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want LFUCache
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				capacity: 2,
			},
			want: LFUCache{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lfuCache := ConstructorLFUCache(2)
			lfuCache.Put(1, 1)
			lfuCache.Put(2, 2)
			val1 := lfuCache.Get(1)
			if val1 != 1 {
				t.Errorf("findMedianSortedArrays() = %v, want %v", val1, 1)
			}
			lfuCache.Put(3, 3)
			val2 := lfuCache.Get(2)
			if val2 != -1 {
				t.Errorf("findMedianSortedArrays() = %v, want %v", val2, -1)
			}
			val3 := lfuCache.Get(3)
			if val3 != 3 {
				t.Errorf("findMedianSortedArrays() = %v, want %v", val3, 3)
			}
			lfuCache.Put(4, 4)
			val1 = lfuCache.Get(1)
			if val1 != -1 {
				t.Errorf("findMedianSortedArrays() = %v, want %v", val1, -1)
			}
			val3 = lfuCache.Get(3)
			if val3 != 3 {
				t.Errorf("findMedianSortedArrays() = %v, want %v", val3, 3)
			}
			val4 := lfuCache.Get(4)
			if val4 != 4 {
				t.Errorf("findMedianSortedArrays() = %v, want %v", val4, 4)
			}
		})
	}
}
