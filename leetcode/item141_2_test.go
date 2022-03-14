package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	node3 := &ListNode{3, nil}
	head := &ListNode{1, &ListNode{2, node3}}
	node3.Next = head
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			},
			want: false,
		}, {
			name: "0",
			args: args{
				head: head,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
