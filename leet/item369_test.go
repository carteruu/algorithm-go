package leet

import (
	"reflect"
	"testing"
)

func Test_plusOne11(t *testing.T) {
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
			name: "0",
			args: args{head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
			want: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		}, {
			name: "0",
			args: args{head: &ListNode{1, nil}},
			want: &ListNode{2, nil},
		}, {
			name: "0",
			args: args{head: &ListNode{1, &ListNode{9, &ListNode{9, nil}}}},
			want: &ListNode{2, &ListNode{0, &ListNode{0, nil}}},
		}, {
			name: "0",
			args: args{head: &ListNode{1, &ListNode{2, &ListNode{9, nil}}}},
			want: &ListNode{1, &ListNode{3, &ListNode{0, nil}}},
		}, {
			name: "0",
			args: args{head: &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			want: &ListNode{1, &ListNode{0, &ListNode{0, &ListNode{0, nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne11(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne1() = %v, want %v", got, tt.want)
			}
		})
	}
}
