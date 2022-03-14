package leetcode

import (
	"reflect"
	"testing"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			},
		}, {
			name: "1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			if got := tt.args.head; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
