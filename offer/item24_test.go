package offer

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
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
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			},
			want: &ListNode{3, &ListNode{2, &ListNode{1, nil}}},
		}, {
			name: "2",
			args: args{
				head: &ListNode{1, nil},
			},
			want: &ListNode{1, nil},
		}, {
			name: "2",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
