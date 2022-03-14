package leetcode

import (
	"algorithm/pkg/link"
	"reflect"
	"testing"
)

func Test_treeToDoublyList(t *testing.T) {
	type args struct {
		root *link.Node
	}
	tests := []struct {
		name string
		args args
		want *link.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeToDoublyList(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeToDoublyList() = %v, want %v", got, tt.want)
			}
		})
	}
}
