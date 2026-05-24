package solution

import (
	"testing"
)

func TestRemoveStars(t *testing.T) {
	t.Skip("TODO: implement removeStars")
	tests := []struct {
		s string
		want string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
	}
	for _, tc := range tests {
		got := removeStars(tc.s)
		if got != tc.want {
			t.Errorf("removeStars(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
