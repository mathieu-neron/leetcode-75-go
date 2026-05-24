package solution

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	t.Skip("TODO: implement reverseVowels")
	tests := []struct {
		s string
		want string
	}{
		{"IceCreAm", "AceCreIm"},
		{"leetcode", "leotcede"},
	}
	for _, tc := range tests {
		got := reverseVowels(tc.s)
		if got != tc.want {
			t.Errorf("reverseVowels(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
