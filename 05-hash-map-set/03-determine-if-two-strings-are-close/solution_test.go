package solution

import (
	"testing"
)

func TestCloseStrings(t *testing.T) {
	t.Skip("TODO: implement closeStrings")
	tests := []struct {
		word1 string
		word2 string
		want bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
	}
	for _, tc := range tests {
		got := closeStrings(tc.word1, tc.word2)
		if got != tc.want {
			t.Errorf("closeStrings(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
