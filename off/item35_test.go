package off

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *Node
	}
	node1 := &Node{1, nil, nil}
	node10 := &Node{10, node1, nil}
	node11 := &Node{11, node10, nil}
	node13 := &Node{13, node11, nil}
	node7 := &Node{7, node13, nil}

	node13.Random = node7
	node11.Random = node1
	node10.Random = node11
	node1.Random = node7
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{head: node7},
			want: node7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
