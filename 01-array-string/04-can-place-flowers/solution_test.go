package solution

import (
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
	}
	for _, tc := range tests {
		got := canPlaceFlowers(tc.flowerbed, tc.n)
		if got != tc.want {
			t.Errorf("canPlaceFlowers(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
