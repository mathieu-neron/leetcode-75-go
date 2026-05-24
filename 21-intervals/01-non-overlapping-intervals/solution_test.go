package solution

import (
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	t.Skip("TODO: implement eraseOverlapIntervals")
	tests := []struct {
		intervals [][]int
		want int
	}{
		{[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{1, 3}}, 1},
		{[][]int{[]int{1, 2}, []int{1, 2}, []int{1, 2}}, 2},
		{[][]int{[]int{1, 2}, []int{2, 3}}, 0},
	}
	for _, tc := range tests {
		got := eraseOverlapIntervals(tc.intervals)
		if got != tc.want {
			t.Errorf("eraseOverlapIntervals(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
