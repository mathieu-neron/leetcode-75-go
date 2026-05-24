package solution

import (
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	t.Skip("TODO: implement findMaxAverage")
	tests := []struct {
		nums []int
		k int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5.0},
	}
	for _, tc := range tests {
		got := findMaxAverage(tc.nums, tc.k)
		if abs(got-tc.want) > 1e-5 {
			t.Errorf("findMaxAverage(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}

func abs(x float64) float64 { if x < 0 { return -x }; return x }
