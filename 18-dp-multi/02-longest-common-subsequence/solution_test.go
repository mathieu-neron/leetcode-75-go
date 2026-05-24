package solution

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	t.Skip("TODO: implement longestCommonSubsequence")
	tests := []struct {
		text1 string
		text2 string
		want int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, tc := range tests {
		got := longestCommonSubsequence(tc.text1, tc.text2)
		if got != tc.want {
			t.Errorf("longestCommonSubsequence(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
