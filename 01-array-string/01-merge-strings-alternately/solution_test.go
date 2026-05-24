package mergestrings

import "testing"

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		word1, word2, want string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
		{"", "xyz", "xyz"},
		{"xyz", "", "xyz"},
	}
	for _, tc := range tests {
		got := mergeAlternately(tc.word1, tc.word2)
		if got != tc.want {
			t.Errorf("mergeAlternately(%q,%q) = %q, want %q", tc.word1, tc.word2, got, tc.want)
		}
	}
}
