package solution

import (
	"testing"
)

func TestNumTilings(t *testing.T) {
	t.Skip("TODO: implement numTilings")
	tests := []struct {
		n int
		want int
	}{
		{3, 5},
		{1, 1},
		{4, 11},
	}
	for _, tc := range tests {
		got := numTilings(tc.n)
		if got != tc.want {
			t.Errorf("numTilings(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
