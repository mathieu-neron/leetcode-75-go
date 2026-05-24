package solution

import (
	"testing"
)

func TestLargestAltitude(t *testing.T) {
	t.Skip("TODO: implement largestAltitude")
	tests := []struct {
		gain []int
		want int
	}{
		{[]int{-5, 1, 5, 0, -7}, 1},
		{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}
	for _, tc := range tests {
		got := largestAltitude(tc.gain)
		if got != tc.want {
			t.Errorf("largestAltitude(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
