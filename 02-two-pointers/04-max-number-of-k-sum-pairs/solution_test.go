package solution

import (
	"testing"
)

func TestMaxOperations(t *testing.T) {
	t.Skip("TODO: implement maxOperations")
	tests := []struct {
		nums []int
		k int
		want int
	}{
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
	}
	for _, tc := range tests {
		got := maxOperations(tc.nums, tc.k)
		if got != tc.want {
			t.Errorf("maxOperations(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
