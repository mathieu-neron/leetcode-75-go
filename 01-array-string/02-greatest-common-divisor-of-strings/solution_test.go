package solution

import (
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		want string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
	}
	for _, tc := range tests {
		got := gcdOfStrings(tc.str1, tc.str2)
		if got != tc.want {
			t.Errorf("gcdOfStrings(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
