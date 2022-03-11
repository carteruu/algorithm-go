package off

import (
	"reflect"
	"testing"
)

func Test_treeToDoublyList(t *testing.T) {
	type args struct {
		root *Node
	}
	node5 := &Node{5, nil, nil}
	node4 := &Node{4, nil, node5}
	node3 := &Node{3, nil, node4}
	node2 := &Node{2, nil, node3}
	node1 := &Node{1, nil, node2}

	node5.Right = node1
	node1.Left = node5
	node2.Left = node1
	node3.Left = node2
	node4.Left = node3
	node5.Left = node4

	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{root: &Node{4, &Node{2, &Node{1, nil, nil}, &Node{3, nil, nil}}, &Node{5, nil, nil}}},
			want: node1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeToDoublyList(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeToDoublyList() = %v, want %v", got, tt.want)
			}
		})
	}
}
