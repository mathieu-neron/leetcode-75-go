package solution

import (
	"testing"
)

func TestMaxVowels(t *testing.T) {
	t.Skip("TODO: implement maxVowels")
	tests := []struct {
		s string
		k int
		want int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
	}
	for _, tc := range tests {
		got := maxVowels(tc.s, tc.k)
		if got != tc.want {
			t.Errorf("maxVowels(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
