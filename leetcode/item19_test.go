package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
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
				n:    1,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				n:    2,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				n:    5,
			},
			want: &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}},
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, nil},
				n:    1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
