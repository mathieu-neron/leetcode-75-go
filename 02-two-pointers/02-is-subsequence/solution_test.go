package solution

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	t.Skip("TODO: implement isSubsequence")
	tests := []struct {
		s string
		t string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for _, tc := range tests {
		got := isSubsequence(tc.s, tc.t)
		if got != tc.want {
			t.Errorf("isSubsequence(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
