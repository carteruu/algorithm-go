package leet

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
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
				head: &ListNode{4, &ListNode{3, &ListNode{3, &ListNode{4, nil}}}},
				x:    3,
			},
			want: &ListNode{4, &ListNode{3, &ListNode{3, &ListNode{4, nil}}}},
		}, // TODO: Add test cases.
		{
			name: "0",
			args: args{
				head: &ListNode{4, &ListNode{3, &ListNode{3, &ListNode{1, nil}}}},
				x:    3,
			},
			want: &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{3, nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
