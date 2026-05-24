package solution

import (
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	t.Skip("TODO: implement findMinArrowShots")
	tests := []struct {
		points [][]int
		want int
	}{
		{[][]int{[]int{10, 16}, []int{2, 8}, []int{1, 6}, []int{7, 12}}, 2},
		{[][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8}}, 4},
		{[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}}, 2},
	}
	for _, tc := range tests {
		got := findMinArrowShots(tc.points)
		if got != tc.want {
			t.Errorf("findMinArrowShots(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
