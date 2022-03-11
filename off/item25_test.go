package off

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
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
				head: &ListNode{Val: 1, Next: &ListNode{2, &ListNode{3, &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
				k:    2,
			},
			want: &ListNode{Val: 2, Next: &ListNode{1, &ListNode{4, &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_reverseK(t *testing.T) {
	type args struct {
		head *ListNode
		tail *ListNode
	}
	tests := []struct {
		name  string
		args  args
		want  *ListNode
		want1 *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{2, &ListNode{3, &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
				tail: nil,
			},
			want:  &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
			want1: &ListNode{Val: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newHead, newTail := reverseK(tt.args.head, tt.args.tail)
			if !reflect.DeepEqual(newHead, tt.want) {
				t.Errorf("reverseK() got = %v, want %v", newHead, tt.want)
			}
			if !reflect.DeepEqual(newTail, tt.want1) {
				t.Errorf("reverseK() got1 = %v, want %v", newTail, tt.want1)
			}
		})
	}
}
