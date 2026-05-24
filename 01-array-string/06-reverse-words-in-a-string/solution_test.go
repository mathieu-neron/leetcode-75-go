package solution

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	t.Skip("TODO: implement reverseWords")
	tests := []struct {
		s string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "hello world"},
		{"a good   example", "example good a"},
	}
	for _, tc := range tests {
		got := reverseWords(tc.s)
		if got != tc.want {
			t.Errorf("reverseWords(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
