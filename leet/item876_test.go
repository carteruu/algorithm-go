package leet

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
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
			args: args{
				&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			},
			want: &ListNode{2, &ListNode{3, nil}},
		}, {
			name: "0",
			args: args{
				&ListNode{1, nil},
			},
			want: &ListNode{1, nil},
		}, {
			name: "0",
			args: args{
				nil,
			},
			want: nil,
		}, {
			name: "0",
			args: args{
				&ListNode{1, &ListNode{2, nil}},
			},
			want: &ListNode{2, nil},
		}, {
			name: "0",
			args: args{
				&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			},
			want: &ListNode{3, &ListNode{4, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
