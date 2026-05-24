package solution

import (
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	t.Skip("TODO: implement minCostClimbingStairs")
	tests := []struct {
		cost []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tc := range tests {
		got := minCostClimbingStairs(tc.cost)
		if got != tc.want {
			t.Errorf("minCostClimbingStairs(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
