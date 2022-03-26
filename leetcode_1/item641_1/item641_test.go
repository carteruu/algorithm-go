package item641_1

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want MyCircularDeque
	}{
		{
			name: "0",
			args: args{k: 3},
			want: Constructor(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Constructor(tt.args.k)
			q.InsertLast(1)
			q.InsertLast(2)
			q.InsertFront(3)
			if got := q.InsertFront(4); got {
				t.Errorf("Constructor() = %v, want %v", got, false)
			}
		})
	}
}
