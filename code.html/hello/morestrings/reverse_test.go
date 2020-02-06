package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				s: "Hello, world",
			},
			want: "dlrow ,olleH",
		},
		{
			name: "t2",
			args: args{
				s: "Hello, 世界",
			},
			want: "界世 ,olleH",
		},
		{
			name: "t3",
			args: args{
				s: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseRunes(tt.args.s); got != tt.want {
				t.Errorf("ReverseRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
