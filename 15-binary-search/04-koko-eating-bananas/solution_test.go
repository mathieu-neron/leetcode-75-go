package solution

import (
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	t.Skip("TODO: implement minEatingSpeed")
	tests := []struct {
		piles []int
		h int
		want int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
	}
	for _, tc := range tests {
		got := minEatingSpeed(tc.piles, tc.h)
		if got != tc.want {
			t.Errorf("minEatingSpeed(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
