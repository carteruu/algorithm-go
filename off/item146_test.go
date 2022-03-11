package off

import (
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	type fields struct {
		capacity int
		count    int
		cache    map[int]*node
		head     *node
		tail     *node
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &LRUCache{
				capacity: tt.fields.capacity,
				count:    tt.fields.count,
				cache:    tt.fields.cache,
				head:     tt.fields.head,
				tail:     tt.fields.tail,
			}
			if got := this.Get(tt.args.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Put(t *testing.T) {
	type fields struct {
		capacity int
		count    int
		cache    map[int]*node
		head     *node
		tail     *node
	}
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &LRUCache{
				capacity: tt.fields.capacity,
				count:    tt.fields.count,
				cache:    tt.fields.cache,
				head:     tt.fields.head,
				tail:     tt.fields.tail,
			}
		})
	}
}
