package solution

import (
	"testing"
)

func TestMaxScore(t *testing.T) {
	t.Skip("TODO: implement maxScore")
	tests := []struct {
		nums1 []int
		nums2 []int
		k int
		want int64
	}{
		{[]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3, 12},
		{[]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1, 30},
	}
	for _, tc := range tests {
		got := maxScore(tc.nums1, tc.nums2, tc.k)
		if got != tc.want {
			t.Errorf("maxScore(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
