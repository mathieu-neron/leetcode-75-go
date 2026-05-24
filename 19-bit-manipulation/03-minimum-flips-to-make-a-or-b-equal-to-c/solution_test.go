package solution

import (
	"testing"
)

func TestMinFlips(t *testing.T) {
	t.Skip("TODO: implement minFlips")
	tests := []struct {
		a int
		b int
		c int
		want int
	}{
		{2, 6, 5, 3},
		{4, 2, 7, 1},
		{1, 2, 3, 0},
	}
	for _, tc := range tests {
		got := minFlips(tc.a, tc.b, tc.c)
		if got != tc.want {
			t.Errorf("minFlips(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
