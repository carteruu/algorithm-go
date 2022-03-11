package off

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
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
				lists: []*ListNode{
					&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
					&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
					&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				},
			},
			want: &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{3, nil}}}}}}}}},
		}, {
			name: "0",
			args: args{
				lists: []*ListNode{
					&ListNode{1, nil},
					&ListNode{0, nil},
				},
			},
			want: &ListNode{0, &ListNode{1, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
