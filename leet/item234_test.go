package leet

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		//{
		//	name: "0",
		//	args: args{
		//		head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		//	},
		//	want: false,
		//},
		{
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{1, nil}}},
			},
			want: true,
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{1, &ListNode{1, nil}}}},
			},
			want: false,
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}},
			},
			want: true,
		}, {
			name: "0",
			args: args{
				head: &ListNode{1, nil},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome4(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome1() = %v, want %v", got, tt.want)
			}
		})
	}
}
