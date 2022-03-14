package offer

import "testing"

func Test_replaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{
				s: "We are happy.",
			},
			want: "We%20are%20happy.",
		}, {
			name: "1",
			args: args{
				s: "",
			},
			want: "",
		}, {
			name: "2",
			args: args{
				s: " ",
			},
			want: "%20",
		}, {
			name: "3",
			args: args{
				s: " 刚 发 的 ",
			},
			want: "%20刚%20发%20的%20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.args.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
