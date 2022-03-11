package leet

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				k:    1,
			},
			want: &ListNode{5, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}},
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				k:    0,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				k:    4,
			},
			want: &ListNode{3, &ListNode{1, &ListNode{2, nil}}},
		}, {
			name: "0",
			args: args{
				head: nil,
				k:    4,
			},
			want: nil,
		}, {
			name: "0",
			args: args{
				head: &ListNode{2, nil},
				k:    1,
			},
			want: &ListNode{2, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
