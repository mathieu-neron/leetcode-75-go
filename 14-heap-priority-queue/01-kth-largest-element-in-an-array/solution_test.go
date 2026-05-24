package solution

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	t.Skip("TODO: implement findKthLargest")
	tests := []struct {
		nums []int
		k int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, tc := range tests {
		got := findKthLargest(tc.nums, tc.k)
		if got != tc.want {
			t.Errorf("findKthLargest(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
