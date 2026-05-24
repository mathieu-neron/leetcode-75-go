package solution

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	t.Skip("TODO: implement uniquePaths")
	tests := []struct {
		m int
		n int
		want int
	}{
		{3, 7, 28},
		{3, 2, 3},
	}
	for _, tc := range tests {
		got := uniquePaths(tc.m, tc.n)
		if got != tc.want {
			t.Errorf("uniquePaths(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
