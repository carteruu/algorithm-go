package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
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
				head: &ListNode{1, &ListNode{1, &ListNode{1, nil}}},
				val:  1,
			},
			want: nil,
		}, {
			name: "0",
			args: args{
				head: nil,
				val:  1,
			},
			want: nil,
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				val:  3,
			},
			want: &ListNode{1, &ListNode{2, nil}},
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				val:  2,
			},
			want: &ListNode{1, &ListNode{3, nil}},
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				val:  4,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
