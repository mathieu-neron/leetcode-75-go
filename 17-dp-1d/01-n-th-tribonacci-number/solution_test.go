package solution

import (
	"testing"
)

func TestTribonacci(t *testing.T) {
	t.Skip("TODO: implement tribonacci")
	tests := []struct {
		n int
		want int
	}{
		{4, 4},
		{25, 1389537},
		{0, 0},
	}
	for _, tc := range tests {
		got := tribonacci(tc.n)
		if got != tc.want {
			t.Errorf("tribonacci(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
