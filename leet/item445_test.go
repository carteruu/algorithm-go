package leet

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers445(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
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
				l1: &ListNode{
					Val: 9,
				},
				l2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val: 8,
											Next: &ListNode{
												Val: 9,
												Next: &ListNode{
													Val: 9,
													Next: &ListNode{
														Val: 9,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val: 9,
											Next: &ListNode{
												Val: 9,
												Next: &ListNode{
													Val: 9,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "0",
			args: args{
				l1: &ListNode{
					Val: 9,
				},
				l2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val: 9,
											Next: &ListNode{
												Val: 9,
												Next: &ListNode{
													Val: 9,
													Next: &ListNode{
														Val: 9,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val: 9,
											Next: &ListNode{
												Val: 9,
												Next: &ListNode{
													Val: 9,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers445(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers445() = %v, want %v", got, tt.want)
			}
		})
	}
}
