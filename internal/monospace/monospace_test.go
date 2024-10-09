package monospace

import (
	"fmt"
	"testing"
)

func TestShorten(t *testing.T) {
	type args struct {
		str string
		max int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				str: "hello world",
				max: 5,
			},
			want: "h...d",
		},
		{
			args: args{
				str: "hello world",
				max: 6,
			},
			want: "he...d",
		},
		{
			args: args{
				str: "hello world!",
				max: 5,
			},
			want: "h...!",
		},
		{
			args: args{
				str: "hello world!",
				max: 6,
			},
			want: "he...!",
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Shorten(%s, %d)", tt.args.str, tt.args.max)
		t.Run(name, func(t *testing.T) {
			if got := Shorten(tt.args.str, tt.args.max); got != tt.want {
				t.Errorf("Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		str     string
		wantLen int
	}{
		{"", 0},
		{"a", 1},
		{"hello", 5},
		{"world", 5},
		{"hello world", 11},
		{"hello 世界", 10},
		{"hello 🌎", 8},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := String(tt.str); got != tt.wantLen {
				t.Errorf("String(%q) = %v, want %v", tt.str, got, tt.wantLen)
			}
		})
	}
}

func TestRune(t *testing.T) {
	tests := []struct {
		r    rune
		want int
	}{
		{'a', 1},
		{'世', 2},
		{'🌎', 2},
		{'\t', 0},
		{'\n', 0},
		{'\x1b', 0},
	}

	for _, tt := range tests {
		t.Run(string(tt.r), func(t *testing.T) {
			if got := Rune(tt.r); got != tt.want {
				t.Errorf("Rune(%q) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}
