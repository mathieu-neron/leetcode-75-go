package solution

import (
	"testing"
)

func TestTotalCost(t *testing.T) {
	t.Skip("TODO: implement totalCost")
	tests := []struct {
		costs []int
		k int
		candidates int
		want int64
	}{
		{[]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4, 11},
		{[]int{1, 2, 4, 1}, 3, 3, 4},
	}
	for _, tc := range tests {
		got := totalCost(tc.costs, tc.k, tc.candidates)
		if got != tc.want {
			t.Errorf("totalCost(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
