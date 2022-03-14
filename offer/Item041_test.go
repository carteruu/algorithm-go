package offer

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want MovingAverage
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			movingAverage := Constructor(3)
			if got := movingAverage.Next(1); !reflect.DeepEqual(got, 1.0) {
				t.Errorf("Constructor() = %v, want %v", got, 1.0)
			}
			if got := movingAverage.Next(10); !reflect.DeepEqual(got, 5.5) {
				t.Errorf("Constructor() = %v, want %v", got, 5.5)
			}
			if got := movingAverage.Next(3); !reflect.DeepEqual(got, 4.6667) {
				t.Errorf("Constructor() = %v, want %v", got, 4.6667)
			}
			if got := movingAverage.Next(5); !reflect.DeepEqual(got, 6.0) {
				t.Errorf("Constructor() = %v, want %v", got, 6.0)
			}

		})
	}
}
